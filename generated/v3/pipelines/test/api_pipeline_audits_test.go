/*
CRM Pipelines

Testing PipelineAuditsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package pipelines

import (
	openapiclient "./openapi"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_pipelines_PipelineAuditsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PipelineAuditsApiService GetAudit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectType string
		var pipelineId string

		resp, httpRes, err := apiClient.PipelineAuditsApi.GetAudit(context.Background(), objectType, pipelineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
