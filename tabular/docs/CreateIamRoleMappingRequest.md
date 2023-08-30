# CreateIamRoleMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**AwsRoleArn** | Pointer to **string** |  | [optional] 
**RoleId** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateIamRoleMappingRequest

`func NewCreateIamRoleMappingRequest() *CreateIamRoleMappingRequest`

NewCreateIamRoleMappingRequest instantiates a new CreateIamRoleMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIamRoleMappingRequestWithDefaults

`func NewCreateIamRoleMappingRequestWithDefaults() *CreateIamRoleMappingRequest`

NewCreateIamRoleMappingRequestWithDefaults instantiates a new CreateIamRoleMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateIamRoleMappingRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIamRoleMappingRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIamRoleMappingRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateIamRoleMappingRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAwsRoleArn

`func (o *CreateIamRoleMappingRequest) GetAwsRoleArn() string`

GetAwsRoleArn returns the AwsRoleArn field if non-nil, zero value otherwise.

### GetAwsRoleArnOk

`func (o *CreateIamRoleMappingRequest) GetAwsRoleArnOk() (*string, bool)`

GetAwsRoleArnOk returns a tuple with the AwsRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRoleArn

`func (o *CreateIamRoleMappingRequest) SetAwsRoleArn(v string)`

SetAwsRoleArn sets AwsRoleArn field to given value.

### HasAwsRoleArn

`func (o *CreateIamRoleMappingRequest) HasAwsRoleArn() bool`

HasAwsRoleArn returns a boolean if a field has been set.

### GetRoleId

`func (o *CreateIamRoleMappingRequest) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *CreateIamRoleMappingRequest) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *CreateIamRoleMappingRequest) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *CreateIamRoleMappingRequest) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


