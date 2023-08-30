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

// checks if the MemberEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberEntry{}

// MemberEntry struct for MemberEntry
type MemberEntry struct {
	Id *string `json:"id,omitempty"`
	Email *string `json:"email,omitempty"`
	WithAdmin *bool `json:"withAdmin,omitempty"`
}

// NewMemberEntry instantiates a new MemberEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberEntry() *MemberEntry {
	this := MemberEntry{}
	return &this
}

// NewMemberEntryWithDefaults instantiates a new MemberEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberEntryWithDefaults() *MemberEntry {
	this := MemberEntry{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MemberEntry) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberEntry) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MemberEntry) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MemberEntry) SetId(v string) {
	o.Id = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *MemberEntry) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberEntry) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *MemberEntry) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *MemberEntry) SetEmail(v string) {
	o.Email = &v
}

// GetWithAdmin returns the WithAdmin field value if set, zero value otherwise.
func (o *MemberEntry) GetWithAdmin() bool {
	if o == nil || IsNil(o.WithAdmin) {
		var ret bool
		return ret
	}
	return *o.WithAdmin
}

// GetWithAdminOk returns a tuple with the WithAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberEntry) GetWithAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.WithAdmin) {
		return nil, false
	}
	return o.WithAdmin, true
}

// HasWithAdmin returns a boolean if a field has been set.
func (o *MemberEntry) HasWithAdmin() bool {
	if o != nil && !IsNil(o.WithAdmin) {
		return true
	}

	return false
}

// SetWithAdmin gets a reference to the given bool and assigns it to the WithAdmin field.
func (o *MemberEntry) SetWithAdmin(v bool) {
	o.WithAdmin = &v
}

func (o MemberEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.WithAdmin) {
		toSerialize["withAdmin"] = o.WithAdmin
	}
	return toSerialize, nil
}

type NullableMemberEntry struct {
	value *MemberEntry
	isSet bool
}

func (v NullableMemberEntry) Get() *MemberEntry {
	return v.value
}

func (v *NullableMemberEntry) Set(val *MemberEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberEntry(val *MemberEntry) *NullableMemberEntry {
	return &NullableMemberEntry{value: val, isSet: true}
}

func (v NullableMemberEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


