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

package firestore

// DocumentEventData: The data within all Firestore document events.
type DocumentEventData struct {
	OldValue   *OldValue   `json:"oldValue,omitempty"`   // A Document object containing a pre-operation document snapshot.; This is only populated for update and delete events.
	UpdateMask *UpdateMask `json:"updateMask,omitempty"` // A DocumentMask object that lists changed fields.; This is only populated for update events.
	Value      *Value      `json:"value,omitempty"`      // A Document object containing a post-operation document snapshot.; This is not populated for delete events. (TODO: check this!)
}

// OldValue: A Document object containing a pre-operation document snapshot.
// This is only populated for update and delete events.
//
// A Firestore document.
type OldValue struct {
	CreateTime *string                  `json:"createTime,omitempty"` // The time at which the document was created.; ; This value increases monotonically when a document is deleted then; recreated. It can also be compared to values from other documents and; the `read_time` of a query.
	Fields     map[string]OldValueField `json:"fields,omitempty"`     // The document's fields.; ; The map keys represent field names.; ; A simple field name contains only characters `a` to `z`, `A` to `Z`,; `0` to `9`, or `_`, and must not start with `0` to `9`. For example,; `foo_bar_17`.; ; Field names matching the regular expression `__.*__` are reserved. Reserved; field names are forbidden except in certain documented contexts. The map; keys, represented as UTF-8, must not exceed 1,500 bytes and cannot be; empty.; ; Field paths may be used in other contexts to refer to structured fields; defined here. For `map_value`, the field path is represented by the simple; or quoted field names of the containing fields, delimited by `.`. For; example, the structured field; `"foo" : { map_value: { "x&y" : { string_value: "hello" }}}` would be; represented by the field path `foo.x&y`.; ; Within a field path, a quoted field name starts and ends with `` ` `` and; may contain any character. Some characters, including `` ` ``, must be; escaped using a `\`. For example, `` `x&y` `` represents `x&y` and; `` `bak\`tik` `` represents `` bak`tik ``.
	Name       *string                  `json:"name,omitempty"`       // The resource name of the document. For example:; `projects/{project_id}/databases/{database_id}/documents/{document_path}`
	UpdateTime *string                  `json:"updateTime,omitempty"` // The time at which the document was last changed.; ; This value is initially set to the `create_time` then increases; monotonically with each change to the document. It can also be; compared to values from other documents and the `read_time` of a query.
}

// OldValueField: A message that can hold any of the supported value types.
type OldValueField struct {
	ArrayValue     *ArrayValue     `json:"arrayValue,omitempty"`          // An array value.; ; Cannot directly contain another array value, though can contain an; map which contains another array.
	BooleanValue   *bool           `json:"booleanValue,omitempty"`        // A boolean value.
	BytesValue     *string         `json:"bytesValue,omitempty"`          // A bytes value.; ; Must not exceed 1 MiB - 89 bytes.; Only the first 1,500 bytes are considered by queries.
	DoubleValue    *float64        `json:"doubleValue,omitempty"`         // A double value.
	GeoPointValue  *GeoPointValue  `json:"geoPointValue,omitempty"`       // A geo point value representing a point on the surface of Earth.
	IntegerValue   *int64          `json:"integerValue,string,omitempty"` // An integer value.
	MapValue       *MapValue       `json:"mapValue,omitempty"`            // A map value.
	NullValue      *NullValueUnion `json:"nullValue"`                     // A null value.
	ReferenceValue *string         `json:"referenceValue,omitempty"`      // A reference to a document. For example:; `projects/{project_id}/databases/{database_id}/documents/{document_path}`.
	StringValue    *string         `json:"stringValue,omitempty"`         // A string value.; ; The string, represented as UTF-8, must not exceed 1 MiB - 89 bytes.; Only the first 1,500 bytes of the UTF-8 representation are considered by; queries.
	TimestampValue *string         `json:"timestampValue,omitempty"`      // A timestamp value.; ; Precise only to microseconds. When stored, any additional precision is; rounded down.
}

// MapValueField: A message that can hold any of the supported value types.
type MapValueField struct {
	ArrayValue     *ArrayValue     `json:"arrayValue,omitempty"`          // An array value.; ; Cannot directly contain another array value, though can contain an; map which contains another array.
	BooleanValue   *bool           `json:"booleanValue,omitempty"`        // A boolean value.
	BytesValue     *string         `json:"bytesValue,omitempty"`          // A bytes value.; ; Must not exceed 1 MiB - 89 bytes.; Only the first 1,500 bytes are considered by queries.
	DoubleValue    *float64        `json:"doubleValue,omitempty"`         // A double value.
	GeoPointValue  *GeoPointValue  `json:"geoPointValue,omitempty"`       // A geo point value representing a point on the surface of Earth.
	IntegerValue   *int64          `json:"integerValue,string,omitempty"` // An integer value.
	MapValue       *MapValue       `json:"mapValue,omitempty"`            // A map value.
	NullValue      *NullValueUnion `json:"nullValue"`                     // A null value.
	ReferenceValue *string         `json:"referenceValue,omitempty"`      // A reference to a document. For example:; `projects/{project_id}/databases/{database_id}/documents/{document_path}`.
	StringValue    *string         `json:"stringValue,omitempty"`         // A string value.; ; The string, represented as UTF-8, must not exceed 1 MiB - 89 bytes.; Only the first 1,500 bytes of the UTF-8 representation are considered by; queries.
	TimestampValue *string         `json:"timestampValue,omitempty"`      // A timestamp value.; ; Precise only to microseconds. When stored, any additional precision is; rounded down.
}

