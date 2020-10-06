// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    logEntryData, err := UnmarshalLogEntryData(bytes)
//    bytes, err = logEntryData.Marshal()

package auditv1

import "encoding/json"

func UnmarshalLogEntryData(data []byte) (LogEntryData, error) {
	var r LogEntryData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *LogEntryData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// This event is triggered when a new audit log entry is written.
type LogEntryData struct {
	InsertID         *string                `json:"insertId,omitempty"`        // A unique identifier for the log entry.
	Labels           map[string]interface{} `json:"labels,omitempty"`          // A set of user-defined (key, value) data that provides additional information about the; log entry.
	LogName          *string                `json:"logName,omitempty"`         // The resource name of the log to which this log entry belongs.
	Operation        *LogEntryOperation     `json:"operation,omitempty"`       // Information about an operation associated with the log entry, if applicable.
	ProtoPayload     *AuditLog              `json:"protoPayload,omitempty"`    // The log entry payload, which is always an AuditLog for Cloud Audit Log events.
	ReceiveTimestamp *string                `json:"receiveTimestamp,omitempty"`// The time the log entry was received by Logging.
	Resource         *MonitoredResource     `json:"resource,omitempty"`        // The monitored resource that produced this log entry. Example: a log entry that reports a; database error would be associated with the monitored resource designating the particular; database that reported the error.
	Severity         *string                `json:"severity,omitempty"`        // The severity of the log entry.
	SpanID           *string                `json:"spanId,omitempty"`          // The span ID within the trace associated with the log entry, if any. For Trace spans, this; is the same format that the Trace API v2 uses: a 16-character hexadecimal encoding of an; 8-byte array, such as `000000000000004a`.
	Timestamp        *string                `json:"timestamp,omitempty"`       // The time the event described by the log entry occurred.
	Trace            *string                `json:"trace,omitempty"`           // Resource name of the trace associated with the log entry, if any. If it contains a; relative resource name, the name is assumed to be relative to `//tracing.googleapis.com`.; Example: `projects/my-projectid/traces/06796866738c859f2f19b7cfb3214824`
}

// Information about an operation associated with the log entry, if applicable.
//
// Additional information about a potentially long-running operation with which a log entry
// is associated.
type LogEntryOperation struct {
	First    *bool   `json:"first,omitempty"`   // True if this is the first log entry in the operation.
	ID       *string `json:"id,omitempty"`      // An arbitrary operation identifier. Log entries with the same identifier are assumed to be; part of the same operation.
	Last     *bool   `json:"last,omitempty"`    // True if this is the last log entry in the operation.
	Producer *string `json:"producer,omitempty"`// An arbitrary producer identifier. The combination of `id` and `producer` must be globally; unique. Examples for `producer`: `"MyDivision.MyBigCompany.com"`,; `"github.com/MyProject/MyApplication"`.
}

// The log entry payload, which is always an AuditLog for Cloud Audit Log events.
type AuditLog struct {
	AuthenticationInfo    *AuthenticationInfo    `json:"authenticationInfo,omitempty"`   // Authentication information.
	AuthorizationInfo     []AuthorizationInfo    `json:"authorizationInfo,omitempty"`    // Authorization information. If there are multiple resources or permissions involved, then; there is one AuthorizationInfo element for each {resource, permission} tuple.
	Metadata              map[string]interface{} `json:"metadata,omitempty"`             // Other service-specific data about the request, response, and other information associated; with the current audited event.
	MethodName            *string                `json:"methodName,omitempty"`           // The name of the service method or operation. For example; "google.datastore.v1.Datastore.RunQuery"
	NumResponseItems      *int64                 `json:"numResponseItems,omitempty"`     // The number of items returned from a List or Query API method, if applicable.
	Request               map[string]interface{} `json:"request,omitempty"`              // The operation request. This may not include all request parameters, such as those that; are too large, privacy-sensitive, or duplicated elsewhere in the log record. It should; never include user-generated data, such as file contents. When the JSON object; represented here has a proto equivalent, the proto name will be indicated in the `@type`; property.
	RequestMetadata       *RequestMetadata       `json:"requestMetadata,omitempty"`      // Metadata about the operation.
	ResourceLocation      *ResourceLocation      `json:"resourceLocation,omitempty"`     // The resource location information.
	ResourceName          *string                `json:"resourceName,omitempty"`         // The resource or collection that is the target of the operation. For example; "shelves/SHELF_ID/books"
	ResourceOriginalState map[string]interface{} `json:"resourceOriginalState,omitempty"`// The resource's original state before mutation.
	Response              map[string]interface{} `json:"response,omitempty"`             // The operation response. This may not include all response elements, such as those that; are too large, privacy-sensitive, or duplicated elsewhere in the log record. It should; never include user-generated data, such as file contents. When the JSON object; represented here has a proto equivalent, the proto name will be indicated in the `@type`; property.
	ServiceData           map[string]interface{} `json:"serviceData,omitempty"`          // Deprecated, use `metadata` field instead. Other service-specific data about the request,; response, and other activities. When the JSON object represented here has a proto; equivalent, the proto name will be indicated in the `@type` property.
	ServiceName           *string                `json:"serviceName,omitempty"`          // The name of the API service performing the operation. For example,; `"datastore.googleapis.com"`.
	Status                *Status                `json:"status,omitempty"`               // The status of the overall operation.
}

// Authentication information.
//
// Authentication information for the operation.
type AuthenticationInfo struct {
	AuthoritySelector            *string                        `json:"authoritySelector,omitempty"`           // The authority selector specified by the requestor, if any. It is not guaranteed that the; principal was allowed to use this authority.
	PrincipalEmail               *string                        `json:"principalEmail,omitempty"`              // The email address of the authenticated user (or service account on behalf of third party; principal) making the request. For privacy reasons, the principal email address is; redacted for all read-only operations that fail with a "permission denied" error.
	PrincipalSubject             *string                        `json:"principalSubject,omitempty"`            // String representation of identity of requesting party. Populated for both first and third; party identities.
	ServiceAccountDelegationInfo []ServiceAccountDelegationInfo `json:"serviceAccountDelegationInfo,omitempty"`// Identity delegation history of an authenticated service account that makes the request.; It contains information on the real authorities that try to access GCP resources by; delegating on a service account. When multiple authorities present, they are guaranteed; to be sorted based on the original ordering of the identity delegation events.
	ServiceAccountKeyName        *string                        `json:"serviceAccountKeyName,omitempty"`       // The name of the service account key used to create or exchange credentials for; authenticating the service account making the request. This is a scheme-less URI full; resource name.
	ThirdPartyPrincipal          map[string]interface{}         `json:"thirdPartyPrincipal,omitempty"`         // The third party identification (if any) of the authenticated user making the request.; When the JSON object represented here has a proto equivalent, the proto name will be; indicated in the @type property.
}

// Identity delegation history of an authenticated service account
type ServiceAccountDelegationInfo struct {
	PrincipalEmail   *string                `json:"principalEmail,omitempty"`  // The email address of a Google account.
	ServiceMetadata  map[string]interface{} `json:"serviceMetadata,omitempty"` // Metadata about the service that uses the service account.
	ThirdPartyClaims map[string]interface{} `json:"thirdPartyClaims,omitempty"`// Metadata about third party identity.
}

// Authorization information. If there are multiple resources or permissions involved, then
// there is one AuthorizationInfo element for each {resource, permission} tuple.
type AuthorizationInfo struct {
	Granted            *bool     `json:"granted,omitempty"`           // Whether or not authorization for resource and permission was granted.
	Permission         *string   `json:"permission,omitempty"`        // The required IAM permission.
	Resource           *string   `json:"resource,omitempty"`          // The resource being accessed, as a REST-style string.
	ResourceAttributes *Resource `json:"resourceAttributes,omitempty"`// Resource attributes used in IAM condition evaluation. This field contains resource; attributes like resource type and resource name. To get the whole view of the attributes; used in IAM condition evaluation, the user must also look into; AuditLog.requestMetadata.requestAttributes.
}

// Resource attributes used in IAM condition evaluation. This field contains resource
// attributes like resource type and resource name. To get the whole view of the attributes
// used in IAM condition evaluation, the user must also look into
// AuditLog.requestMetadata.requestAttributes.
type Resource struct {
	Labels  map[string]interface{} `json:"labels,omitempty"` 
	Name    *string                `json:"name,omitempty"`   
	Service *string                `json:"service,omitempty"`
	Type    *string                `json:"type,omitempty"`   
}

// Metadata about the operation.
type RequestMetadata struct {
	CallerIP                *string  `json:"callerIp,omitempty"`               // The IP address of the caller. For caller from internet, this will be public IPv4 or IPv6; address. For caller from a Compute Engine VM with external IP address, this will be the; VM's external IP address. For caller from a Compute Engine VM without external IP; address, if the VM is in the same organization (or project) as the accessed resource,; `callerIp` will be the VM's internal IPv4 address, otherwise the `callerIp` will be; redacted to "gce-internal-ip". See https://cloud.google.com/compute/docs/vpc/ for more; information."
	CallerNetwork           *string  `json:"callerNetwork,omitempty"`          // The network of the caller.
	CallerSuppliedUserAgent *string  `json:"callerSuppliedUserAgent,omitempty"`// The user agent of the caller. This information is not authenticated and should be treated; accordingly.
	DestinationAttributes   *Peer    `json:"destinationAttributes,omitempty"`  // The destination of a network activity, such as accepting a TCP connection.
	RequestAttributes       *Request `json:"requestAttributes,omitempty"`      // Request attributes used in IAM condition evaluation. This field contains request; attributes like request time and access levels associated with the request.
}

// The destination of a network activity, such as accepting a TCP connection.
type Peer struct {
	IP         *string                `json:"ip,omitempty"`        
	Labels     map[string]interface{} `json:"labels,omitempty"`    
	Port       *int64                 `json:"port,omitempty"`      
	Principal  *string                `json:"principal,omitempty"` 
	RegionCode *string                `json:"regionCode,omitempty"`
}

// Request attributes used in IAM condition evaluation. This field contains request
// attributes like request time and access levels associated with the request.
type Request struct {
	Auth     *Auth                  `json:"auth,omitempty"`    
	Headers  map[string]interface{} `json:"headers,omitempty"` 
	Host     *string                `json:"host,omitempty"`    
	ID       *string                `json:"id,omitempty"`      
	Method   *string                `json:"method,omitempty"`  
	Path     *string                `json:"path,omitempty"`    
	Protocol *string                `json:"protocol,omitempty"`
	Query    *string                `json:"query,omitempty"`   
	Reason   *string                `json:"reason,omitempty"`  
	Scheme   *string                `json:"scheme,omitempty"`  
	Size     *int64                 `json:"size,omitempty"`    
	Time     *string                `json:"time,omitempty"`    
}

type Auth struct {
	AccessLevels []string               `json:"accessLevels,omitempty"`
	Audiences    []string               `json:"audiences,omitempty"`   
	Claims       map[string]interface{} `json:"claims,omitempty"`      
	Presenter    *string                `json:"presenter,omitempty"`   
	Principal    *string                `json:"principal,omitempty"`   
}

// The resource location information.
//
// Location information about a resource.
type ResourceLocation struct {
	CurrentLocations  []string `json:"currentLocations,omitempty"` // The locations of a resource after the execution of the operation. Requests to create or; delete a location based resource must populate the 'currentLocations' field and not the; 'originalLocations' field.
	OriginalLocations []string `json:"originalLocations,omitempty"`// The locations of a resource prior to the execution of the operation. Requests that mutate; the resource's location must populate both the 'originalLocations' as well as the; 'currentLocations' fields. For example:
}

// The status of the overall operation.
type Status struct {
	Code    *int64      `json:"code,omitempty"`   // The status code, which should be an enum value of [google.rpc.Code][google.rpc.Code].
	Details interface{} `json:"details"`          // A list of messages that carry the error details.  There is a common set of message types; for APIs to use.
	Message *string     `json:"message,omitempty"`// A developer-facing error message, which should be in English. Any user-facing error; message should be localized and sent in the; [google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client.
}

// The monitored resource that produced this log entry. Example: a log entry that reports a
// database error would be associated with the monitored resource designating the particular
// database that reported the error.
//
// The monitored resource that produced this log entry.
type MonitoredResource struct {
	Labels map[string]interface{} `json:"labels,omitempty"`// Values for all of the labels listed in the associated monitored resource descriptor. For; example, Compute Engine VM instances use the labels `"projectId"`, `"instanceId"`, and; `"zone"`.
	Type   *string                `json:"type,omitempty"`  // Required. The monitored resource type. For example, the type of a Compute Engine VM; instance is `gceInstance`.
}
