/*
Contacts

Testing AssociationsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package contacts

import (
	openapiclient "./openapi"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_contacts_AssociationsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AssociationsApiService AssociationsGetAll", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var contactId int64
		var toObjectType string

		resp, httpRes, err := apiClient.AssociationsApi.AssociationsGetAll(context.Background(), contactId, toObjectType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssociationsApiService DeleteCrmV4ObjectsContactsContactIdAssociationsToObjectTypeToObjectIdArchive", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var contactId int64
		var toObjectType string
		var toObjectId int64

		resp, httpRes, err := apiClient.AssociationsApi.DeleteCrmV4ObjectsContactsContactIdAssociationsToObjectTypeToObjectIdArchive(context.Background(), contactId, toObjectType, toObjectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssociationsApiService PutCrmV4ObjectsContactsContactIdAssociationsToObjectTypeToObjectIdCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var contactId int64
		var toObjectType string
		var toObjectId int64

		resp, httpRes, err := apiClient.AssociationsApi.PutCrmV4ObjectsContactsContactIdAssociationsToObjectTypeToObjectIdCreate(context.Background(), contactId, toObjectType, toObjectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
