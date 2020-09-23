// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    event, err := UnmarshalEvent(bytes)
//    bytes, err = event.Marshal()

package auditv1

import "bytes"
import "errors"
import "encoding/json"

func UnmarshalAuditLogWrittenEvent(data []byte) (AuditLogWrittenEvent, error) {
	var r AuditLogWrittenEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AuditLogWrittenEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}


type EventClass struct {
	AuditLogWrittenEvent *AuditLogWrittenEvent `json:"AuditLogWrittenEvent,omitempty"`// This event is triggered when a new audit log entry is written.
}

// This event is triggered when a new audit log entry is written.
type AuditLogWrittenEvent struct {
	InsertID         *string                `json:"insert_id,omitempty"`        
	Labels           map[string]interface{} `json:"labels,omitempty"`           
	LogName          *string                `json:"log_name,omitempty"`         
	Operation        *LogEntryOperation     `json:"operation,omitempty"`        
	ProtoPayload     *AuditLog              `json:"proto_payload,omitempty"`    
	ReceiveTimestamp *string                `json:"receive_timestamp,omitempty"`
	Resource         *MonitoredResource     `json:"resource,omitempty"`         
	Severity         *int64                 `json:"severity,omitempty"`         
	SpanID           *string                `json:"span_id,omitempty"`          
	Timestamp        *string                `json:"timestamp,omitempty"`        
	Trace            *string                `json:"trace,omitempty"`            
}

type LogEntryOperation struct {
	First    *bool   `json:"first,omitempty"`   
	ID       *string `json:"id,omitempty"`      
	Last     *bool   `json:"last,omitempty"`    
	Producer *string `json:"producer,omitempty"`
}

type AuditLog struct {
	AuthenticationInfo    *AuthenticationInfo    `json:"authentication_info,omitempty"`    
	AuthorizationInfo     []AuthorizationInfo    `json:"authorization_info,omitempty"`     
	Metadata              map[string]interface{} `json:"metadata,omitempty"`               
	MethodName            *string                `json:"method_name,omitempty"`            
	NumResponseItems      *int64                 `json:"num_response_items,omitempty"`     
	Request               map[string]interface{} `json:"request,omitempty"`                
	RequestMetadata       *RequestMetadata       `json:"request_metadata,omitempty"`       
	ResourceLocation      *ResourceLocation      `json:"resource_location,omitempty"`      
	ResourceName          *string                `json:"resource_name,omitempty"`          
	ResourceOriginalState map[string]interface{} `json:"resource_original_state,omitempty"`
	Response              map[string]interface{} `json:"response,omitempty"`               
	ServiceData           map[string]interface{} `json:"service_data,omitempty"`           
	ServiceName           *string                `json:"service_name,omitempty"`           
	Status                *Status                `json:"status,omitempty"`                 
}

// Authentication information for the operation.
type AuthenticationInfo struct {
	AuthoritySelector            *string                        `json:"authority_selector,omitempty"`             // The authority selector specified by the requestor, if any. It is not guaranteed that the; principal was allowed to use this authority.
	PrincipalEmail               *string                        `json:"principal_email,omitempty"`                // The email address of the authenticated user (or service account on behalf of third party; principal) making the request. For privacy reasons, the principal email address is; redacted for all read-only operations that fail with a "permission denied" error.
	PrincipalSubject             *string                        `json:"principal_subject,omitempty"`              // String representation of identity of requesting party. Populated for both first and third; party identities.
	ServiceAccountDelegationInfo []ServiceAccountDelegationInfo `json:"service_account_delegation_info,omitempty"`// Identity delegation history of an authenticated service account that makes the request.; It contains information on the real authorities that try to access GCP resources by; delegating on a service account. When multiple authorities present, they are guaranteed; to be sorted based on the original ordering of the identity delegation events.
	ServiceAccountKeyName        *string                        `json:"service_account_key_name,omitempty"`       // The name of the service account key used to create or exchange credentials for; authenticating the service account making the request. This is a scheme-less URI full; resource name.
	ThirdPartyPrincipal          map[string]interface{}         `json:"third_party_principal,omitempty"`          // The third party identification (if any) of the authenticated user making the request.; When the JSON object represented here has a proto equivalent, the proto name will be; indicated in the @type property.
}

