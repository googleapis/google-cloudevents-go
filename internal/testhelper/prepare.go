package testhelper

import (
	"encoding/json"
	"fmt"
)

// PreparePubSubMessagePublishedData modifies data for protojson parsing.
// - Remove fields protojson considers duplicate field names
// - Remove @type properties
func PreparePubSubMessagePublishedData(b []byte) ([]byte, error) {
	var j map[string]interface{}
	if err := json.Unmarshal(b, &j); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}
	m := j["message"].(map[string]interface{})

	// Remove reserved @type field.
	delete(m, "@type")
	// Remove message_id field name conflict with messageId.
	delete(m, "message_id")
	// Remove publish_time field name conflict with publishTime.
	delete(m, "publish_time")

	j["message"] = m
	return json.Marshal(j)
}

// PrepareAuditLogEntryData modifies data for protojson parsing.
// - Remove @type properties
func PrepareAuditLogEntryData(b []byte) ([]byte, error) {
	var j map[string]interface{}
	if err := json.Unmarshal(b, &j); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}
	m := j["protoPayload"].(map[string]interface{})

	// Remove reserved @type field.
	delete(m, "@type")

	j["protoPayload"] = m
	return json.Marshal(j)
}
