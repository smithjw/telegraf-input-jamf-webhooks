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
	ManagementID        string `json:"managementId"`
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
}

type MobileDevice struct {
	BluetoothMacAddress string `json:"bluetoothMacAddress"`
	DeviceName          string `json:"deviceName"`
	IcciID              string `json:"icciID"`
	Imei                string `json:"imei"`
	IPAddress           string `json:"ipAddress"`
	JssID               int    `json:"jssID"`
	ManagementID        string `json:"managementId"`
	Model               string `json:"model"`
	ModelDisplay        string `json:"modelDisplay"`
	OSBuild             string `json:"osBuild"`
	OSVersion           string `json:"osVersion"`
	Product             string `json:"product"`
	Room                string `json:"room"`
	SerialNumber        string `json:"serialNumber"`
	UDID                string `json:"udid"`
	UserDirectoryID     string `json:"userDirectoryID"`
	Username            string `json:"username"`
	Version             string `json:"version"`
	WifiMacAddress      string `json:"wifiMacAddress"`
}

// This struct/function is used across the following events:
//   - ComputerAdded
//   - ComputerInventoryCompleted
//   - ComputerPushCapabilityChanged
type ComputerEvent struct {
	Event   Computer `json:"event"`
	Webhook Webhook  `json:"webhook"`
}

func (s ComputerEvent) NewMetric() telegraf.Metric {
	tags := map[string]string{
		"event":         s.Webhook.Event,
		"jss_id":        strconv.Itoa(s.Event.JssID),
		"management_id": s.Event.ManagementID,
		"serial_number": s.Event.SerialNumber,
		"username":      s.Event.Username,
	}
	fields := map[string]interface{}{
		"alternate_mac_address": s.Event.AlternateMacAddress,
		"building":              s.Event.Building,
		"department":            s.Event.Department,
		"device_mac_address":    s.Event.MacAddress,
		"device_model":          s.Event.Model,
		"device_name":           s.Event.DeviceName,
		"device_udid":           s.Event.UDID,
		"email_address":         s.Event.EmailAddress,
		"ip_address":            s.Event.IPAddress,
		"os_buid":               s.Event.OSBuild,
		"os_version":            s.Event.OSVersion,
		"phone":                 s.Event.Phone,
		"position":              s.Event.Position,
		"realname":              s.Event.RealName,
		"reported_ip_address":   s.Event.ReportedIPAddress,
		"room":                  s.Event.Room,
		"user_directory_id":     s.Event.UserDirectoryID,
		"user_real_name":        s.Event.RealName,
	}
	m := metric.New(measurement, tags, fields, time.UnixMilli(s.Webhook.Timestamp))
	return m
}

type ComputerPolicyFinishedEvent struct {
	Event struct {
		Computer   Computer `json:"computer"`
		PolicyID   int      `json:"policyId"`
		Successful bool     `json:"successful"`
	} `json:"event"`
	Webhook Webhook `json:"webhook"`
}

func (s ComputerPolicyFinishedEvent) NewMetric() telegraf.Metric {
	tags := map[string]string{
		"event":         s.Webhook.Event,
		"policy_id":     strconv.Itoa(s.Event.PolicyID),
		"serial_number": s.Event.Computer.SerialNumber,
		"management_id": s.Event.Computer.ManagementID,
		"jss_id":        strconv.Itoa(s.Event.Computer.JssID),
	}
	fields := map[string]interface{}{
		"successful":  s.Event.Successful,
		"device_name": s.Event.Computer.DeviceName,
	}
	m := metric.New(measurement, tags, fields, time.UnixMilli(s.Webhook.Timestamp))
	return m
}

type ComputerCheckInEvent struct {
	Event struct {
		Computer Computer `json:"computer"`
		Trigger  string   `json:"trigger"`
		Username string   `json:"username"`
	} `json:"event"`
	Webhook Webhook `json:"webhook"`
}

func (s ComputerCheckInEvent) NewMetric() telegraf.Metric {
	tags := map[string]string{
		"event":         s.Webhook.Event,
		"serial_number": s.Event.Computer.SerialNumber,
		"management_id": s.Event.Computer.ManagementID,
		"jss_id":        strconv.Itoa(s.Event.Computer.JssID),
	}
	fields := map[string]interface{}{
		"check_in_username": s.Event.Username,
		"check_in_trigger":  s.Event.Trigger,
		"username":          s.Event.Computer.Username,
		"device_name":       s.Event.Computer.DeviceName,
	}
	m := metric.New(measurement, tags, fields, time.UnixMilli(s.Webhook.Timestamp))
	return m
}

