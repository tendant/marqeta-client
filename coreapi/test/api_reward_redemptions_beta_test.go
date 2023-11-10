/*
Core API

Testing RewardRedemptionsBetaAPIService

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

func Test_marqeta_coreapi_client_RewardRedemptionsBetaAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RewardRedemptionsBetaAPIService GetRedemption", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string
		var redemptionToken string

		resp, httpRes, err := apiClient.RewardRedemptionsBetaAPI.GetRedemption(context.Background(), token, redemptionToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RewardRedemptionsBetaAPIService PostRedemptionTransition", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string
		var redemptionToken string

		resp, httpRes, err := apiClient.RewardRedemptionsBetaAPI.PostRedemptionTransition(context.Background(), token, redemptionToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RewardRedemptionsBetaAPIService PostRewardProgramRedemption", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string

		resp, httpRes, err := apiClient.RewardRedemptionsBetaAPI.PostRewardProgramRedemption(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RewardRedemptionsBetaAPIService RetrieveRedemptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string

		resp, httpRes, err := apiClient.RewardRedemptionsBetaAPI.RetrieveRedemptions(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RewardRedemptionsBetaAPIService RetrieveRedemptionsBalance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string

		resp, httpRes, err := apiClient.RewardRedemptionsBetaAPI.RetrieveRedemptionsBalance(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
