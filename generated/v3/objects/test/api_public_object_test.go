/*
CRM Objects

Testing PublicObjectApiService

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

func Test_objects_PublicObjectApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PublicObjectApiService Merge", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectType string

		resp, httpRes, err := apiClient.PublicObjectApi.Merge(context.Background(), objectType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
