/*
CRM Objects

Testing AssociationsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package objects

import (
	openapiclient "./openapi"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_objects_AssociationsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AssociationsApiService AssociationsArchive", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectType string
		var objectId string
		var toObjectType string
		var toObjectId string
		var associationType string

		resp, httpRes, err := apiClient.AssociationsApi.AssociationsArchive(context.Background(), objectType, objectId, toObjectType, toObjectId, associationType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssociationsApiService AssociationsCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectType string
		var objectId string
		var toObjectType string
		var toObjectId string
		var associationType string

		resp, httpRes, err := apiClient.AssociationsApi.AssociationsCreate(context.Background(), objectType, objectId, toObjectType, toObjectId, associationType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssociationsApiService AssociationsGetAll", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectType string
		var objectId string
		var toObjectType string

		resp, httpRes, err := apiClient.AssociationsApi.AssociationsGetAll(context.Background(), objectType, objectId, toObjectType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
