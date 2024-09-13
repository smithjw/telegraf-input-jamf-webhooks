package jamf

func ComputerInventoryCompletedJSON() string {
	return `{
	"event": {
		"alternateMacAddress": "string",
		"building": "string",
		"department": "string",
		"deviceName": "string",
		"emailAddress": "string",
		"ipAddress": "string",
		"jssID": 999,
		"macAddress": "string",
		"model": "string",
		"osBuild": "string",
		"osVersion": "string",
		"phone": "string",
		"position": "string",
		"realName": "string",
		"reportedIpAddress": "string",
		"room": "string",
		"serialNumber": "string",
		"udid": "string",
		"userDirectoryID": "string",
		"username": "string",
		"managementId": "string"
	},
	"webhook": {
		"eventTimestamp": 1667601111,
		"id": 999,
		"name": "string",
		"webhookEvent": "ComputerInventoryCompleted"
	}
}`
}

func ComputerPolicyFinishedJSON() string {
	return `{
	"event": {
		"computer": {
			"alternateMacAddress": "string",
			"building": "string",
			"department": "string",
			"deviceName": "string",
			"emailAddress": "string",
			"ipAddress": "string",
			"jssID": 999,
			"macAddress": "string",
			"model": "string",
			"osBuild": "string",
			"osVersion": "string",
			"phone": "string",
			"position": "string",
			"realName": "string",
			"reportedIpAddress": "string",
			"room": "string",
			"serialNumber": "string",
			"udid": "string",
			"userDirectoryID": "string",
			"username": "string",
			"managementId": "string"
		},
		"policyId": 999,
		"successful": true
	},
	"webhook": {
		"eventTimestamp": 1667601111,
		"id": 999,
		"name": "string",
		"webhookEvent": "ComputerPolicyFinished"
	}
}`
}
