// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package analytics

// AnalyticsLogData: The data within Firebase Analytics log events.
type AnalyticsLogData struct {
	EventDim []EventDimensions `json:"eventDim,omitempty"` // A repeated record of event related dimensions.
	UserDim  *UserDim          `json:"userDim,omitempty"`  // User related dimensions.
}

// EventDimensions: Message containing information pertaining to the event.
type EventDimensions struct {
	Date                    *string                   `json:"date,omitempty"`                           // The date on which this event was logged.; (YYYYMMDD format in the registered timezone of your app.)
	Name                    *string                   `json:"name,omitempty"`                           // The name of this event.
	Params                  map[string]AnalyticsValue `json:"params,omitempty"`                         // A repeated record of the parameters associated with this event.
	PreviousTimestampMicros *int64                    `json:"previousTimestampMicros,string,omitempty"` // UTC client time when the previous event happened.
	TimestampMicros         *int64                    `json:"timestampMicros,string,omitempty"`         // UTC client time when the event happened.
	ValueInUsd              *float64                  `json:"valueInUsd,omitempty"`                     // Value param in USD.
}

// AnalyticsValue: Value for Event Params and UserProperty can be of type string or int or
// float or double.
type AnalyticsValue struct {
	DoubleValue *float64 `json:"doubleValue,omitempty"`
	FloatValue  *float64 `json:"floatValue,omitempty"`
	IntValue    *int64   `json:"intValue,string,omitempty"`
	StringValue *string  `json:"stringValue,omitempty"`
}

// UserDim: User related dimensions.
//
// Message containing information about the user associated with the event.
type UserDim struct {
	AppInfo                  *AppInfo                     `json:"appInfo,omitempty"`                         // App information.
	BundleInfo               *BundleInfo                  `json:"bundleInfo,omitempty"`                      // Information regarding the bundle in which these events were uploaded.
	DeviceInfo               *DeviceInfo                  `json:"deviceInfo,omitempty"`                      // Device information.
	FirstOpenTimestampMicros *int64                       `json:"firstOpenTimestampMicros,string,omitempty"` // The time (in microseconds) at which the user first opened the app.
	GeoInfo                  *GeoInfo                     `json:"geoInfo,omitempty"`                         // User's geographic information.
	LtvInfo                  *LtvInfo                     `json:"ltvInfo,omitempty"`                         // Lifetime Value information about this user.
	TrafficSource            *TrafficSource               `json:"trafficSource,omitempty"`                   // Information about marketing campaign which acquired the user.
	UserID                   *string                      `json:"userId,omitempty"`                          // The user ID set via the setUserId API.
	UserProperties           map[string]UserPropertyValue `json:"userProperties,omitempty"`                  // A repeated record of user properties set with the setUserProperty API.; https://firebase.google.com/docs/analytics/android/properties
}

// AppInfo: App information.
//
// Message which contains App Information.
type AppInfo struct {
	AppID         *string `json:"appId,omitempty"`         // Unique application identifier within an app store.
	AppInstanceID *string `json:"appInstanceId,omitempty"` // Unique id for this instance of the app.; Example: "71683BF9FA3B4B0D9535A1F05188BAF3"
	AppPlatform   *string `json:"appPlatform,omitempty"`   // The app platform.; Eg "ANDROID", "IOS".
	AppStore      *string `json:"appStore,omitempty"`      // The identifier of the store that installed the app.; Eg. "com.sec.android.app.samsungapps", "com.amazon.venezia",; "com.nokia.nstore"
	AppVersion    *string `json:"appVersion,omitempty"`    // The app's version name; Examples: "1.0", "4.3.1.1.213361", "2.3 (1824253)", "v1.8b22p6"
}

// BundleInfo: Information regarding the bundle in which these events were uploaded.
//
// Message containing information regarding the bundle in which these
// events were uploaded.
type BundleInfo struct {
	BundleSequenceID            *int64 `json:"bundleSequenceId,string,omitempty"`            // Monotonically increasing index for each bundle set by SDK.
	ServerTimestampOffsetMicros *int64 `json:"serverTimestampOffsetMicros,string,omitempty"` // Timestamp offset between collection time and upload time.
}

