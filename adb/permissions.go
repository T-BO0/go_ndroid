package adb

type AndroidPermission int

const (
	StoragePermission AndroidPermission = iota
	CameraPermission
	SendNotificationPermission
	RecordAudioPermission
	PhonePermission
	ContactsPermission
	SmsPermission
	InternetPermission
	BluetoothPermission
	LocationBackgroundPermission
	AccessFineLocationPermission
	AccessCoarseLocationPermission
	AccessNetworkStatePermission
	AccessWifiStatePermission
	ChangeWifiStatePermission
	ChangeNetworkStatePermission
	AccessMockLocationPermission
	AccessLocationExtraCommandsPermission
	AccessCheckinPropertiesPermission
	AccessLocationPermission
	AccessCheckinPermission
	AccessCoarseUpdatesPermission
	AccessFineUpdatesPermission
)

// For those permissions you need root access
var PermissionStrings = map[AndroidPermission]string{
	StoragePermission:                     "android.permission.READ_EXTERNAL_STORAGE",      // Storage permission
	CameraPermission:                      "android.permission.CAMERA",                     // Camera permission
	SendNotificationPermission:            "android.permission.POST_NOTIFICATION",          // Send notification permission
	RecordAudioPermission:                 "android.permission.RECORD_AUDIO",               // Record audio permission
	PhonePermission:                       "android.permission.READ_PHONE_STATE",           // Phone permission
	ContactsPermission:                    "android.permission.READ_CONTACTS",              // Contacts permission
	SmsPermission:                         "android.permission.SEND_SMS",                   // SMS permission
	InternetPermission:                    "android.permission.INTERNET",                   // Internet permission
	BluetoothPermission:                   "android.permission.BLUETOOTH",                  // Bluetooth permission
	LocationBackgroundPermission:          "android.permission.ACCESS_BACKGROUND_LOCATION", // Location permission, background
	AccessFineLocationPermission:          "android.permission.ACCESS_FINE_LOCATION",       // Location permission, high accuracy
	AccessCoarseLocationPermission:        "android.permission.ACCESS_COARSE_LOCATION",     // Location permission, low accuracy
	AccessNetworkStatePermission:          "android.permission.ACCESS_NETWORK_STATE",       // Network state permission
	AccessWifiStatePermission:             "android.permission.ACCESS_WIFI_STATE",          // Wifi state permission
	ChangeWifiStatePermission:             "android.permission.CHANGE_WIFI_STATE",          // Change wifi state permission
	ChangeNetworkStatePermission:          "android.permission.CHANGE_NETWORK_STATE",       // Change network state permission
	AccessMockLocationPermission:          "android.permission.ACCESS_MOCK_LOCATION",       // Mock location permission
	AccessLocationExtraCommandsPermission: "android.permission.LOCATION_EXTRA_COMMANDS",    // Location extra commands permission
	AccessCheckinPropertiesPermission:     "android.permission.CHECKIN_PROPERTIES",         // Checkin properties permission
	AccessLocationPermission:              "android.permission.ACCESS_LOCATION",            // Access location permission
	AccessCheckinPermission:               "android.permission.ACCESS_CHECKIN_PROPERTIES",  // Access checkin permission
	AccessCoarseUpdatesPermission:         "android.permission.ACCESS_COARSE_UPDATES",      // Access coarse updates permission
	AccessFineUpdatesPermission:           "android.permission.ACCESS_FINE_UPDATES",        // Access fine updates permission
}
