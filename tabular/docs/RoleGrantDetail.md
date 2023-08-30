# RoleGrantDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Role** | Pointer to [**RoleRef**](RoleRef.md) |  | [optional] 
**Privilege** | Pointer to **string** |  | [optional] 
**WithGrant** | Pointer to **bool** |  | [optional] 

## Methods

### NewRoleGrantDetail

`func NewRoleGrantDetail() *RoleGrantDetail`

NewRoleGrantDetail instantiates a new RoleGrantDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleGrantDetailWithDefaults

`func NewRoleGrantDetailWithDefaults() *RoleGrantDetail`

NewRoleGrantDetailWithDefaults instantiates a new RoleGrantDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoleGrantDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleGrantDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleGrantDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoleGrantDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRole

`func (o *RoleGrantDetail) GetRole() RoleRef`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RoleGrantDetail) GetRoleOk() (*RoleRef, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RoleGrantDetail) SetRole(v RoleRef)`

SetRole sets Role field to given value.

### HasRole

`func (o *RoleGrantDetail) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetPrivilege

`func (o *RoleGrantDetail) GetPrivilege() string`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *RoleGrantDetail) GetPrivilegeOk() (*string, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *RoleGrantDetail) SetPrivilege(v string)`

SetPrivilege sets Privilege field to given value.

### HasPrivilege

`func (o *RoleGrantDetail) HasPrivilege() bool`

HasPrivilege returns a boolean if a field has been set.

### GetWithGrant

`func (o *RoleGrantDetail) GetWithGrant() bool`

GetWithGrant returns the WithGrant field if non-nil, zero value otherwise.

### GetWithGrantOk

`func (o *RoleGrantDetail) GetWithGrantOk() (*bool, bool)`

GetWithGrantOk returns a tuple with the WithGrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithGrant

`func (o *RoleGrantDetail) SetWithGrant(v bool)`

SetWithGrant sets WithGrant field to given value.

### HasWithGrant

`func (o *RoleGrantDetail) HasWithGrant() bool`

HasWithGrant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


