# ListMembersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to [**[]MemberRef**](MemberRef.md) |  | [optional] 

## Methods

### NewListMembersResponse

`func NewListMembersResponse() *ListMembersResponse`

NewListMembersResponse instantiates a new ListMembersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMembersResponseWithDefaults

`func NewListMembersResponseWithDefaults() *ListMembersResponse`

NewListMembersResponseWithDefaults instantiates a new ListMembersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *ListMembersResponse) GetMembers() []MemberRef`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ListMembersResponse) GetMembersOk() (*[]MemberRef, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ListMembersResponse) SetMembers(v []MemberRef)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ListMembersResponse) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


