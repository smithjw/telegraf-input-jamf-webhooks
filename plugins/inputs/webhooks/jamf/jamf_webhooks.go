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
	Path       string
	URLQueries []string
	acc        telegraf.Accumulator
	log        telegraf.Logger
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

	if !jwh.Verify(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

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
		tags := p.Tags()

		for _, queryParam := range jwh.URLQueries {
			value := r.URL.Query().Get(queryParam)
			if value != "" {
				tags[queryParam] = value
			}
		}
		jwh.acc.AddFields("jamf_webhooks", p.Fields(), tags, p.Time())
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

func (jwh *JamfWebhook) NewEvent(data []byte, name string) (Event, error) {
	jwh.log.Debugf("New %v event received", name)
	switch name {
	case "ComputerInventoryCompleted":
		return generateEvent(data, &ComputerEvent{})
	case "ComputerPolicyFinished":
		return generateEvent(data, &ComputerPolicyFinishedEvent{})
	case "ComputerAdded":
		return generateEvent(data, &ComputerEvent{})
	case "ComputerCheckIn":
		return generateEvent(data, &ComputerCheckInEvent{})
	// case "ComputerPatchPolicyCompleted":
	// 	return generateEvent(data, &ComputerPatchPolicyCompletedEvent{})
	case "ComputerPushCapabilityChanged":
		return generateEvent(data, &ComputerEvent{})
	case "DeviceAddedToDEP":
		return generateEvent(data, &DeviceAddedToDEPEvent{})
	case "JSSShutdown":
		return generateEvent(data, &JSSEvent{})
	case "JSSStartup":
		return generateEvent(data, &JSSEvent{})
	case "MobileDeviceCheckIn":
		return generateEvent(data, &MobileDeviceEvent{})
	case "MobileDeviceCommandCompleted":
		return generateEvent(data, &MobileDeviceEvent{})
	case "MobileDeviceEnrolled":
		return generateEvent(data, &MobileDeviceEvent{})
	case "MobileDeviceInventoryCompleted":
		return generateEvent(data, &MobileDeviceEvent{})
	case "MobileDevicePushSent":
		return generateEvent(data, &MobileDeviceEvent{})
	case "MobileDeviceUnEnrolled":
		return generateEvent(data, &MobileDeviceEvent{})
	// case "PatchSoftwareTitleUpdated":
	// 	return generateEvent(data, &PatchSoftwareTitleUpdatedEvent{})
	case "PushSent":
		return generateEvent(data, &PushSentEvent{})
	// case "RestAPIOperation":
	// 	return generateEvent(data, &RestAPIOperationEvent{})
	// case "SmartGroupComputerMembershipChange":
	// 	return generateEvent(data, &SmartGroupComputerMembershipChangeEvent{})
	// case "SmartGroupMobileDeviceMembershipChange":
	// 	return generateEvent(data, &SmartGroupMobileDeviceMembershipChangeEvent{})
	// case "SmartGroupUserMembershipChange":
	// 	return generateEvent(data, &SmartGroupUserMembershipChangeEvent{})
	default:
		return nil, &newEventError{"Not a recognized event type"}
	}
}

func (jwh *JamfWebhook) SetURLQueries(urlQueries []string) {
	jwh.URLQueries = urlQueries
}
