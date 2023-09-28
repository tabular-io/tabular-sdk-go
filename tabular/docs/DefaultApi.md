# \DefaultApi

All URIs are relative to *https://api.tabular.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddChildToRole**](DefaultApi.md#AddChildToRole) | **Put** /v1/organizations/{organizationId}/roles/{roleName}/children | Add child to role
[**AddRoleMembers**](DefaultApi.md#AddRoleMembers) | **Put** /v1/organizations/{organizationId}/roles/{roleName}/members | Add members to a role
[**CreateCustomerIdentityProviderCredential**](DefaultApi.md#CreateCustomerIdentityProviderCredential) | **Post** /v1/organizations/{organizationId}/iam/credentials/custom-idp | Create a custom identity provider credential
[**CreateDatabase**](DefaultApi.md#CreateDatabase) | **Post** /v1/organizations/{organizationId}/warehouses/{warehouseId}/databases/ | Creates database with extended information
[**CreateIamRoleMapping**](DefaultApi.md#CreateIamRoleMapping) | **Post** /v1/organizations/{organizationId}/iam/credentials/aws | Create an IAM Role Mapping
[**CreateMemberCredential**](DefaultApi.md#CreateMemberCredential) | **Post** /v1/organizations/{organizationId}/iam/credentials/member | Create a member credential
[**CreateOIDCIntegration**](DefaultApi.md#CreateOIDCIntegration) | **Post** /v1/organizations/{organizationId}/iam/credentials/oidc | Create an OIDC Integration
[**CreateRole**](DefaultApi.md#CreateRole) | **Post** /v1/organizations/{organizationId}/roles/ | Create role
[**CreateServiceAccountCredential**](DefaultApi.md#CreateServiceAccountCredential) | **Post** /v1/organizations/{organizationId}/iam/credentials/service-account | Create a service credential
[**CreateStorageProfile**](DefaultApi.md#CreateStorageProfile) | **Post** /v1/organizations/{organizationId}/storage-profiles/ | Create a storage profile
[**CreateWarehouse**](DefaultApi.md#CreateWarehouse) | **Post** /v1/organizations/{organizationId}/warehouses/ | Create a warehouse
[**DeleteDatabase**](DefaultApi.md#DeleteDatabase) | **Delete** /v1/organizations/{organizationId}/warehouses/{warehouseId}/databases/{database} | Drops a database
[**DeleteMemberCredential**](DefaultApi.md#DeleteMemberCredential) | **Delete** /v1/organizations/{organizationId}/iam/credentials/{credentialKey} | Delete a credential
[**DeleteRole**](DefaultApi.md#DeleteRole) | **Delete** /v1/organizations/{organizationId}/roles/{roleName} | Delete Role
[**DeleteStorageProfile**](DefaultApi.md#DeleteStorageProfile) | **Delete** /v1/organizations/{organizationId}/storage-profiles/{storageProfileId} | Delete a storage profile by ID
[**DeleteWarehouse**](DefaultApi.md#DeleteWarehouse) | **Delete** /v1/organizations/{organizationId}/warehouses/{warehouseId} | Delete warehouse by id
[**GetCredential**](DefaultApi.md#GetCredential) | **Get** /v1/organizations/{organizationId}/iam/credentials/{credentialKey} | Fetch information about a credential
[**GetDatabase**](DefaultApi.md#GetDatabase) | **Get** /v1/organizations/{organizationId}/warehouses/{warehouseId}/databases/{database} | Gets extended information on a database
[**GetDatabaseRoleGrants**](DefaultApi.md#GetDatabaseRoleGrants) | **Get** /v1/organizations/{organizationId}/warehouses/{warehouseId}/databases/{database}/grants | List all grants for database
[**GetOrganizationMembers**](DefaultApi.md#GetOrganizationMembers) | **Get** /v1/organizations/{organizationId}/members/ | Get organization members
[**GetRole**](DefaultApi.md#GetRole) | **Get** /v1/organizations/{organizationId}/roles/{roleName} | Get role
[**GetRoleWarehouseGrants**](DefaultApi.md#GetRoleWarehouseGrants) | **Get** /v1/organizations/{organizationId}/warehouses/{warehouseId}/grants/roles/{roleId} | Get warehouse grants by role name
[**GetStorageProfile**](DefaultApi.md#GetStorageProfile) | **Get** /v1/organizations/{organizationId}/storage-profiles/{storageProfileId} | Get a storage profile by ID
[**GetWarehouse**](DefaultApi.md#GetWarehouse) | **Get** /v1/organizations/{organizationId}/warehouses/{warehouseId} | Get a warehouse by id
[**GrantPrivilegesOnDatabase**](DefaultApi.md#GrantPrivilegesOnDatabase) | **Put** /v1/organizations/{organizationId}/warehouses/{warehouseId}/databases/{database}/grants | Grant privileges on database
[**GrantPrivilegesOnWarehouse**](DefaultApi.md#GrantPrivilegesOnWarehouse) | **Put** /v1/organizations/{organizationId}/warehouses/{warehouseId}/grants | Grant privileges on a warehouse
[**ListWarehouses**](DefaultApi.md#ListWarehouses) | **Get** /v1/organizations/{organizationId}/warehouses/ | List all warehouses
[**RemoveChildFromRole**](DefaultApi.md#RemoveChildFromRole) | **Delete** /v1/organizations/{organizationId}/roles/{roleName}/children | Remove child from role
[**RemoveRoleMembers**](DefaultApi.md#RemoveRoleMembers) | **Delete** /v1/organizations/{organizationId}/roles/{roleName}/members | Remove members from a role
[**RevokePrivilegesOnDatabase**](DefaultApi.md#RevokePrivilegesOnDatabase) | **Delete** /v1/organizations/{organizationId}/warehouses/{warehouseId}/databases/{database}/grants | Revoke privileges on database
[**RevokePrivilegesOnWarehouse**](DefaultApi.md#RevokePrivilegesOnWarehouse) | **Delete** /v1/organizations/{organizationId}/warehouses/{warehouseId}/grants | Revoke privileges on a warehouse
[**UpdateRoleName**](DefaultApi.md#UpdateRoleName) | **Put** /v1/organizations/{organizationId}/roles/{roleName} | Update role



## AddChildToRole

> AddChildToRole(ctx, organizationId, roleName).UpdateRoleRequest(updateRoleRequest).Execute()

Add child to role

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    roleName := "roleName_example" // string | 
    updateRoleRequest := *openapiclient.NewUpdateRoleRequest() // UpdateRoleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.AddChildToRole(context.Background(), organizationId, roleName).UpdateRoleRequest(updateRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddChildToRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddChildToRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRoleRequest** | [**UpdateRoleRequest**](UpdateRoleRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddRoleMembers

> AddRoleMembers(ctx, organizationId, roleName).UpdateRoleMemberRequest(updateRoleMemberRequest).Execute()

Add members to a role

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    roleName := "roleName_example" // string | 
    updateRoleMemberRequest := []openapiclient.UpdateRoleMemberRequest{*openapiclient.NewUpdateRoleMemberRequest()} // []UpdateRoleMemberRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.AddRoleMembers(context.Background(), organizationId, roleName).UpdateRoleMemberRequest(updateRoleMemberRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddRoleMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRoleMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRoleMemberRequest** | [**[]UpdateRoleMemberRequest**](UpdateRoleMemberRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomerIdentityProviderCredential

> CreateCredentialResponse CreateCustomerIdentityProviderCredential(ctx, organizationId).CreateCustomIdpCredentialRequest(createCustomIdpCredentialRequest).Execute()

Create a custom identity provider credential

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    createCustomIdpCredentialRequest := *openapiclient.NewCreateCustomIdpCredentialRequest() // CreateCustomIdpCredentialRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateCustomerIdentityProviderCredential(context.Background(), organizationId).CreateCustomIdpCredentialRequest(createCustomIdpCredentialRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCustomerIdentityProviderCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomerIdentityProviderCredential`: CreateCredentialResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCustomerIdentityProviderCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerIdentityProviderCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCustomIdpCredentialRequest** | [**CreateCustomIdpCredentialRequest**](CreateCustomIdpCredentialRequest.md) |  | 

### Return type

[**CreateCredentialResponse**](CreateCredentialResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDatabase

> CreateDatabaseResponse CreateDatabase(ctx, organizationId, warehouseId).CreateDatabaseRequest(createDatabaseRequest).Execute()

Creates database with extended information

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    createDatabaseRequest := *openapiclient.NewCreateDatabaseRequest() // CreateDatabaseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateDatabase(context.Background(), organizationId, warehouseId).CreateDatabaseRequest(createDatabaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDatabase`: CreateDatabaseResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createDatabaseRequest** | [**CreateDatabaseRequest**](CreateDatabaseRequest.md) |  | 

### Return type

[**CreateDatabaseResponse**](CreateDatabaseResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIamRoleMapping

> CreateCredentialResponse CreateIamRoleMapping(ctx, organizationId).CreateIamRoleMappingRequest(createIamRoleMappingRequest).Execute()

Create an IAM Role Mapping

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    createIamRoleMappingRequest := *openapiclient.NewCreateIamRoleMappingRequest() // CreateIamRoleMappingRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateIamRoleMapping(context.Background(), organizationId).CreateIamRoleMappingRequest(createIamRoleMappingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIamRoleMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIamRoleMapping`: CreateCredentialResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIamRoleMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIamRoleMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createIamRoleMappingRequest** | [**CreateIamRoleMappingRequest**](CreateIamRoleMappingRequest.md) |  | 

### Return type

[**CreateCredentialResponse**](CreateCredentialResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMemberCredential

> CreateCredentialResponse CreateMemberCredential(ctx, organizationId).CreateMemberCredentialRequest(createMemberCredentialRequest).Execute()

Create a member credential

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    createMemberCredentialRequest := *openapiclient.NewCreateMemberCredentialRequest() // CreateMemberCredentialRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateMemberCredential(context.Background(), organizationId).CreateMemberCredentialRequest(createMemberCredentialRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateMemberCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMemberCredential`: CreateCredentialResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateMemberCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMemberCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createMemberCredentialRequest** | [**CreateMemberCredentialRequest**](CreateMemberCredentialRequest.md) |  | 

### Return type

[**CreateCredentialResponse**](CreateCredentialResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOIDCIntegration

> CreateCredentialResponse CreateOIDCIntegration(ctx, organizationId).CreateOidcCredentialRequest(createOidcCredentialRequest).Execute()

Create an OIDC Integration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    createOidcCredentialRequest := *openapiclient.NewCreateOidcCredentialRequest() // CreateOidcCredentialRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateOIDCIntegration(context.Background(), organizationId).CreateOidcCredentialRequest(createOidcCredentialRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateOIDCIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOIDCIntegration`: CreateCredentialResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateOIDCIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOIDCIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOidcCredentialRequest** | [**CreateOidcCredentialRequest**](CreateOidcCredentialRequest.md) |  | 

### Return type

[**CreateCredentialResponse**](CreateCredentialResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> CreateRoleResponse CreateRole(ctx, organizationId).CreateRoleRequest(createRoleRequest).Execute()

Create role

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    createRoleRequest := *openapiclient.NewCreateRoleRequest() // CreateRoleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateRole(context.Background(), organizationId).CreateRoleRequest(createRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRole`: CreateRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRoleRequest** | [**CreateRoleRequest**](CreateRoleRequest.md) |  | 

### Return type

[**CreateRoleResponse**](CreateRoleResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceAccountCredential

> CreateCredentialResponse CreateServiceAccountCredential(ctx, organizationId).CreateServiceAccountCredentialRequest(createServiceAccountCredentialRequest).Execute()

Create a service credential

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    createServiceAccountCredentialRequest := *openapiclient.NewCreateServiceAccountCredentialRequest() // CreateServiceAccountCredentialRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateServiceAccountCredential(context.Background(), organizationId).CreateServiceAccountCredentialRequest(createServiceAccountCredentialRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateServiceAccountCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceAccountCredential`: CreateCredentialResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateServiceAccountCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAccountCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createServiceAccountCredentialRequest** | [**CreateServiceAccountCredentialRequest**](CreateServiceAccountCredentialRequest.md) |  | 

### Return type

[**CreateCredentialResponse**](CreateCredentialResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStorageProfile

> CreateS3StorageProfileResponse CreateStorageProfile(ctx, organizationId).CreateS3StorageProfileRequest(createS3StorageProfileRequest).CreatorRoleId(creatorRoleId).Execute()

Create a storage profile

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    createS3StorageProfileRequest := *openapiclient.NewCreateS3StorageProfileRequest() // CreateS3StorageProfileRequest | 
    creatorRoleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateStorageProfile(context.Background(), organizationId).CreateS3StorageProfileRequest(createS3StorageProfileRequest).CreatorRoleId(creatorRoleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateStorageProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorageProfile`: CreateS3StorageProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateStorageProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorageProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createS3StorageProfileRequest** | [**CreateS3StorageProfileRequest**](CreateS3StorageProfileRequest.md) |  | 
 **creatorRoleId** | **string** |  | 

### Return type

[**CreateS3StorageProfileResponse**](CreateS3StorageProfileResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWarehouse

> CreateWarehouseResponse CreateWarehouse(ctx, organizationId).CreateWarehouseRequest(createWarehouseRequest).CreatorRoleId(creatorRoleId).Execute()

Create a warehouse

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "organizationId_example" // string | 
    createWarehouseRequest := *openapiclient.NewCreateWarehouseRequest() // CreateWarehouseRequest | 
    creatorRoleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateWarehouse(context.Background(), organizationId).CreateWarehouseRequest(createWarehouseRequest).CreatorRoleId(creatorRoleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateWarehouse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWarehouse`: CreateWarehouseResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateWarehouse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createWarehouseRequest** | [**CreateWarehouseRequest**](CreateWarehouseRequest.md) |  | 
 **creatorRoleId** | **string** |  | 

### Return type

[**CreateWarehouseResponse**](CreateWarehouseResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatabase

> DeleteDatabase(ctx, organizationId, warehouseId, database).Execute()

Drops a database

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    database := "database_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteDatabase(context.Background(), organizationId, warehouseId, database).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 
**database** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMemberCredential

> DeleteMemberCredential(ctx, organizationId, credentialKey).Execute()

Delete a credential

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    credentialKey := "credentialKey_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteMemberCredential(context.Background(), organizationId, credentialKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteMemberCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**credentialKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMemberCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> DeleteRole(ctx, organizationId, roleName).Force(force).Execute()

Delete Role

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    roleName := "roleName_example" // string | 
    force := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteRole(context.Background(), organizationId, roleName).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **bool** |  | [default to true]

### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStorageProfile

> DeleteStorageProfile(ctx, organizationId, storageProfileId).Execute()

Delete a storage profile by ID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    storageProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteStorageProfile(context.Background(), organizationId, storageProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteStorageProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**storageProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStorageProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWarehouse

> DeleteWarehouse(ctx, organizationId, warehouseId).Execute()

Delete warehouse by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteWarehouse(context.Background(), organizationId, warehouseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteWarehouse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCredential

> GetCredentialResponse GetCredential(ctx, organizationId, credentialKey).Execute()

Fetch information about a credential

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    credentialKey := "credentialKey_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetCredential(context.Background(), organizationId, credentialKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCredential`: GetCredentialResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**credentialKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetCredentialResponse**](GetCredentialResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabase

> GetDatabaseResponse GetDatabase(ctx, organizationId, warehouseId, database).Execute()

Gets extended information on a database

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    database := "database_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetDatabase(context.Background(), organizationId, warehouseId, database).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabase`: GetDatabaseResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 
**database** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetDatabaseResponse**](GetDatabaseResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseRoleGrants

> ListDatabaseRoleGrantsResponse GetDatabaseRoleGrants(ctx, organizationId, warehouseId, database).Execute()

List all grants for database

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    database := "database_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetDatabaseRoleGrants(context.Background(), organizationId, warehouseId, database).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDatabaseRoleGrants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseRoleGrants`: ListDatabaseRoleGrantsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDatabaseRoleGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 
**database** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseRoleGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ListDatabaseRoleGrantsResponse**](ListDatabaseRoleGrantsResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationMembers

> ListMembersResponse GetOrganizationMembers(ctx, organizationId).Execute()

Get organization members

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetOrganizationMembers(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetOrganizationMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationMembers`: ListMembersResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetOrganizationMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListMembersResponse**](ListMembersResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> GetRoleResponse GetRole(ctx, organizationId, roleName).Execute()

Get role

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    roleName := "roleName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRole(context.Background(), organizationId, roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRole`: GetRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetRoleResponse**](GetRoleResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleWarehouseGrants

> GetRoleWarehouseGrantsResponse GetRoleWarehouseGrants(ctx, organizationId, warehouseId, roleId).Execute()

Get warehouse grants by role name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    roleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRoleWarehouseGrants(context.Background(), organizationId, warehouseId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRoleWarehouseGrants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleWarehouseGrants`: GetRoleWarehouseGrantsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRoleWarehouseGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleWarehouseGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetRoleWarehouseGrantsResponse**](GetRoleWarehouseGrantsResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageProfile

> GetStorageProfileResponse GetStorageProfile(ctx, organizationId, storageProfileId).Execute()

Get a storage profile by ID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    storageProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetStorageProfile(context.Background(), organizationId, storageProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStorageProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorageProfile`: GetStorageProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStorageProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**storageProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetStorageProfileResponse**](GetStorageProfileResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWarehouse

> GetWarehouseResponse GetWarehouse(ctx, organizationId, warehouseId).Execute()

Get a warehouse by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetWarehouse(context.Background(), organizationId, warehouseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetWarehouse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWarehouse`: GetWarehouseResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetWarehouse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetWarehouseResponse**](GetWarehouseResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GrantPrivilegesOnDatabase

> GrantPrivilegesOnDatabase(ctx, organizationId, warehouseId, database).ChangeRoleGrantRequest(changeRoleGrantRequest).Execute()

Grant privileges on database

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    database := "database_example" // string | 
    changeRoleGrantRequest := []openapiclient.ChangeRoleGrantRequest{*openapiclient.NewChangeRoleGrantRequest()} // []ChangeRoleGrantRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.GrantPrivilegesOnDatabase(context.Background(), organizationId, warehouseId, database).ChangeRoleGrantRequest(changeRoleGrantRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GrantPrivilegesOnDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 
**database** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGrantPrivilegesOnDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **changeRoleGrantRequest** | [**[]ChangeRoleGrantRequest**](ChangeRoleGrantRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GrantPrivilegesOnWarehouse

> GrantPrivilegesOnWarehouse(ctx, organizationId, warehouseId).RoleWarehouseGrantRequest(roleWarehouseGrantRequest).Execute()

Grant privileges on a warehouse

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    roleWarehouseGrantRequest := []openapiclient.RoleWarehouseGrantRequest{*openapiclient.NewRoleWarehouseGrantRequest()} // []RoleWarehouseGrantRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.GrantPrivilegesOnWarehouse(context.Background(), organizationId, warehouseId).RoleWarehouseGrantRequest(roleWarehouseGrantRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GrantPrivilegesOnWarehouse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGrantPrivilegesOnWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleWarehouseGrantRequest** | [**[]RoleWarehouseGrantRequest**](RoleWarehouseGrantRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWarehouses

> ListWarehouseResponse ListWarehouses(ctx, organizationId).Execute()

List all warehouses

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListWarehouses(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListWarehouses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWarehouses`: ListWarehouseResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListWarehouses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWarehousesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListWarehouseResponse**](ListWarehouseResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveChildFromRole

> RemoveChildFromRole(ctx, organizationId, roleName).UpdateRoleRequest(updateRoleRequest).Execute()

Remove child from role

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    roleName := "roleName_example" // string | 
    updateRoleRequest := *openapiclient.NewUpdateRoleRequest() // UpdateRoleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.RemoveChildFromRole(context.Background(), organizationId, roleName).UpdateRoleRequest(updateRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RemoveChildFromRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveChildFromRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRoleRequest** | [**UpdateRoleRequest**](UpdateRoleRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveRoleMembers

> RemoveRoleMembers(ctx, organizationId, roleName).UpdateRoleMemberRequest(updateRoleMemberRequest).Execute()

Remove members from a role

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    roleName := "roleName_example" // string | 
    updateRoleMemberRequest := []openapiclient.UpdateRoleMemberRequest{*openapiclient.NewUpdateRoleMemberRequest()} // []UpdateRoleMemberRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.RemoveRoleMembers(context.Background(), organizationId, roleName).UpdateRoleMemberRequest(updateRoleMemberRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RemoveRoleMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRoleMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRoleMemberRequest** | [**[]UpdateRoleMemberRequest**](UpdateRoleMemberRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokePrivilegesOnDatabase

> RevokePrivilegesOnDatabase(ctx, organizationId, warehouseId, database).ChangeRoleGrantRequest(changeRoleGrantRequest).Execute()

Revoke privileges on database

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    database := "database_example" // string | 
    changeRoleGrantRequest := []openapiclient.ChangeRoleGrantRequest{*openapiclient.NewChangeRoleGrantRequest()} // []ChangeRoleGrantRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.RevokePrivilegesOnDatabase(context.Background(), organizationId, warehouseId, database).ChangeRoleGrantRequest(changeRoleGrantRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RevokePrivilegesOnDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 
**database** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokePrivilegesOnDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **changeRoleGrantRequest** | [**[]ChangeRoleGrantRequest**](ChangeRoleGrantRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokePrivilegesOnWarehouse

> RevokePrivilegesOnWarehouse(ctx, organizationId, warehouseId).RoleWarehouseGrantRequest(roleWarehouseGrantRequest).Execute()

Revoke privileges on a warehouse

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    roleWarehouseGrantRequest := []openapiclient.RoleWarehouseGrantRequest{*openapiclient.NewRoleWarehouseGrantRequest()} // []RoleWarehouseGrantRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.RevokePrivilegesOnWarehouse(context.Background(), organizationId, warehouseId).RoleWarehouseGrantRequest(roleWarehouseGrantRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RevokePrivilegesOnWarehouse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokePrivilegesOnWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleWarehouseGrantRequest** | [**[]RoleWarehouseGrantRequest**](RoleWarehouseGrantRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoleName

> UpdateRoleNameResponse UpdateRoleName(ctx, organizationId, roleName).UpdateRoleRequest(updateRoleRequest).Execute()

Update role

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    roleName := "roleName_example" // string | 
    updateRoleRequest := *openapiclient.NewUpdateRoleRequest() // UpdateRoleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateRoleName(context.Background(), organizationId, roleName).UpdateRoleRequest(updateRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoleName`: UpdateRoleNameResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateRoleName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRoleRequest** | [**UpdateRoleRequest**](UpdateRoleRequest.md) |  | 

### Return type

[**UpdateRoleNameResponse**](UpdateRoleNameResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

