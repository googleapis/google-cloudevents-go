// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    events, err := UnmarshalEvents(bytes)
//    bytes, err = events.Marshal()

package auditv1

import "encoding/json"


func UnmarshalAuditLogWrittenEvent(data []byte) (AuditLogWrittenEvent, error) {
	var r AuditLogWrittenEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AuditLogWrittenEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Events struct {
	AuditLogWrittenEvent *AuditLogWrittenEvent `json:"AuditLogWrittenEvent,omitempty"`// The event is triggered when a new Cloud Audit Log entry is written.
}

// The event is triggered when a new Cloud Audit Log entry is written.
type AuditLogWrittenEvent struct {
	AuthenticationInfo    *AuthenticationInfo    `json:"authenticationInfo,omitempty"`   
	AuthorizationInfo     []AuthorizationInfo    `json:"authorizationInfo,omitempty"`    
	Metadata              map[string]interface{} `json:"metadata,omitempty"`             
	MethodName            *string                `json:"methodName,omitempty"`           
	NumResponseItems      *int64                 `json:"numResponseItems,omitempty"`     
	Request               map[string]interface{} `json:"request,omitempty"`              
	RequestMetadata       *RequestMetadata       `json:"requestMetadata,omitempty"`      
	ResourceLocation      *ResourceLocation      `json:"resourceLocation,omitempty"`     
	ResourceName          *string                `json:"resourceName,omitempty"`         
	ResourceOriginalState map[string]interface{} `json:"resourceOriginalState,omitempty"`
	Response              map[string]interface{} `json:"response,omitempty"`             
	ServiceData           map[string]interface{} `json:"serviceData,omitempty"`          
	ServiceName           *string                `json:"serviceName,omitempty"`          
	Status                *Status                `json:"status,omitempty"`               
}

type AuthenticationInfo struct {
	AuthoritySelector            *string                        `json:"authoritySelector,omitempty"`           
	PrincipalSubject             *string                        `json:"principalSubject,omitempty"`            
	PrincipleEmail               *string                        `json:"principleEmail,omitempty"`              
	ServiceAccountDelegationInfo []ServiceAccountDelegationInfo `json:"serviceAccountDelegationInfo,omitempty"`
	ServiceAccountKeyName        *string                        `json:"serviceAccountKeyName,omitempty"`       
	ThirdPartyPrincipal          map[string]interface{}         `json:"thirdPartyPrincipal,omitempty"`         
}

type ServiceAccountDelegationInfo struct {
	PrincipalEmail   *string                `json:"principalEmail,omitempty"`  
	ServiceMetadata  map[string]interface{} `json:"serviceMetadata,omitempty"` 
	ThirdPartyClaims map[string]interface{} `json:"thirdPartyClaims,omitempty"`
}

type AuthorizationInfo struct {
	Granted            *bool               `json:"granted,omitempty"`           
	Permission         *string             `json:"permission,omitempty"`        
	ResourceAttributes *ResourceAttributes `json:"resourceAttributes,omitempty"`
	Response           *string             `json:"response,omitempty"`          
}

type ResourceAttributes struct {
	Labels  map[string]interface{} `json:"labels,omitempty"` 
	Name    *string                `json:"name,omitempty"`   
	Service *string                `json:"service,omitempty"`
	Type    *string                `json:"type,omitempty"`   
}

type RequestMetadata struct {
	CallerIP                *string  `json:"callerIp,omitempty"`               
	CallerNetwork           *string  `json:"callerNetwork,omitempty"`          
	CallerSuppliedUserAgent *string  `json:"callerSuppliedUserAgent,omitempty"`
	DestinationAttributes   *Peer    `json:"destinationAttributes,omitempty"`  
	RequestAttributes       *Request `json:"requestAttributes,omitempty"`      
}

type Peer struct {
	IP         *string                `json:"ip,omitempty"`        
	Labels     map[string]interface{} `json:"labels,omitempty"`    
	Port       *int64                 `json:"port,omitempty"`      
	Principal  *string                `json:"principal,omitempty"` 
	RegionCode *string                `json:"regionCode,omitempty"`
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
	AccessLevels []string               `json:"accessLevels,omitempty"`
	Audiences    []string               `json:"audiences,omitempty"`   
	Claims       map[string]interface{} `json:"claims,omitempty"`      
	Presenter    *string                `json:"presenter,omitempty"`   
	Principal    *string                `json:"principal,omitempty"`   
}

type ResourceLocation struct {
	CurrentLocations  []string `json:"currentLocations,omitempty"` 
	OriginalLocations []string `json:"originalLocations,omitempty"`
}

type Status struct {
	Code    *int64        `json:"code,omitempty"`   
	Details []interface{} `json:"details,omitempty"`
	Message *string       `json:"message,omitempty"`
}