// Identity delegation history of an authenticated service account
type ServiceAccountDelegationInfo struct {
	PrincipalEmail   *string                `json:"principal_email,omitempty"`   
	ServiceMetadata  map[string]interface{} `json:"service_metadata,omitempty"`  
	ThirdPartyClaims map[string]interface{} `json:"third_party_claims,omitempty"`
}

type AuthorizationInfo struct {
	Granted            *bool     `json:"granted,omitempty"`            // Whether or not authorization for resource and permission was granted.
	Permission         *string   `json:"permission,omitempty"`         // The required IAM permission.
	Resource           *string   `json:"resource,omitempty"`           // The resource being accessed, as a REST-style string.
	ResourceAttributes *Resource `json:"resource_attributes,omitempty"`// Resource attributes used in IAM condition evaluation. This field contains resource; attributes like resource type and resource name. To get the whole view of the attributes; used in IAM condition evaluation, the user must also look into; AuditLog.request_metadata.request_attributes.
}

// Resource attributes used in IAM condition evaluation. This field contains resource
// attributes like resource type and resource name. To get the whole view of the attributes
// used in IAM condition evaluation, the user must also look into
// AuditLog.request_metadata.request_attributes.
type Resource struct {
	Labels  map[string]interface{} `json:"labels,omitempty"` 
	Name    *string                `json:"name,omitempty"`   
	Service *string                `json:"service,omitempty"`
	Type    *string                `json:"type,omitempty"`   
}

type RequestMetadata struct {
	CallerIP                *string  `json:"caller_ip,omitempty"`                 
	CallerNetwork           *string  `json:"caller_network,omitempty"`            
	CallerSuppliedUserAgent *string  `json:"caller_supplied_user_agent,omitempty"`
	DestinationAttributes   *Peer    `json:"destination_attributes,omitempty"`    
	RequestAttributes       *Request `json:"request_attributes,omitempty"`        
}

type Peer struct {
	IP         *string                `json:"ip,omitempty"`         
	Labels     map[string]interface{} `json:"labels,omitempty"`     
	Port       *int64                 `json:"port,omitempty"`       
	Principal  *string                `json:"principal,omitempty"`  
	RegionCode *string                `json:"region_code,omitempty"`
}

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
	AccessLevels []string               `json:"access_levels,omitempty"`
	Audiences    []string               `json:"audiences,omitempty"`    
	Claims       map[string]interface{} `json:"claims,omitempty"`       
	Presenter    *string                `json:"presenter,omitempty"`    
	Principal    *string                `json:"principal,omitempty"`    
}

type ResourceLocation struct {
	CurrentLocations  []string `json:"current_locations,omitempty"` // The locations of a resource after the execution of the operation. Requests to create or; delete a location based resource must populate the 'current_locations' field and not the; 'original_locations' field.
	OriginalLocations []string `json:"original_locations,omitempty"`// The locations of a resource prior to the execution of the operation. Requests that mutate; the resource's location must populate both the 'original_locations' as well as the; 'current_locations' fields. For example:
}

type Status struct {
	Code    *int64      `json:"code,omitempty"`   
	Details interface{} `json:"details"`          
	Message *string     `json:"message,omitempty"`
}

type MonitoredResource struct {
	Labels map[string]interface{} `json:"labels,omitempty"`
	Type   *string                `json:"type,omitempty"`  
}

type Event struct {
	AnythingArray []interface{}
	Bool          *bool
	Double        *float64
	EventClass    *EventClass
	Integer       *int64
	String        *string
}

func (x *Event) UnmarshalJSON(data []byte) error {
	x.AnythingArray = nil
	x.EventClass = nil
	var c EventClass
	object, err := unmarshalUnion(data, &x.Integer, &x.Double, &x.Bool, &x.String, true, &x.AnythingArray, true, &c, false, nil, false, nil, true)
	if err != nil {
		return err
	}
	if object {
		x.EventClass = &c
	}
	return nil
}

func (x *Event) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, x.Double, x.Bool, x.String, x.AnythingArray != nil, x.AnythingArray, x.EventClass != nil, x.EventClass, false, nil, false, nil, true)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
