# MemberEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**WithAdmin** | Pointer to **bool** |  | [optional] 

## Methods

### NewMemberEntry

`func NewMemberEntry() *MemberEntry`

NewMemberEntry instantiates a new MemberEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberEntryWithDefaults

`func NewMemberEntryWithDefaults() *MemberEntry`

NewMemberEntryWithDefaults instantiates a new MemberEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MemberEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MemberEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *MemberEntry) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MemberEntry) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MemberEntry) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MemberEntry) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetWithAdmin

`func (o *MemberEntry) GetWithAdmin() bool`

GetWithAdmin returns the WithAdmin field if non-nil, zero value otherwise.

### GetWithAdminOk

`func (o *MemberEntry) GetWithAdminOk() (*bool, bool)`

GetWithAdminOk returns a tuple with the WithAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithAdmin

`func (o *MemberEntry) SetWithAdmin(v bool)`

SetWithAdmin sets WithAdmin field to given value.

### HasWithAdmin

`func (o *MemberEntry) HasWithAdmin() bool`

HasWithAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


