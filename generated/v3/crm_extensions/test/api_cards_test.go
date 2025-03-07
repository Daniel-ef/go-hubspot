/*
CRM cards

Testing CardsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package crm_extensions

import (
	openapiclient "./openapi"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_crm_extensions_CardsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CardsApiService CardsArchive", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32
		var cardId string

		resp, httpRes, err := apiClient.CardsApi.CardsArchive(context.Background(), appId, cardId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CardsApiService CardsCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32

		resp, httpRes, err := apiClient.CardsApi.CardsCreate(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CardsApiService CardsGetAll", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32

		resp, httpRes, err := apiClient.CardsApi.CardsGetAll(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CardsApiService CardsGetByID", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32
		var cardId string

		resp, httpRes, err := apiClient.CardsApi.CardsGetByID(context.Background(), appId, cardId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CardsApiService CardsUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32
		var cardId string

		resp, httpRes, err := apiClient.CardsApi.CardsUpdate(context.Background(), appId, cardId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
