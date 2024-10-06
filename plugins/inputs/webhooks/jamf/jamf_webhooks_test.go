package jamf

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/influxdata/telegraf/testutil"
)

func JamfWebhookRequest(t *testing.T, event, jsonString string) {
	var acc testutil.Accumulator
	jwh := &JamfWebhook{Path: "/jamf", acc: &acc, log: testutil.Logger{}}
	req, err := http.NewRequest("POST", "/jamf", strings.NewReader(jsonString))
	require.NoError(t, err)
	w := httptest.NewRecorder()
	jwh.eventHandler(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("POST "+event+" returned HTTP status code %v.\nExpected %v", w.Code, http.StatusOK)
	}
}

func JamfWebhookRequestWithQueryParams(t *testing.T, event, jsonString string, queryParams map[string]string, expectedStatus int) {
	var acc testutil.Accumulator
	jwh := &JamfWebhook{
		Path:       "/jamf",
		URLQueries: []string{"hostname", "other"},
		acc:        &acc,
		log:        testutil.Logger{},
	}

	queryStr := "?"
	for key, value := range queryParams {
		queryStr += key + "=" + value + "&"
	}
	queryStr = strings.TrimSuffix(queryStr, "&")
	req, err := http.NewRequest("POST", "/jamf"+queryStr, strings.NewReader(jsonString))
	require.NoError(t, err)
	w := httptest.NewRecorder()
	jwh.eventHandler(w, req)
	if w.Code != expectedStatus {
		t.Errorf("POST "+event+" returned HTTP status code %v.\nExpected %v", w.Code, expectedStatus)
	}
	for _, query := range jwh.URLQueries {
		if value, exists := queryParams[query]; exists && value != "" {
			if !acc.HasTag("jamf_webhooks", query) {
				t.Errorf("Expected tag '%s' with value '%s', but it was not found", query, value)
			}
		}
	}
}

func TestComputerPolicyFinishedEvent(t *testing.T) {
	JamfWebhookRequest(t, "ComputerPolicyFinished", ComputerPolicyFinishedJSON())
}

func TestComputerInventoryCompletedEvent(t *testing.T) {
	JamfWebhookRequest(t, "ComputerInventoryCompleted", ComputerInventoryCompletedJSON())
}

func TestDeviceAddedToDEPEvent(t *testing.T) {
	JamfWebhookRequest(t, "DeviceAddedToDEP", DeviceAddedToDEPJSON())
}

func TestComputerAddedEvent(t *testing.T) {
	JamfWebhookRequest(t, "ComputerAdded", ComputerAddedJSON())
}

// func TestSmartGroupComputerMembershipChangeEvent(t *testing.T) {
// 	JamfWebhookRequest(t, "SmartGroupComputerMembershipChange", SmartGroupComputerMembershipChangeJSON())
// }

func TestComputerCheckInEvent(t *testing.T) {
	JamfWebhookRequest(t, "ComputerCheckIn", ComputerCheckInJSON())
}

// func TestComputerPatchPolicyCompletedEvent(t *testing.T) {
// 	JamfWebhookRequest(t, "ComputerPatchPolicyCompleted", ComputerPatchPolicyCompletedJSON())
// }

func TestJSSShutdownEvent(t *testing.T) {
	JamfWebhookRequest(t, "JSSShutdown", JSSShutdownJSON())
}

func TestJSSStartupEvent(t *testing.T) {
	JamfWebhookRequest(t, "JSSStartup", JSSStartupJSON())
}

func TestComputerPushCapabilityChangedEvent(t *testing.T) {
	JamfWebhookRequest(t, "ComputerPushCapabilityChanged", ComputerPushCapabilityChangedJSON())
}

func TestMobileDeviceCheckInEvent(t *testing.T) {
	JamfWebhookRequest(t, "MobileDeviceCheckIn", MobileDeviceCheckInJSON())
}

func TestMobileDeviceCommandCompletedEvent(t *testing.T) {
	JamfWebhookRequest(t, "MobileDeviceCommandCompleted", MobileDeviceCommandCompletedJSON())
}

func TestMobileDeviceEnrolledEvent(t *testing.T) {
	JamfWebhookRequest(t, "MobileDeviceEnrolled", MobileDeviceEnrolledJSON())
}

func TestMobileDeviceInventoryCompletedEvent(t *testing.T) {
	JamfWebhookRequest(t, "MobileDeviceInventoryCompleted", MobileDeviceInventoryCompletedJSON())
}

func TestMobileDevicePushSentEvent(t *testing.T) {
	JamfWebhookRequest(t, "MobileDevicePushSent", MobileDevicePushSentJSON())
}

func TestMobileDeviceUnEnrolledEvent(t *testing.T) {
	JamfWebhookRequest(t, "MobileDeviceUnEnrolled", MobileDeviceUnEnrolledJSON())
}

// func TestPatchSoftwareTitleUpdatedEvent(t *testing.T) {
// 	JamfWebhookRequest(t, "PatchSoftwareTitleUpdated", PatchSoftwareTitleUpdatedJSON())
// }

func TestPushSentEvent(t *testing.T) {
	JamfWebhookRequest(t, "PushSent", PushSentJSON())
}

// func TestRestAPIOperationEvent(t *testing.T) {
// 	JamfWebhookRequest(t, "RestAPIOperation", RestAPIOperationJSON())
// }

// func TestSmartGroupMobileDeviceMembershipChangeEvent(t *testing.T) {
// 	JamfWebhookRequest(t, "SmartGroupMobileDeviceMembershipChange", SmartGroupMobileDeviceMembershipChangeJSON())
// }

// func TestSmartGroupUserMembershipChangeEvent(t *testing.T) {
// 	JamfWebhookRequest(t, "SmartGroupUserMembershipChange", SmartGroupUserMembershipChangeJSON())
// }

func TestJamfWebhookWithQueryParamsSuccess(t *testing.T) {
	queryParams := map[string]string{
		"hostname": "test1.example.com",
		"other":    "value2",
		"unused":   "test",
	}
	JamfWebhookRequestWithQueryParams(
		t,
		"ComputerInventoryCompleted",
		`{"webhook": {"webhookEvent": "ComputerInventoryCompleted"}}`, // JSON payload
		queryParams,
		http.StatusOK, // Expected status
	)
}
