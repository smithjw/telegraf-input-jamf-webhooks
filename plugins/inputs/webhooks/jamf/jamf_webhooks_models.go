package jamf

import (
	"strconv"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/metric"
)

const measurement = "jamf_webhooks"

type Event interface {
	NewMetric() telegraf.Metric
}
type Webhook struct {
	Timestamp int64  `json:"eventTimestamp"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Event     string `json:"webhookEvent"`
}

type Computer struct {
	AlternateMacAddress string `json:"alternateMacAddress"`
	Building            string `json:"building"`
	Department          string `json:"department"`
	DeviceName          string `json:"deviceName"`
	EmailAddress        string `json:"emailAddress"`
	IPAddress           string `json:"ipAddress"`
	JssID               int    `json:"jssID"`
	MacAddress          string `json:"macAddress"`
	Model               string `json:"model"`
	OSBuild             string `json:"osBuild"`
	OSVersion           string `json:"osVersion"`
	Phone               string `json:"phone"`
	Position            string `json:"position"`
	RealName            string `json:"realName"`
	ReportedIPAddress   string `json:"reportedIpAddress"`
	Room                string `json:"room"`
	SerialNumber        string `json:"serialNumber"`
	UDID                string `json:"udid"`
	UserDirectoryID     string `json:"userDirectoryID"`
	Username            string `json:"username"`
	ManagementID        string `json:"managementId"`
}

type ComputerInventoryCompletedEvent struct {
	Event   Computer `json:"event"`
	Webhook Webhook  `json:"webhook"`
}

func (s ComputerInventoryCompletedEvent) NewMetric() telegraf.Metric {
	t := map[string]string{
		"event":         s.Webhook.Event,
		"jss_id":        strconv.Itoa(s.Event.JssID),
		"serial_number": s.Event.SerialNumber,
		"management_id": s.Event.ManagementID,
	}
	f := map[string]interface{}{
		"device_name":    s.Event.DeviceName,
		"os_version":     s.Event.OSVersion,
		"user_real_name": s.Event.RealName,
		"email_address":  s.Event.EmailAddress,
		"username":       s.Event.Username,
		"device_model":   s.Event.Model,
		"device_udid":    s.Event.UDID,
	}
	m := metric.New(measurement, t, f, time.Unix(s.Webhook.Timestamp, 0))
	return m
}

type ComputerPolicyEvent struct {
	Computer   Computer `json:"computer"`
	PolicyID   int      `json:"policyId"`
	Successful bool     `json:"successful"`
}

type ComputerPolicyFinishedEvent struct {
	Event   ComputerPolicyEvent `json:"event"`
	Webhook Webhook             `json:"webhook"`
}

func (s ComputerPolicyFinishedEvent) NewMetric() telegraf.Metric {
	t := map[string]string{
		"event":         s.Webhook.Event,
		"policy_id":     strconv.Itoa(s.Event.PolicyID),
		"serial_number": s.Event.Computer.SerialNumber,
		"management_id": s.Event.Computer.ManagementID,
		"jss_id":        strconv.Itoa(s.Event.Computer.JssID),
	}
	f := map[string]interface{}{
		"successful":  s.Event.Successful,
		"device_name": s.Event.Computer.DeviceName,
	}
	m := metric.New(measurement, t, f, time.Unix(s.Webhook.Timestamp, 0))
	return m
}
