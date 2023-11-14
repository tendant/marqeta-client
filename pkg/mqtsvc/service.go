package mqtsvc

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	mqtcore "github.com/tendant/marqeta-client/coreapi"
	"golang.org/x/exp/slog"
)

type MqtSvc struct {
	acl *mqtcore.APIClient
	ctx context.Context
}

func NewFromBasic(url, appToken, accessToken string) *MqtSvc {
	configuration := mqtcore.NewConfiguration(url)
	ctx := context.WithValue(context.Background(), mqtcore.ContextBasicAuth, mqtcore.BasicAuth{
		UserName: appToken,
		Password: accessToken,
	})
	acl := mqtcore.NewAPIClient(configuration)
	return &MqtSvc{acl, ctx}
}

type Tokens struct {
	BusinessToken                    string `json:"business_token"`
	ActingUserToken                  string `json:"acting_user_token"`
	CardToken                        string `json:"card_token"`
	PrecedingRelatedTransactionToken string `json:"preceding_related_transaction_token,omitempty"`
	CardProductToken                 string `json:"card_product_token,omitempty"`
	ApplicationToken                 string `json:"application_token,omitempty"`
}

func (m *MqtSvc) decodeRespError(r *http.Response) (*mqtcore.Error, error) {
	result := mqtcore.NewError()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	if err != nil {
		slog.Warn("Failed decoding:", "err", err)
		return nil, err
	} else {
		slog.Info("Got error from resp body!", "result", result)
		return result, nil
	}
}

func (m *MqtSvc) BusinessesInfo(tokens Tokens) (*mqtcore.BusinessCardHolderResponse, int, error) {
	fields := "status"
	slog.Info("BusinessInfo for", "Tokens", tokens, "fields", fields)
	resp, r, err := m.acl.BusinessesAPI.GetBusinessesToken(m.ctx, tokens.BusinessToken).Fields(fields).Execute()
	if err != nil {
		slog.Error("Error when calling `BusinesssApi.GetBusinesssToken`", "err", err)
		result, err := m.decodeRespError(r)
		if err != nil {
			return nil, r.StatusCode, err
		} else {
			return nil, r.StatusCode, errors.New(result.GetErrorCode())
		}
	}
	return resp, r.StatusCode, nil
}
