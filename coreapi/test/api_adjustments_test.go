/*
Core API

Testing AdjustmentsAPIService

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

func Test_marqeta_coreapi_client_AdjustmentsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AdjustmentsAPIService CreateAdjustmentForAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string

		resp, httpRes, err := apiClient.AdjustmentsAPI.CreateAdjustmentForAccount(context.Background(), accountToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdjustmentsAPIService GetAdjustmentsByAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string

		resp, httpRes, err := apiClient.AdjustmentsAPI.GetAdjustmentsByAccount(context.Background(), accountToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdjustmentsAPIService RetrieveAdjustment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string
		var adjustmentToken string

		resp, httpRes, err := apiClient.AdjustmentsAPI.RetrieveAdjustment(context.Background(), accountToken, adjustmentToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
