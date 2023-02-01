/*
Accounting Extension

Testing UserAccountsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package accounting

import (
	openapiclient "./openapi"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_accounting_UserAccountsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserAccountsApiService UserAccountsArchive", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.UserAccountsApi.UserAccountsArchive(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAccountsApiService UserAccountsReplace", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserAccountsApi.UserAccountsReplace(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
