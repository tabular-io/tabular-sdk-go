/*
OpenAPI definition

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tabular

import (
	"encoding/json"
)

// checks if the CreateRoleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRoleResponse{}

// CreateRoleResponse struct for CreateRoleResponse
type CreateRoleResponse struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewCreateRoleResponse instantiates a new CreateRoleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRoleResponse() *CreateRoleResponse {
	this := CreateRoleResponse{}
	return &this
}

// NewCreateRoleResponseWithDefaults instantiates a new CreateRoleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRoleResponseWithDefaults() *CreateRoleResponse {
	this := CreateRoleResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateRoleResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateRoleResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateRoleResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateRoleResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateRoleResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateRoleResponse) SetName(v string) {
	o.Name = &v
}

func (o CreateRoleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRoleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCreateRoleResponse struct {
	value *CreateRoleResponse
	isSet bool
}

func (v NullableCreateRoleResponse) Get() *CreateRoleResponse {
	return v.value
}

func (v *NullableCreateRoleResponse) Set(val *CreateRoleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRoleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRoleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRoleResponse(val *CreateRoleResponse) *NullableCreateRoleResponse {
	return &NullableCreateRoleResponse{value: val, isSet: true}
}

func (v NullableCreateRoleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRoleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


