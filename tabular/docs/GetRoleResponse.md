# GetRoleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Members** | Pointer to [**[]MemberEntry**](MemberEntry.md) |  | [optional] 
**Children** | Pointer to [**[]RoleRef**](RoleRef.md) |  | [optional] 

## Methods

### NewGetRoleResponse

`func NewGetRoleResponse() *GetRoleResponse`

NewGetRoleResponse instantiates a new GetRoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRoleResponseWithDefaults

`func NewGetRoleResponseWithDefaults() *GetRoleResponse`

NewGetRoleResponseWithDefaults instantiates a new GetRoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetRoleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRoleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRoleResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetRoleResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetRoleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRoleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRoleResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetRoleResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMembers

`func (o *GetRoleResponse) GetMembers() []MemberEntry`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GetRoleResponse) GetMembersOk() (*[]MemberEntry, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GetRoleResponse) SetMembers(v []MemberEntry)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GetRoleResponse) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetChildren

`func (o *GetRoleResponse) GetChildren() []RoleRef`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *GetRoleResponse) GetChildrenOk() (*[]RoleRef, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *GetRoleResponse) SetChildren(v []RoleRef)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *GetRoleResponse) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


