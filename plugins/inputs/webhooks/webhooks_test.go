package webhooks

import (
	"reflect"
	"testing"

	"github.com/smithjw/telegraf-input-jamf-webhooks/plugins/inputs/webhooks"
)

func TestAvailableWebhooks(t *testing.T) {
	wb := NewWebhooks()
	expected := make([]Webhook, 0)
	if !reflect.DeepEqual(wb.AvailableWebhooks(), expected) {
		t.Errorf("expected to %v.\nGot %v", expected, wb.AvailableWebhooks())
	}

	wb.Jamf = &jamf.JamfWebhook{Path: "/jamf"}
	expected = append(expected, wb.Jamf)
	if !reflect.DeepEqual(wb.AvailableWebhooks(), expected) {
		t.Errorf("expected to be %v.\nGot %v", expected, wb.AvailableWebhooks())
	}
}
