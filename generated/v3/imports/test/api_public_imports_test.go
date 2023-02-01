/*
CRM Imports

Testing PublicImportsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package imports

import (
	openapiclient "./openapi"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_imports_PublicImportsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PublicImportsApiService GetErrors", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var importId int64

		resp, httpRes, err := apiClient.PublicImportsApi.GetErrors(context.Background(), importId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
