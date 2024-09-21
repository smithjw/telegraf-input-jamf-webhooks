//go:generate ../../../tools/readme_config_includer/generator
package webhooks

import (
	_ "embed"
	"fmt"
	"net"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/mux"
	"github.com/smithjw/telegraf-input-jamf-webhooks/plugins/inputs/webhooks/jamf"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/config"
	"github.com/influxdata/telegraf/plugins/inputs"
)

const (
	defaultReadTimeout  = 10 * time.Second
	defaultWriteTimeout = 10 * time.Second
)

type Webhook interface {
	Register(router *mux.Router, acc telegraf.Accumulator, log telegraf.Logger)
}

func init() {
	inputs.Add("external_webhooks", func() telegraf.Input { return NewWebhooks() })
}

type Webhooks struct {
	ServiceAddress string          `toml:"service_address"`
	ReadTimeout    config.Duration `toml:"read_timeout"`
	WriteTimeout   config.Duration `toml:"write_timeout"`

	Jamf *jamf.JamfWebhook `toml:"jamf"`

	Log telegraf.Logger `toml:"-"`
	srv *http.Server
}

func NewWebhooks() *Webhooks {
	return &Webhooks{}
}

func (wb *Webhooks) SampleConfig() string {
	return `
  ## Address and port to host Webhook listener on
  service_address = ":1619"

  [inputs.external_webhooks.plex]
    path = "/jamf"
`
}

func (wb *Webhooks) Gather(_ telegraf.Accumulator) error {
	return nil
}

// AvailableWebhooks Looks for fields which implement Webhook interface
func (wb *Webhooks) AvailableWebhooks() []Webhook {
	webhooks := make([]Webhook, 0)
	s := reflect.ValueOf(wb).Elem()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)

		if !f.CanInterface() {
			continue
		}

		if wbPlugin, ok := f.Interface().(Webhook); ok {
			if !reflect.ValueOf(wbPlugin).IsNil() {
				webhooks = append(webhooks, wbPlugin)
			}
		}
	}

	return webhooks
}

func (wb *Webhooks) Start(acc telegraf.Accumulator) error {
	if wb.ReadTimeout < config.Duration(time.Second) {
		wb.ReadTimeout = config.Duration(defaultReadTimeout)
	}
	if wb.WriteTimeout < config.Duration(time.Second) {
		wb.WriteTimeout = config.Duration(defaultWriteTimeout)
	}

	r := mux.NewRouter()

	for _, webhook := range wb.AvailableWebhooks() {
		webhook.Register(r, acc, wb.Log)
	}

	wb.srv = &http.Server{
		Handler:      r,
		ReadTimeout:  time.Duration(wb.ReadTimeout),
		WriteTimeout: time.Duration(wb.WriteTimeout),
	}

	ln, err := net.Listen("tcp", wb.ServiceAddress)
	if err != nil {
		return fmt.Errorf("error starting server: %w", err)
	}

	go func() {
		if err := wb.srv.Serve(ln); err != nil {
			if err != http.ErrServerClosed {
				acc.AddError(fmt.Errorf("error listening: %w", err))
			}
		}
	}()

	wb.Log.Infof("Started the webhooks service on %s", wb.ServiceAddress)

	return nil
}

func (wb *Webhooks) Stop() {
	wb.srv.Close()
	wb.Log.Infof("Stopping the Webhooks service")
}
