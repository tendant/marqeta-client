/*
Core API

Testing GPAOrdersAPIService

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

func Test_marqeta_coreapi_client_GPAOrdersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GPAOrdersAPIService GetGpaordersToken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string

		resp, httpRes, err := apiClient.GPAOrdersAPI.GetGpaordersToken(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GPAOrdersAPIService GetGpaordersUnloads", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GPAOrdersAPI.GetGpaordersUnloads(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GPAOrdersAPIService GetGpaordersUnloadsUnloadtoken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var unloadToken string

		resp, httpRes, err := apiClient.GPAOrdersAPI.GetGpaordersUnloadsUnloadtoken(context.Background(), unloadToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GPAOrdersAPIService PostGpaorders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GPAOrdersAPI.PostGpaorders(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GPAOrdersAPIService PostGpaordersUnloads", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GPAOrdersAPI.PostGpaordersUnloads(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