// MapValue: A map value.
type MapValue struct {
	Fields map[string]MapValueField `json:"fields,omitempty"` // The map's fields.; ; The map keys represent field names. Field names matching the regular; expression `__.*__` are reserved. Reserved field names are forbidden except; in certain documented contexts. The map keys, represented as UTF-8, must; not exceed 1,500 bytes and cannot be empty.
}

// ValueElement: A message that can hold any of the supported value types.
type ValueElement struct {
	ArrayValue     *ArrayValue     `json:"arrayValue,omitempty"`          // An array value.; ; Cannot directly contain another array value, though can contain an; map which contains another array.
	BooleanValue   *bool           `json:"booleanValue,omitempty"`        // A boolean value.
	BytesValue     *string         `json:"bytesValue,omitempty"`          // A bytes value.; ; Must not exceed 1 MiB - 89 bytes.; Only the first 1,500 bytes are considered by queries.
	DoubleValue    *float64        `json:"doubleValue,omitempty"`         // A double value.
	GeoPointValue  *GeoPointValue  `json:"geoPointValue,omitempty"`       // A geo point value representing a point on the surface of Earth.
	IntegerValue   *int64          `json:"integerValue,string,omitempty"` // An integer value.
	MapValue       *MapValue       `json:"mapValue,omitempty"`            // A map value.
	NullValue      *NullValueUnion `json:"nullValue"`                     // A null value.
	ReferenceValue *string         `json:"referenceValue,omitempty"`      // A reference to a document. For example:; `projects/{project_id}/databases/{database_id}/documents/{document_path}`.
	StringValue    *string         `json:"stringValue,omitempty"`         // A string value.; ; The string, represented as UTF-8, must not exceed 1 MiB - 89 bytes.; Only the first 1,500 bytes of the UTF-8 representation are considered by; queries.
	TimestampValue *string         `json:"timestampValue,omitempty"`      // A timestamp value.; ; Precise only to microseconds. When stored, any additional precision is; rounded down.
}

// ArrayValue: An array value.
//
// Cannot directly contain another array value, though can contain an
// map which contains another array.
type ArrayValue struct {
	Values []ValueElement `json:"values,omitempty"` // Values in the array.
}

// GeoPointValue: A geo point value representing a point on the surface of Earth.
type GeoPointValue struct {
	Latitude  *float64 `json:"latitude,omitempty"`  // The latitude in degrees. It must be in the range [-90.0, +90.0].
	Longitude *float64 `json:"longitude,omitempty"` // The longitude in degrees. It must be in the range [-180.0, +180.0].
}

// UpdateMask: A DocumentMask object that lists changed fields.
// This is only populated for update events.
type UpdateMask struct {
	FieldPaths []string `json:"fieldPaths,omitempty"` // The list of field paths in the mask.; See [Document.fields][google.cloud.firestore.v1.events.Document.fields]; for a field path syntax reference.
}

// Value: A Document object containing a post-operation document snapshot.
// This is not populated for delete events. (TODO: check this!)
//
// A Document object containing a pre-operation document snapshot.
// This is only populated for update and delete events.
//
// A Firestore document.
type Value struct {
	CreateTime *string                  `json:"createTime,omitempty"` // The time at which the document was created.; ; This value increases monotonically when a document is deleted then; recreated. It can also be compared to values from other documents and; the `read_time` of a query.
	Fields     map[string]OldValueField `json:"fields,omitempty"`     // The document's fields.; ; The map keys represent field names.; ; A simple field name contains only characters `a` to `z`, `A` to `Z`,; `0` to `9`, or `_`, and must not start with `0` to `9`. For example,; `foo_bar_17`.; ; Field names matching the regular expression `__.*__` are reserved. Reserved; field names are forbidden except in certain documented contexts. The map; keys, represented as UTF-8, must not exceed 1,500 bytes and cannot be; empty.; ; Field paths may be used in other contexts to refer to structured fields; defined here. For `map_value`, the field path is represented by the simple; or quoted field names of the containing fields, delimited by `.`. For; example, the structured field; `"foo" : { map_value: { "x&y" : { string_value: "hello" }}}` would be; represented by the field path `foo.x&y`.; ; Within a field path, a quoted field name starts and ends with `` ` `` and; may contain any character. Some characters, including `` ` ``, must be; escaped using a `\`. For example, `` `x&y` `` represents `x&y` and; `` `bak\`tik` `` represents `` bak`tik ``.
	Name       *string                  `json:"name,omitempty"`       // The resource name of the document. For example:; `projects/{project_id}/databases/{database_id}/documents/{document_path}`
	UpdateTime *string                  `json:"updateTime,omitempty"` // The time at which the document was last changed.; ; This value is initially set to the `create_time` then increases; monotonically with each change to the document. It can also be; compared to values from other documents and the `read_time` of a query.
}

type CreateTime string

const (
	NullValue CreateTime = "NULL_VALUE"
)

// NullValueUnion: A null value.
type NullValueUnion struct {
	Enum    *CreateTime
	Integer *int64
}