// DeviceInfo: Device information.
//
// Message containing device informations.
type DeviceInfo struct {
	DeviceCategory              *string `json:"deviceCategory,omitempty"`                     // Device category.; Eg. tablet or mobile.
	DeviceID                    *string `json:"deviceId,omitempty"`                           // Vendor specific device identifier. This is IDFV on iOS. Not used for; Android.; Example: "599F9C00-92DC-4B5C-9464-7971F01F8370"
	DeviceModel                 *string `json:"deviceModel,omitempty"`                        // Device model.; Eg. GT-I9192
	DeviceTimeZoneOffsetSeconds *int64  `json:"deviceTimeZoneOffsetSeconds,string,omitempty"` // The timezone of the device when data was uploaded as seconds skew from UTC.
	LimitedAdTracking           *bool   `json:"limitedAdTracking,omitempty"`                  // The device's Limit Ad Tracking setting.; When true, we cannot use device_id for remarketing, demographics or; influencing ads serving behaviour. However, we can use device_id for; conversion tracking and campaign attribution.
	MobileBrandName             *string `json:"mobileBrandName,omitempty"`                    // Device brand name.; Eg. Samsung, HTC, etc.
	MobileMarketingName         *string `json:"mobileMarketingName,omitempty"`                // Device marketing name.; Eg. Galaxy S4 Mini
	MobileModelName             *string `json:"mobileModelName,omitempty"`                    // Device model name.; Eg. GT-I9192
	PlatformVersion             *string `json:"platformVersion,omitempty"`                    // Device OS version when data capture ended.; Eg. 4.4.2
	ResettableDeviceID          *string `json:"resettableDeviceId,omitempty"`                 // The type of the resettable_device_id is always IDFA on iOS and AdId; on Android.; Example: "71683BF9-FA3B-4B0D-9535-A1F05188BAF3"
	UserDefaultLanguage         *string `json:"userDefaultLanguage,omitempty"`                // The user language.; Eg. "en-us", "en-za", "zh-tw", "jp"
}

// GeoInfo: User's geographic information.
//
// User's geographic informaiton.
type GeoInfo struct {
	City      *string `json:"city,omitempty"`      // The geographic city.; Eg. Sao Paulo
	Continent *string `json:"continent,omitempty"` // The geographic continent.; Eg. Americas
	Country   *string `json:"country,omitempty"`   // The geographic country.; Eg. Brazil
	Region    *string `json:"region,omitempty"`    // The geographic region.; Eg. State of Sao Paulo
}

// LtvInfo: Lifetime Value information about this user.
type LtvInfo struct {
	Currency *string  `json:"currency,omitempty"` // The currency corresponding to the revenue.
	Revenue  *float64 `json:"revenue,omitempty"`  // The Lifetime Value revenue of this user.
}

// TrafficSource: Information about marketing campaign which acquired the user.
//
// Mesage containing marketing campaign information which acquired the user.
type TrafficSource struct {
	UserAcquiredCampaign *string `json:"userAcquiredCampaign,omitempty"` // The name of the campaign which acquired the user.
	UserAcquiredMedium   *string `json:"userAcquiredMedium,omitempty"`   // The name of the medium which acquired the user.
	UserAcquiredSource   *string `json:"userAcquiredSource,omitempty"`   // The name of the network which acquired the user.
}

// UserPropertyValue: Predefined (eg: LTV) or custom properties (eg: birthday) stored on client
// side and associated with subsequent HitBundles.
type UserPropertyValue struct {
	Index            *int64 `json:"index,string,omitempty"`            // Index for user property (one-based).
	SetTimestampUsec *int64 `json:"setTimestampUsec,string,omitempty"` // UTC client time when user property was last set.
	Value            *Value `json:"value,omitempty"`                   // Last set value of user property.
}

// Value: Last set value of user property.
//
// Value for Event Params and UserProperty can be of type string or int or
// float or double.
type Value struct {
	DoubleValue *float64 `json:"doubleValue,omitempty"`
	FloatValue  *float64 `json:"floatValue,omitempty"`
	IntValue    *int64   `json:"intValue,string,omitempty"`
	StringValue *string  `json:"stringValue,omitempty"`
}
