/*
Core API

Testing CreditDisputesAPIService

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

func Test_marqeta_coreapi_client_CreditDisputesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CreditDisputesAPIService CreateDisputeForAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string

		resp, httpRes, err := apiClient.CreditDisputesAPI.CreateDisputeForAccount(context.Background(), accountToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CreditDisputesAPIService GetDisputesByAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string

		resp, httpRes, err := apiClient.CreditDisputesAPI.GetDisputesByAccount(context.Background(), accountToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CreditDisputesAPIService RetrieveDispute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string
		var disputeToken string

		resp, httpRes, err := apiClient.CreditDisputesAPI.RetrieveDispute(context.Background(), accountToken, disputeToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CreditDisputesAPIService TransitionDispute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string
		var disputeToken string

		resp, httpRes, err := apiClient.CreditDisputesAPI.TransitionDispute(context.Background(), accountToken, disputeToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}