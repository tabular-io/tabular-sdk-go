# UpdateRoleMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberId** | Pointer to **string** |  | [optional] 
**WithAdmin** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateRoleMemberRequest

`func NewUpdateRoleMemberRequest() *UpdateRoleMemberRequest`

NewUpdateRoleMemberRequest instantiates a new UpdateRoleMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleMemberRequestWithDefaults

`func NewUpdateRoleMemberRequestWithDefaults() *UpdateRoleMemberRequest`

NewUpdateRoleMemberRequestWithDefaults instantiates a new UpdateRoleMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberId

`func (o *UpdateRoleMemberRequest) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *UpdateRoleMemberRequest) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *UpdateRoleMemberRequest) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *UpdateRoleMemberRequest) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetWithAdmin

`func (o *UpdateRoleMemberRequest) GetWithAdmin() bool`

GetWithAdmin returns the WithAdmin field if non-nil, zero value otherwise.

### GetWithAdminOk

`func (o *UpdateRoleMemberRequest) GetWithAdminOk() (*bool, bool)`

GetWithAdminOk returns a tuple with the WithAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithAdmin

`func (o *UpdateRoleMemberRequest) SetWithAdmin(v bool)`

SetWithAdmin sets WithAdmin field to given value.

### HasWithAdmin

`func (o *UpdateRoleMemberRequest) HasWithAdmin() bool`

HasWithAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


