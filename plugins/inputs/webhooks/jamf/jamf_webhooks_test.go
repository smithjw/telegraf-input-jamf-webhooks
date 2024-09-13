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

// Webhooks to add:
//   ComputerAdded
//   ComputerCheckIn
//   ComputerPatchPolicyCompleted
//   ComputerPushCapabilityChanged
//   DeviceAddedToDEP
//   JSSShutdown
//   JSSStartup
//   MobileDeviceCheckIn
//   MobileDeviceCommandCompleted
//   MobileDeviceEnrolled
//   MobileDeviceInventoryCompleted
//   MobileDevicePushSent
//   MobileDeviceUnenrolled
//   PatchSoftwareTitleUpdated
//   PushSent
//   RestAPIOperation
//   SCEPChallenge
//   SmartGroupComputerMembershipChange
//   SmartGroupMobileDeviceMembershipChange
//   SmartGroupUserMembershipChange
