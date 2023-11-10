/*
Core API

Testing DelinquencyAPIService

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

func Test_marqeta_coreapi_client_DelinquencyAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DelinquencyAPIService ResendWebhookEvent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var eventType string
		var resourceToken string

		resp, httpRes, err := apiClient.DelinquencyAPI.ResendWebhookEvent(context.Background(), eventType, resourceToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DelinquencyAPIService RetrieveDelinquencyState", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string

		resp, httpRes, err := apiClient.DelinquencyAPI.RetrieveDelinquencyState(context.Background(), accountToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DelinquencyAPIService RetrieveDelinquencyTransition", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string
		var delinquencyTransitionToken string

		resp, httpRes, err := apiClient.DelinquencyAPI.RetrieveDelinquencyTransition(context.Background(), accountToken, delinquencyTransitionToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DelinquencyAPIService RetrieveDelinquencyTransitions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string

		resp, httpRes, err := apiClient.DelinquencyAPI.RetrieveDelinquencyTransitions(context.Background(), accountToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
