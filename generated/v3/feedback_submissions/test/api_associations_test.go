/*
Feedback Submissions

Testing AssociationsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package feedback_submissions

import (
	openapiclient "./openapi"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_feedback_submissions_AssociationsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AssociationsApiService AssociationsGetPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var feedbackSubmissionId int64
		var toObjectType string

		resp, httpRes, err := apiClient.AssociationsApi.AssociationsGetPage(context.Background(), feedbackSubmissionId, toObjectType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssociationsApiService DeleteCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdArchive", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var feedbackSubmissionId int64
		var toObjectType string
		var toObjectId int64

		resp, httpRes, err := apiClient.AssociationsApi.DeleteCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdArchive(context.Background(), feedbackSubmissionId, toObjectType, toObjectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssociationsApiService PutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var feedbackSubmissionId int64
		var toObjectType string
		var toObjectId int64

		resp, httpRes, err := apiClient.AssociationsApi.PutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreate(context.Background(), feedbackSubmissionId, toObjectType, toObjectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
