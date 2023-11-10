/*
Core API

Testing RewardProgramsBetaAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package marqeta_coreapi_client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/tendant/marqeta-client"
)

func Test_marqeta_coreapi_client_RewardProgramsBetaAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RewardProgramsBetaAPIService PostRewardProgramEntry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string

		resp, httpRes, err := apiClient.RewardProgramsBetaAPI.PostRewardProgramEntry(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RewardProgramsBetaAPIService RetrieveRewardProgram", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string

		resp, httpRes, err := apiClient.RewardProgramsBetaAPI.RetrieveRewardProgram(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RewardProgramsBetaAPIService RetrieveRewardProgramBalance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string

		resp, httpRes, err := apiClient.RewardProgramsBetaAPI.RetrieveRewardProgramBalance(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RewardProgramsBetaAPIService RetrieveRewardProgramEntries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string

		resp, httpRes, err := apiClient.RewardProgramsBetaAPI.RetrieveRewardProgramEntries(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RewardProgramsBetaAPIService RetrieveRewardProgramEntriesBalance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string

		resp, httpRes, err := apiClient.RewardProgramsBetaAPI.RetrieveRewardProgramEntriesBalance(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RewardProgramsBetaAPIService RetrieveRewardProgramEntry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string
		var entryToken string

		resp, httpRes, err := apiClient.RewardProgramsBetaAPI.RetrieveRewardProgramEntry(context.Background(), token, entryToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RewardProgramsBetaAPIService RetrieveRewardPrograms", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RewardProgramsBetaAPI.RetrieveRewardPrograms(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RewardProgramsBetaAPIService RetrieveRewardProgramsRulesConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string

		resp, httpRes, err := apiClient.RewardProgramsBetaAPI.RetrieveRewardProgramsRulesConfig(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RewardProgramsBetaAPIService RetrieveRewardProgramsRulesConfigApplied", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string

		resp, httpRes, err := apiClient.RewardProgramsBetaAPI.RetrieveRewardProgramsRulesConfigApplied(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RewardProgramsBetaAPIService UpdateRewardProgram", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string

		resp, httpRes, err := apiClient.RewardProgramsBetaAPI.UpdateRewardProgram(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
