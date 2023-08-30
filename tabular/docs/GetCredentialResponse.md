# GetCredentialResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**MemberId** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**RoleId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Cidrs** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Creator** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetCredentialResponse

`func NewGetCredentialResponse() *GetCredentialResponse`

NewGetCredentialResponse instantiates a new GetCredentialResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCredentialResponseWithDefaults

`func NewGetCredentialResponseWithDefaults() *GetCredentialResponse`

NewGetCredentialResponseWithDefaults instantiates a new GetCredentialResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCredentialResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCredentialResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCredentialResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetCredentialResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *GetCredentialResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetCredentialResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetCredentialResponse) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GetCredentialResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetUserId

`func (o *GetCredentialResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetCredentialResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetCredentialResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GetCredentialResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetMemberId

`func (o *GetCredentialResponse) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *GetCredentialResponse) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *GetCredentialResponse) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *GetCredentialResponse) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *GetCredentialResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetCredentialResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetCredentialResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *GetCredentialResponse) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetRoleId

`func (o *GetCredentialResponse) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *GetCredentialResponse) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *GetCredentialResponse) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *GetCredentialResponse) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetName

`func (o *GetCredentialResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCredentialResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCredentialResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCredentialResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetCredentialResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetCredentialResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetCredentialResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetCredentialResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCidrs

`func (o *GetCredentialResponse) GetCidrs() []string`

GetCidrs returns the Cidrs field if non-nil, zero value otherwise.

### GetCidrsOk

`func (o *GetCredentialResponse) GetCidrsOk() (*[]string, bool)`

GetCidrsOk returns a tuple with the Cidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrs

`func (o *GetCredentialResponse) SetCidrs(v []string)`

SetCidrs sets Cidrs field to given value.

### HasCidrs

`func (o *GetCredentialResponse) HasCidrs() bool`

HasCidrs returns a boolean if a field has been set.

### GetType

`func (o *GetCredentialResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCredentialResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCredentialResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetCredentialResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreator

`func (o *GetCredentialResponse) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *GetCredentialResponse) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *GetCredentialResponse) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *GetCredentialResponse) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetActive

`func (o *GetCredentialResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetCredentialResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetCredentialResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetCredentialResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


