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

// TODO:
func ComputerAddedJSON() string {
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
		"webhookEvent": "ComputerAdded"
	}
}`
}

// TODO:
func ComputerCheckInJSON() string {
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
		"trigger": "string",
		"username": "string"
	},
	"webhook": {
		"eventTimestamp": 1667601111,
		"id": 999,
		"name": "string",
		"webhookEvent": "ComputerCheckIn"
	}
}`
}

// TODO:
func ComputerPatchPolicyCompletedJSON() string {
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
		"deployedVersion": "string",
		"eventActions": {
			"action": [
				"string",
				"string"
			]
		},
		"patchPolicyId": 999,
		"patchPolicyName": "string",
		"softwareTitleId": 999,
		"successful": true
	},
	"webhook": {
		"eventTimestamp": 1667601111,
		"id": 999,
		"name": "string",
		"webhookEvent": "ComputerPatchPolicyCompleted"
	}
}`
}

// TODO:
func ComputerPushCapabilityChangedJSON() string {
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
		"webhookEvent": "ComputerPushCapabilityChanged"
	}
}`
}

// TODO:
func DeviceAddedToDEPJSON() string {
	return `{
	"event": {
		"assetTag": "string",
		"description": "string",
		"deviceAssignedDate": 999,
		"deviceEnrollmentProgramInstanceId": 999,
		"model": "string",
		"serialNumber": "string"
	},
	"webhook": {
		"eventTimestamp": 1667601111,
		"id": 999,
		"name": "string",
		"webhookEvent": "DeviceAddedToDEP"
	}
}`
}

// TODO:
func JSSShutdownJSON() string {
	return `{
	"event": {
		"hostAddress": "string",
		"institution": "string",
		"isClusterMaster": true,
		"jssUrl": "string",
		"webApplicationPath": "string"
	},
	"webhook": {
		"eventTimestamp": 1667601111,
		"id": 999,
		"name": "string",
		"webhookEvent": "JSSShutdown"
	}
}`
}

// TODO:
func JSSStartupJSON() string {
	return `{
	"event": {
		"hostAddress": "string",
		"institution": "string",
		"isClusterMaster": true,
		"jssUrl": "string",
		"webApplicationPath": "string"
	},
	"webhook": {
		"eventTimestamp": 1667601111,
		"id": 999,
		"name": "string",
		"webhookEvent": "JSSStartup"
	}
}`
}

// TODO:
func MobileDeviceCheckInJSON() string {
	return `{
	"event": {
		"bluetoothMacAddress": "string",
		"deviceName": "string",
		"icciID": "string",
		"imei": "string",
		"ipAddress": "string",
		"jssID": 999,
		"model": "string",
		"modelDisplay": "string",
		"osBuild": "string",
		"osVersion": "string",
		"product": "string",
		"room": "string",
		"serialNumber": "string",
		"udid": "string",
		"userDirectoryID": "string",
		"username": "string",
		"version": "string",
		"wifiMacAddress": "string",
		"managementId": "string"
	},
	"webhook": {
		"eventTimestamp": 1667601111,
		"id": 999,
		"name": "string",
		"webhookEvent": "MobileDeviceCheckIn"
	}
}`
}

// TODO:
func MobileDeviceCommandCompletedJSON() string {
	return `{
	"event": {
		"bluetoothMacAddress": "string",
		"deviceName": "string",
		"icciID": "string",
		"imei": "string",
		"ipAddress": "string",
		"jssID": 999,
		"model": "string",
		"modelDisplay": "string",
		"osBuild": "string",
		"osVersion": "string",
		"product": "string",
		"room": "string",
		"serialNumber": "string",
		"udid": "string",
		"userDirectoryID": "string",
		"username": "string",
		"version": "string",
		"wifiMacAddress": "string",
		"managementId": "string"
	},
	"webhook": {
		"eventTimestamp": 1667601111,
		"id": 999,
		"name": "string",
		"webhookEvent": "MobileDeviceCommandCompleted"
	}
}`
}

// TODO:
func MobileDeviceEnrolledJSON() string {
	return `{
	"event": {
		"bluetoothMacAddress": "string",
		"deviceName": "string",
		"icciID": "string",
		"imei": "string",
		"ipAddress": "string",
		"jssID": 999,
		"model": "string",
		"modelDisplay": "string",
		"osBuild": "string",
		"osVersion": "string",
		"product": "string",
		"room": "string",
		"serialNumber": "string",
		"udid": "string",
		"userDirectoryID": "string",
		"username": "string",
		"version": "string",
		"wifiMacAddress": "string",
		"managementId": "string"
	},
	"webhook": {
		"eventTimestamp": 1667601111,
		"id": 999,
		"name": "string",
		"webhookEvent": "MobileDeviceEnrolled"
	}
}`
}

// TODO:
func MobileDeviceInventoryCompletedJSON() string {
	return `{
	"event": {
		"udid": "string",
		"deviceName": "string",
		"version": "string",
		"model": "string",
		"bluetoothMacAddress": "string",
		"wifiMacAddress": "string",
		"imei": "string",
		"icciID": "string",
		"product": "string",
		"serialNumber": "string",
		"userDirectoryID": "string",
		"room": "string",
		"osVersion": "string",
		"osBuild": "string",
		"modelDisplay": "string",
		"username": "string",
		"jssID": 999,
		"ipAddress": "string",
		"managementId": "string"
	},
	"webhook": {
		"eventTimestamp": 1667601111,
		"id": 999,
		"name": "string",
		"webhookEvent": "MobileDeviceInventoryCompleted"
	}
}`
}

