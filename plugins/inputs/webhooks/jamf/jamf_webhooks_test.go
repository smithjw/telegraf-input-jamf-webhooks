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

func TestComputerPolicyFinishedEvent(t *testing.T) {
	JamfWebhookRequest(t, "ComputerPolicyFinished", ComputerPolicyFinishedJSON())
}

func TestComputerInventoryCompletedEvent(t *testing.T) {
	JamfWebhookRequest(t, "ComputerInventoryCompleted", ComputerInventoryCompletedJSON())
}

// To Do
// func TestDeviceAddedToDEPEvent(t *testing.T) {
// 	JamfWebhookRequest(t, "DeviceAddedToDEP", DeviceAddedToDEPJSON())
// }

// func TestComputerAddedEvent(t *testing.T) {
// 	JamfWebhookRequest(t, "ComputerAdded", ComputerAddedJSON())
// }

// func TestSmartGroupComputerMembershipChangeEvent(t *testing.T) {
// 	JamfWebhookRequest(t, "SmartGroupComputerMembershipChange", SmartGroupComputerMembershipChangeJSON())
// }

// func TestComputerCheckInEvent(t *testing.T) {
// 	JamfWebhookRequest(t, "ComputerCheckIn", ComputerCheckInJSON())
// }
// func TestComputerPatchPolicyCompletedEvent(t *testing.T) {
// 	JamfWebhookRequest(t, "ComputerPatchPolicyCompleted", ComputerPatchPolicyCompletedJSON())
// }

func TestJSSShutdownEvent(t *testing.T) {
	JamfWebhookRequest(t, "JSSShutdown", JSSShutdownJSON())
}

func TestJSSStartupEvent(t *testing.T) {
	JamfWebhookRequest(t, "JSSStartup", JSSStartupJSON())
}

// func TestComputerPushCapabilityChangedEvent(t *testing.T) {
// 	JamfWebhookRequest(t, "ComputerPushCapabilityChanged", ComputerPushCapabilityChangedJSON())
// }

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

func TestMobileDeviceUnenrolledEvent(t *testing.T) {
	JamfWebhookRequest(t, "MobileDeviceUnenrolled", MobileDeviceUnenrolledJSON())
}

// func TestPatchSoftwareTitleUpdatedEvent(t *testing.T) {
// 	JamfWebhookRequest(t, "PatchSoftwareTitleUpdated", PatchSoftwareTitleUpdatedJSON())
// }

// func TestPushSentEvent(t *testing.T) {
// 	JamfWebhookRequest(t, "PushSent", PushSentJSON())
// }

// func TestRestAPIOperationEvent(t *testing.T) {
// 	JamfWebhookRequest(t, "RestAPIOperation", RestAPIOperationJSON())
// }

// func TestSmartGroupMobileDeviceMembershipChangeEvent(t *testing.T) {
// 	JamfWebhookRequest(t, "SmartGroupMobileDeviceMembershipChange", SmartGroupMobileDeviceMembershipChangeJSON())
// }

// func TestSmartGroupUserMembershipChangeEvent(t *testing.T) {
// 	JamfWebhookRequest(t, "SmartGroupUserMembershipChange", SmartGroupUserMembershipChangeJSON())
// }
