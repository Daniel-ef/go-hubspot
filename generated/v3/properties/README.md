# Go API client for properties

All HubSpot objects store data in default and custom properties. These endpoints provide access to read and modify object properties in HubSpot.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v3
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import properties "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), properties.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), properties.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), properties.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), properties.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.hubapi.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BatchApi* | [**BatchArchive**](docs/BatchApi.md#batcharchive) | **Post** /crm/v3/properties/{objectType}/batch/archive | Archive a batch of properties
*BatchApi* | [**BatchCreate**](docs/BatchApi.md#batchcreate) | **Post** /crm/v3/properties/{objectType}/batch/create | Create a batch of properties
*BatchApi* | [**BatchRead**](docs/BatchApi.md#batchread) | **Post** /crm/v3/properties/{objectType}/batch/read | Read a batch of properties
*CoreApi* | [**Archive**](docs/CoreApi.md#archive) | **Delete** /crm/v3/properties/{objectType}/{propertyName} | Archive a property
*CoreApi* | [**Create**](docs/CoreApi.md#create) | **Post** /crm/v3/properties/{objectType} | Create a property
*CoreApi* | [**GetAll**](docs/CoreApi.md#getall) | **Get** /crm/v3/properties/{objectType} | Read all properties
*CoreApi* | [**GetByName**](docs/CoreApi.md#getbyname) | **Get** /crm/v3/properties/{objectType}/{propertyName} | Read a property
*CoreApi* | [**Update**](docs/CoreApi.md#update) | **Patch** /crm/v3/properties/{objectType}/{propertyName} | Update a property
*GroupsApi* | [**GroupsArchive**](docs/GroupsApi.md#groupsarchive) | **Delete** /crm/v3/properties/{objectType}/groups/{groupName} | Archive a property group
*GroupsApi* | [**GroupsCreate**](docs/GroupsApi.md#groupscreate) | **Post** /crm/v3/properties/{objectType}/groups | Create a property group
*GroupsApi* | [**GroupsGetAll**](docs/GroupsApi.md#groupsgetall) | **Get** /crm/v3/properties/{objectType}/groups | Read all property groups
*GroupsApi* | [**GroupsGetByName**](docs/GroupsApi.md#groupsgetbyname) | **Get** /crm/v3/properties/{objectType}/groups/{groupName} | Read a property group
*GroupsApi* | [**GroupsUpdate**](docs/GroupsApi.md#groupsupdate) | **Patch** /crm/v3/properties/{objectType}/groups/{groupName} | Update a property group


## Documentation For Models

 - [BatchInputPropertyCreate](docs/BatchInputPropertyCreate.md)
 - [BatchInputPropertyName](docs/BatchInputPropertyName.md)
 - [BatchReadInputPropertyName](docs/BatchReadInputPropertyName.md)
 - [BatchResponseProperty](docs/BatchResponseProperty.md)
 - [CollectionResponseProperty](docs/CollectionResponseProperty.md)
 - [CollectionResponsePropertyGroup](docs/CollectionResponsePropertyGroup.md)
 - [Error](docs/Error.md)
 - [ErrorCategory](docs/ErrorCategory.md)
 - [ErrorDetail](docs/ErrorDetail.md)
 - [NextPage](docs/NextPage.md)
 - [Option](docs/Option.md)
 - [OptionInput](docs/OptionInput.md)
 - [Paging](docs/Paging.md)
 - [Property](docs/Property.md)
 - [PropertyCreate](docs/PropertyCreate.md)
 - [PropertyGroup](docs/PropertyGroup.md)
 - [PropertyGroupCreate](docs/PropertyGroupCreate.md)
 - [PropertyGroupUpdate](docs/PropertyGroupUpdate.md)
 - [PropertyModificationMetadata](docs/PropertyModificationMetadata.md)
 - [PropertyName](docs/PropertyName.md)
 - [PropertyUpdate](docs/PropertyUpdate.md)
 - [StandardError](docs/StandardError.md)


## Documentation For Authorization



### hapikey

- **Type**: API key
- **API key parameter name**: hapikey
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: hapikey and passed in as the auth context for each request.


### oauth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://app.hubspot.com/oauth/authorize
- **Scopes**: 
 - **crm.objects.contacts.read**:  
 - **crm.schemas.quotes.read**: Quotes schemas
 - **crm.objects.contacts.write**:  
 - **crm.schemas.contacts.read**:  
 - **crm.objects.companies.write**:  
 - **crm.objects.companies.read**:  
 - **crm.schemas.companies.read**:  
 - **crm.objects.deals.read**:  
 - **crm.schemas.line_items.read**: Line Items schemas
 - **crm.objects.deals.write**:  
 - **crm.schemas.deals.read**:  
 - **crm.schemas.deals.write**:  
 - **crm.schemas.contacts.write**:  
 - **crm.schemas.companies.write**:  

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


### oauth2_legacy


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://app.hubspot.com/oauth/authorize
- **Scopes**: 
 - **timeline**: Create timeline events
 - **tickets**: Read and write tickets
 - **media_bridge.read**: Read media and media events
 - **crm.schemas.custom.read**: View custom object definitions
 - **e-commerce**: e-commerce

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


### private_apps

- **Type**: API key
- **API key parameter name**: private-app
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: private-app and passed in as the auth context for each request.


### private_apps_legacy

- **Type**: API key
- **API key parameter name**: private-app-legacy
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: private-app-legacy and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



