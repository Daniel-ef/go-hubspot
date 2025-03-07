# Go API client for schemas

The CRM uses schemas to define how custom objects should store and represent information in the HubSpot CRM. Schemas define details about an object's type, properties, and associations. The schema can be uniquely identified by its **object type ID**.

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
import schemas "github.com/GIT_USER_ID/GIT_REPO_ID"
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
ctx := context.WithValue(context.Background(), schemas.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), schemas.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), schemas.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), schemas.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.hubapi.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CoreApi* | [**Archive**](docs/CoreApi.md#archive) | **Delete** /crm/v3/schemas/{objectType} | Delete a schema
*CoreApi* | [**ArchiveAssociation**](docs/CoreApi.md#archiveassociation) | **Delete** /crm/v3/schemas/{objectType}/associations/{associationIdentifier} | Remove an association
*CoreApi* | [**Create**](docs/CoreApi.md#create) | **Post** /crm/v3/schemas | Create a new schema
*CoreApi* | [**CreateAssociation**](docs/CoreApi.md#createassociation) | **Post** /crm/v3/schemas/{objectType}/associations | Create an association
*CoreApi* | [**GetAll**](docs/CoreApi.md#getall) | **Get** /crm/v3/schemas | Get all schemas
*CoreApi* | [**GetByID**](docs/CoreApi.md#getbyid) | **Get** /crm/v3/schemas/{objectType} | Get an existing schema
*CoreApi* | [**Update**](docs/CoreApi.md#update) | **Patch** /crm/v3/schemas/{objectType} | Update a schema
*PublicObjectSchemasApi* | [**Purge**](docs/PublicObjectSchemasApi.md#purge) | **Delete** /crm/v3/schemas/{objectType}/purge | 


## Documentation For Models

 - [AssociationDefinition](docs/AssociationDefinition.md)
 - [AssociationDefinitionEgg](docs/AssociationDefinitionEgg.md)
 - [CollectionResponseObjectSchemaNoPaging](docs/CollectionResponseObjectSchemaNoPaging.md)
 - [Error](docs/Error.md)
 - [ErrorDetail](docs/ErrorDetail.md)
 - [ObjectSchema](docs/ObjectSchema.md)
 - [ObjectSchemaEgg](docs/ObjectSchemaEgg.md)
 - [ObjectTypeDefinition](docs/ObjectTypeDefinition.md)
 - [ObjectTypeDefinitionLabels](docs/ObjectTypeDefinitionLabels.md)
 - [ObjectTypeDefinitionPatch](docs/ObjectTypeDefinitionPatch.md)
 - [ObjectTypePropertyCreate](docs/ObjectTypePropertyCreate.md)
 - [Option](docs/Option.md)
 - [OptionInput](docs/OptionInput.md)
 - [Property](docs/Property.md)
 - [PropertyModificationMetadata](docs/PropertyModificationMetadata.md)


## Documentation For Authorization



### hapikey

- **Type**: API key
- **API key parameter name**: hapikey
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: hapikey and passed in as the auth context for each request.


### oauth2_legacy


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://app.hubspot.com/oauth/authorize
- **Scopes**: 
 - **crm.schemas.custom.read**: View custom object definitions
 - **crm.objects.custom.read**: View custom object records

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



