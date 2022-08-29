// Code generated by protoc-gen-go-snowpea. DO NOT EDIT.

package v1

import (
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
)

// UnmarshalLogEntryData provides a typed object.
func UnmarshalLogEntryData(datatype string, data []byte) (*LogEntryData, error) {
	out := LogEntryData{}

	switch datatype {
	case "application/json":
		pj := protojson.UnmarshalOptions{DiscardUnknown: true}
		if err := pj.Unmarshal(data, &out); err != nil {
			return nil, fmt.Errorf("protojson.Unmarshal: %w", err)
		}
	default:
		return nil, fmt.Errorf("data type %q not supported", datatype)
	}

	return &out, nil
}

// UnmarshalStrictLogEntryData provides a typed object but errors on unknown fields.
func UnmarshalStrictLogEntryData(data []byte) (*LogEntryData, error) {
	out := LogEntryData{}
	if err := protojson.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("protojson.Unmarshal: %w", err)
	}

	return &out, nil
}

// JSON provides the protobuf as a serialized JSON object.
func (m *LogEntryData) JSON() ([]byte, error) {
	b, err := protojson.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("protojson.Marshal: %w", err)
	}

	return b, nil
}

// CloudEventTypePrefix returns the prefix of the CloudEvent Type this data type supports.
func CloudEventTypePrefix() string {
	return "google.events.cloud.audit.v1"
}
