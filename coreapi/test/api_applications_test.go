/*
Core API

Testing ApplicationsAPIService

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

func Test_marqeta_coreapi_client_ApplicationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ApplicationsAPIService CreditApplicationsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ApplicationsAPI.CreditApplicationsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService GetFileByType", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ FileType

		resp, httpRes, err := apiClient.ApplicationsAPI.GetFileByType(context.Background(), type_).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService PageApplicationTransitions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string

		resp, httpRes, err := apiClient.ApplicationsAPI.PageApplicationTransitions(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService RetrieveApplication", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string

		resp, httpRes, err := apiClient.ApplicationsAPI.RetrieveApplication(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService RetrieveFiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ApplicationsAPI.RetrieveFiles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService TransitionApplication", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string

		resp, httpRes, err := apiClient.ApplicationsAPI.TransitionApplication(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}