type DeviceAddedToDEPEvent struct {
	Event struct {
		AssetTag                          string `json:"assetTag"`
		Description                       string `json:"description"`
		DeviceAssignedDate                int    `json:"deviceAssignedDate"`
		DeviceEnrollmentProgramInstanceID int    `json:"deviceEnrollmentProgramInstanceId"`
		Model                             string `json:"model"`
		SerialNumber                      string `json:"serialNumber"`
	} `json:"event"`
	Webhook Webhook `json:"webhook"`
}

func (s DeviceAddedToDEPEvent) NewMetric() telegraf.Metric {
	tags := map[string]string{
		"event":           s.Webhook.Event,
		"serial_number":   s.Event.SerialNumber,
		"asset_tag":       s.Event.AssetTag,
		"ade_instance_id": strconv.Itoa(s.Event.DeviceEnrollmentProgramInstanceID),
	}
	fields := map[string]interface{}{
		"description":   s.Event.Description,
		"assigned_date": s.Event.DeviceAssignedDate,
		"model":         s.Event.Model,
	}
	m := metric.New(measurement, tags, fields, time.UnixMilli(s.Webhook.Timestamp))
	return m
}

// type PatchPolicyAction struct {
// 	Action []string `json:"action"`
// }

// type ComputerPatchPolicyCompletedEvent struct {
// 	Computer        Computer          `json:"computer"`
// 	DeployedVersion string            `json:"deployedVersion"`
// 	EventActions    PatchPolicyAction `json:"eventActions"`
// 	PatchPolicyID   int               `json:"patchPolicyId"`
// 	PatchPolicyName string            `json:"patchPolicyName"`
// 	SoftwareTitleID int               `json:"softwareTitleId"`
// 	Successful      bool              `json:"successful"`
// 	Webhook         Webhook           `json:"webhook"`
// }

// This struct/function is used across the following events:
//   - MobileDeviceCheckIn
//   - MobileDeviceCommandCompleted
//   - MobileDeviceEnrolled
//   - MobileDeviceInventoryCompleted
//   - MobileDevicePushSent
//   - MobileDeviceUnEnrolled
type MobileDeviceEvent struct {
	Event   MobileDevice `json:"mobileDevice"`
	Webhook Webhook      `json:"webhook"`
}

func (s MobileDeviceEvent) NewMetric() telegraf.Metric {
	tags := map[string]string{
		"event":         s.Webhook.Event,
		"jss_id":        strconv.Itoa(s.Event.JssID),
		"serial_number": s.Event.SerialNumber,
		"management_id": s.Event.ManagementID,
		"username":      s.Event.Username,
	}
	fields := map[string]interface{}{
		"device_model_display": s.Event.ModelDisplay,
		"device_model":         s.Event.Model,
		"device_name":          s.Event.DeviceName,
		"device_udid":          s.Event.UDID,
		"icci_id":              s.Event.IcciID,
		"imei":                 s.Event.Imei,
		"ip_address":           s.Event.IPAddress,
		"os_build":             s.Event.OSBuild,
		"os_version":           s.Event.OSVersion,
		"product":              s.Event.Product,
		"room":                 s.Event.Room,
		"user_directory_id":    s.Event.UserDirectoryID,
		"version":              s.Event.Version,
		"wifi_mac_address":     s.Event.WifiMacAddress,
	}
	m := metric.New(measurement, tags, fields, time.UnixMilli(s.Webhook.Timestamp))
	return m
}

// This struct/function is used across the following events:
//   - JSSShutdown
//   - JSSStartup
type JSSEvent struct {
	Event struct {
		HostAddress        string `json:"hostAddress"`
		Institution        string `json:"institution"`
		IsClusterMaster    bool   `json:"isClusterMaster"`
		JssURL             string `json:"jssUrl"`
		WebApplicationPath string `json:"webApplicationPath"`
	} `json:"event"`
	Webhook Webhook `json:"webhook"`
}

func (s JSSEvent) NewMetric() telegraf.Metric {
	tags := map[string]string{
		"event":   s.Webhook.Event,
		"jss_url": s.Event.JssURL,
	}
	fields := map[string]interface{}{
		"host_address":      s.Event.HostAddress,
		"is_cluster_master": s.Event.IsClusterMaster,
	}
	m := metric.New(measurement, tags, fields, time.UnixMilli(s.Webhook.Timestamp))
	return m
}

type PushSentEvent struct {
	Event struct {
		ManagementID int    `json:"managementId"`
		Type         string `json:"type"`
	} `json:"event"`
	Webhook Webhook `json:"webhook"`
}

func (s PushSentEvent) NewMetric() telegraf.Metric {
	tags := map[string]string{
		"event":         s.Webhook.Event,
		"management_id": strconv.Itoa(s.Event.ManagementID),
	}
	fields := map[string]interface{}{
		"type": s.Event.Type,
	}
	m := metric.New(measurement, tags, fields, time.UnixMilli(s.Webhook.Timestamp))
	return m
}