// TODO:
func MobileDevicePushSentJSON() string {
	return `{
	"event": {
		"bluetoothMacAddress": "string",
		"deviceName": "string",
		"icciID": "string",
		"imei": "string",
		"ipAddress": "string",
		"jssID": 999,
		"model": "string",
		"modelDisplay": "string",
		"osBuild": "string",
		"osVersion": "string",
		"product": "string",
		"room": "string",
		"serialNumber": "string",
		"udid": "string",
		"userDirectoryID": "string",
		"username": "string",
		"version": "string",
		"wifiMacAddress": "string",
		"managementId": "string"
	},
	"webhook": {
		"eventTimestamp": 1667601111,
		"id": 999,
		"name": "string",
		"webhookEvent": "MobileDevicePushSent"
	}
}`
}

// TODO:
func MobileDeviceUnenrolledJSON() string {
	return `{
	"event": {
		"bluetoothMacAddress": "string",
		"deviceName": "string",
		"icciID": "string",
		"imei": "string",
		"ipAddress": "string",
		"jssID": 999,
		"model": "string",
		"modelDisplay": "string",
		"osBuild": "string",
		"osVersion": "string",
		"product": "string",
		"room": "string",
		"serialNumber": "string",
		"udid": "string",
		"userDirectoryID": "string",
		"username": "string",
		"version": "string",
		"wifiMacAddress": "string",
		"managementId": "string"
	},
	"webhook": {
		"eventTimestamp": 1667601111,
		"id": 999,
		"name": "string",
		"webhookEvent": "MobileDeviceUnEnrolled"
	}
}`
}

// TODO:
func PatchSoftwareTitleUpdatedJSON() string {
	return `{
	"event": {
		"jssID": 999,
		"lastUpdate": 1667601111,
		"latestVersion": "string",
		"name": "string",
		"reportUrls": [
			"string",
			"string"
		]
	},
	"webhook": {
		"eventTimestamp": 1667601111,
		"id": 999,
		"name": "string",
		"webhookEvent": "PatchSoftwareTitleUpdated"
	}
}`
}

// TODO:
func PushSentJSON() string {
	return `{
	"event": {
		"managementId": 999,
		"type": "string"
	},
	"webhook": {
		"eventTimestamp": 1667601111,
		"id": 999,
		"name": "string",
		"webhookEvent": "PushSent"
	}
}`
}

// TODO:
func RestAPIOperationJSON() string {
	return `{
	"event": {
		"authorizedUsername": "string",
		"objectID": 999,
		"objectName": "string",
		"objectTypeName": "string",
		"operationSuccessful": true,
		"restAPIOperationType": "string"
	},
	"webhook": {
		"eventTimestamp": 1667601111,
		"id": 999,
		"name": "string",
		"webhookEvent": "RestAPIOperation"
	}
}`
}

// TODO:
func SmartGroupComputerMembershipChangeJSON() string {
	return `{
	"event": {
		"computer": true,
		"groupAddedDevices": [],
		"groupAddedDevicesIds": [],
		"groupRemovedDevices": [],
		"groupRemovedDevicesIds": [
			999
		],
		"jssid": 999,
		"name": "string",
		"smartGroup": true
	},
	"webhook": {
		"eventTimestamp": 1667601111,
		"id": 999,
		"name": "string",
		"webhookEvent": "SmartGroupComputerMembershipChange"
	}
}`
}

// TODO:
func SmartGroupMobileDeviceMembershipChangeJSON() string {
	return `{
	"event": {
		"computer": true,
		"groupAddedDevices": [
			{
				"bluetoothMacAddress": "string",
				"deviceName": "string",
				"icciID": "string",
				"imei": "string",
				"ipAddress": "string",
				"jssID": 999,
				"model": "string",
				"modelDisplay": "string",
				"osBuild": "string",
				"osVersion": "string",
				"product": "string",
				"room": "string",
				"serialNumber": "string",
				"udid": "string",
				"userDirectoryID": "string",
				"username": "string",
				"version": "string",
				"wifiMacAddress": "string"
			}
		],
		"groupAddedDevicesIds": [
			999
		],
		"groupRemovedDevices": [],
		"groupRemovedDevicesIds": [],
		"jssid": 999,
		"name": "string",
		"smartGroup": true
	},
	"webhook": {
		"eventTimestamp": 1667601111,
		"id": 999,
		"name": "string",
		"webhookEvent": "SmartGroupMobileDeviceMembershipChange"
	}
}`
}

// TODO:
func SmartGroupUserMembershipChangeJSON() string {
	return `{
	"event": {
		"groupAddedUserIds": [
			999
		],
		"groupRemovedUserIds": [],
		"jssid": 999,
		"name": "string",
		"smartGroup": true
	},
	"webhook": {
		"eventTimestamp": 1667601111,
		"id": 999,
		"name": "string",
		"webhookEvent": "SmartGroupUserMembershipChange"
	}
}`
}
