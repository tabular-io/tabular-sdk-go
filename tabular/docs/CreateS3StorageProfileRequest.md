# CreateS3StorageProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** |  | [optional] 
**Bucket** | Pointer to **string** |  | [optional] 
**RoleArn** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateS3StorageProfileRequest

`func NewCreateS3StorageProfileRequest() *CreateS3StorageProfileRequest`

NewCreateS3StorageProfileRequest instantiates a new CreateS3StorageProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateS3StorageProfileRequestWithDefaults

`func NewCreateS3StorageProfileRequestWithDefaults() *CreateS3StorageProfileRequest`

NewCreateS3StorageProfileRequestWithDefaults instantiates a new CreateS3StorageProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *CreateS3StorageProfileRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateS3StorageProfileRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateS3StorageProfileRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreateS3StorageProfileRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetBucket

`func (o *CreateS3StorageProfileRequest) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *CreateS3StorageProfileRequest) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *CreateS3StorageProfileRequest) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *CreateS3StorageProfileRequest) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetRoleArn

`func (o *CreateS3StorageProfileRequest) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *CreateS3StorageProfileRequest) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *CreateS3StorageProfileRequest) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *CreateS3StorageProfileRequest) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


