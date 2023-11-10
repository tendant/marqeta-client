/*
Core API

Testing StatementsAPIService

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

func Test_marqeta_coreapi_client_StatementsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StatementsAPIService GetStatementFilesByAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string

		resp, httpRes, err := apiClient.StatementsAPI.GetStatementFilesByAccount(context.Background(), accountToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StatementsAPIService GetStatementSummariesByAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string

		resp, httpRes, err := apiClient.StatementsAPI.GetStatementSummariesByAccount(context.Background(), accountToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StatementsAPIService ListStatementJournalEntries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string
		var statementSummaryToken string

		resp, httpRes, err := apiClient.StatementsAPI.ListStatementJournalEntries(context.Background(), accountToken, statementSummaryToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StatementsAPIService ListStatementLedgerEntries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string
		var statementSummaryToken string

		resp, httpRes, err := apiClient.StatementsAPI.ListStatementLedgerEntries(context.Background(), accountToken, statementSummaryToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StatementsAPIService ResendWebhookEvent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var eventType string
		var resourceToken string

		resp, httpRes, err := apiClient.StatementsAPI.ResendWebhookEvent(context.Background(), eventType, resourceToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StatementsAPIService RetrieveStatementFiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string
		var statementSummaryToken string

		resp, httpRes, err := apiClient.StatementsAPI.RetrieveStatementFiles(context.Background(), accountToken, statementSummaryToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StatementsAPIService RetrieveStatementInterestCharges", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string
		var statementSummaryToken string

		resp, httpRes, err := apiClient.StatementsAPI.RetrieveStatementInterestCharges(context.Background(), accountToken, statementSummaryToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StatementsAPIService RetrieveStatementPaymentInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string
		var statementSummaryToken string

		resp, httpRes, err := apiClient.StatementsAPI.RetrieveStatementPaymentInfo(context.Background(), accountToken, statementSummaryToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StatementsAPIService RetrieveStatementReward", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string
		var statementSummaryToken string

		resp, httpRes, err := apiClient.StatementsAPI.RetrieveStatementReward(context.Background(), accountToken, statementSummaryToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StatementsAPIService RetrieveStatementSummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string
		var statementSummaryToken string

		resp, httpRes, err := apiClient.StatementsAPI.RetrieveStatementSummary(context.Background(), accountToken, statementSummaryToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StatementsAPIService RetrieveYearToDateForStatementSummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string
		var statementSummaryToken string

		resp, httpRes, err := apiClient.StatementsAPI.RetrieveYearToDateForStatementSummary(context.Background(), accountToken, statementSummaryToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}