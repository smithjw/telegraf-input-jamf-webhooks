package jamf

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/common/auth"
)

type JamfWebhook struct {
	Path   string
	Secret string
	acc    telegraf.Accumulator
	log    telegraf.Logger
	auth.BasicAuth
}

func (jwh *JamfWebhook) Register(router *mux.Router, acc telegraf.Accumulator, log telegraf.Logger) {
	router.HandleFunc(jwh.Path, jwh.eventHandler).Methods("POST")

	jwh.log = log
	jwh.log.Infof("Started the webhooks_jamf on %s", jwh.Path)
	jwh.acc = acc
}

func (jwh *JamfWebhook) eventHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		jwh.log.Errorf("Failed to read request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Define a temporary structure to extract the event type from the request body
	var webhookPayload struct {
		Webhook struct {
			Event string `json:"webhookEvent"`
		} `json:"webhook"`
	}

	// Unmarshal the request body to extract the event type
	if err := json.Unmarshal(data, &webhookPayload); err != nil {
		jwh.log.Errorf("Failed to parse request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract the event type from the parsed payload
	eventType := webhookPayload.Webhook.Event
	if eventType == "" {
		jwh.log.Error("No webhookEvent found in the request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	e, err := jwh.NewEvent(data, eventType)
	if err != nil {
		jwh.log.Errorf("Failed to create new event: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if e != nil {
		p := e.NewMetric()
		jwh.acc.AddFields("jamf_webhooks", p.Fields(), p.Tags(), p.Time())
	}

	w.WriteHeader(http.StatusOK)
}

func generateEvent(data []byte, event Event) (Event, error) {
	err := json.Unmarshal(data, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

type newEventError struct {
	s string
}

func (e *newEventError) Error() string {
	return e.s
}

func (jh *JamfWebhook) NewEvent(data []byte, name string) (Event, error) {
	jh.log.Debugf("New %v event received", name)
	switch name {
	case "ComputerInventoryCompleted":
		return generateEvent(data, &ComputerInventoryCompletedEvent{})
	case "ComputerPolicyFinished":
		return generateEvent(data, &ComputerPolicyFinishedEvent{})
	default:
		return nil, &newEventError{"Not a recognized event type"}
	}
}
