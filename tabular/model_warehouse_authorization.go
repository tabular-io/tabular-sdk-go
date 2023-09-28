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

// checks if the WarehouseAuthorization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WarehouseAuthorization{}

// WarehouseAuthorization struct for WarehouseAuthorization
type WarehouseAuthorization struct {
	Id *string `json:"id,omitempty"`
	SubjectId *string `json:"subjectId,omitempty"`
	Privilege *string `json:"privilege,omitempty"`
	WithGrant *bool `json:"withGrant,omitempty"`
	Resource *string `json:"resource,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
}

// NewWarehouseAuthorization instantiates a new WarehouseAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarehouseAuthorization() *WarehouseAuthorization {
	this := WarehouseAuthorization{}
	return &this
}

// NewWarehouseAuthorizationWithDefaults instantiates a new WarehouseAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarehouseAuthorizationWithDefaults() *WarehouseAuthorization {
	this := WarehouseAuthorization{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WarehouseAuthorization) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseAuthorization) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WarehouseAuthorization) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WarehouseAuthorization) SetId(v string) {
	o.Id = &v
}

// GetSubjectId returns the SubjectId field value if set, zero value otherwise.
func (o *WarehouseAuthorization) GetSubjectId() string {
	if o == nil || IsNil(o.SubjectId) {
		var ret string
		return ret
	}
	return *o.SubjectId
}

// GetSubjectIdOk returns a tuple with the SubjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseAuthorization) GetSubjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectId) {
		return nil, false
	}
	return o.SubjectId, true
}

// HasSubjectId returns a boolean if a field has been set.
func (o *WarehouseAuthorization) HasSubjectId() bool {
	if o != nil && !IsNil(o.SubjectId) {
		return true
	}

	return false
}

// SetSubjectId gets a reference to the given string and assigns it to the SubjectId field.
func (o *WarehouseAuthorization) SetSubjectId(v string) {
	o.SubjectId = &v
}

// GetPrivilege returns the Privilege field value if set, zero value otherwise.
func (o *WarehouseAuthorization) GetPrivilege() string {
	if o == nil || IsNil(o.Privilege) {
		var ret string
		return ret
	}
	return *o.Privilege
}

// GetPrivilegeOk returns a tuple with the Privilege field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseAuthorization) GetPrivilegeOk() (*string, bool) {
	if o == nil || IsNil(o.Privilege) {
		return nil, false
	}
	return o.Privilege, true
}

// HasPrivilege returns a boolean if a field has been set.
func (o *WarehouseAuthorization) HasPrivilege() bool {
	if o != nil && !IsNil(o.Privilege) {
		return true
	}

	return false
}

// SetPrivilege gets a reference to the given string and assigns it to the Privilege field.
func (o *WarehouseAuthorization) SetPrivilege(v string) {
	o.Privilege = &v
}

// GetWithGrant returns the WithGrant field value if set, zero value otherwise.
func (o *WarehouseAuthorization) GetWithGrant() bool {
	if o == nil || IsNil(o.WithGrant) {
		var ret bool
		return ret
	}
	return *o.WithGrant
}

// GetWithGrantOk returns a tuple with the WithGrant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseAuthorization) GetWithGrantOk() (*bool, bool) {
	if o == nil || IsNil(o.WithGrant) {
		return nil, false
	}
	return o.WithGrant, true
}

// HasWithGrant returns a boolean if a field has been set.
func (o *WarehouseAuthorization) HasWithGrant() bool {
	if o != nil && !IsNil(o.WithGrant) {
		return true
	}

	return false
}

// SetWithGrant gets a reference to the given bool and assigns it to the WithGrant field.
func (o *WarehouseAuthorization) SetWithGrant(v bool) {
	o.WithGrant = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *WarehouseAuthorization) GetResource() string {
	if o == nil || IsNil(o.Resource) {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseAuthorization) GetResourceOk() (*string, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *WarehouseAuthorization) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *WarehouseAuthorization) SetResource(v string) {
	o.Resource = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *WarehouseAuthorization) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseAuthorization) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *WarehouseAuthorization) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *WarehouseAuthorization) SetResourceType(v string) {
	o.ResourceType = &v
}

func (o WarehouseAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WarehouseAuthorization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SubjectId) {
		toSerialize["subjectId"] = o.SubjectId
	}
	if !IsNil(o.Privilege) {
		toSerialize["privilege"] = o.Privilege
	}
	if !IsNil(o.WithGrant) {
		toSerialize["withGrant"] = o.WithGrant
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !IsNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	return toSerialize, nil
}

type NullableWarehouseAuthorization struct {
	value *WarehouseAuthorization
	isSet bool
}

func (v NullableWarehouseAuthorization) Get() *WarehouseAuthorization {
	return v.value
}

func (v *NullableWarehouseAuthorization) Set(val *WarehouseAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullableWarehouseAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableWarehouseAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarehouseAuthorization(val *WarehouseAuthorization) *NullableWarehouseAuthorization {
	return &NullableWarehouseAuthorization{value: val, isSet: true}
}

func (v NullableWarehouseAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarehouseAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

