/*
Core API

Marqeta's Core API endpoints, conveniently annotated to enable code generation (including SDKs), test cases, and documentation. Currently in beta.

API version: 3.0.11
Contact: support@marqeta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marqeta_coreapi_client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// AccountRewardsAPIService AccountRewardsAPI service
type AccountRewardsAPIService service

type ApiCreateRewardRequest struct {
	ctx context.Context
	ApiService *AccountRewardsAPIService
	accountToken string
	rewardCreateReq *RewardCreateReq
}

func (r ApiCreateRewardRequest) RewardCreateReq(rewardCreateReq RewardCreateReq) ApiCreateRewardRequest {
	r.rewardCreateReq = &rewardCreateReq
	return r
}

func (r ApiCreateRewardRequest) Execute() (*RewardResponse, *http.Response, error) {
	return r.ApiService.CreateRewardExecute(r)
}

/*
CreateReward Create account reward

Create a reward for an existing credit account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountToken The unique identifier of the credit account for which you want to create a reward.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
 @return ApiCreateRewardRequest
*/
func (a *AccountRewardsAPIService) CreateReward(ctx context.Context, accountToken string) ApiCreateRewardRequest {
	return ApiCreateRewardRequest{
		ApiService: a,
		ctx: ctx,
		accountToken: accountToken,
	}
}

// Execute executes the request
//  @return RewardResponse
func (a *AccountRewardsAPIService) CreateRewardExecute(r ApiCreateRewardRequest) (*RewardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RewardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountRewardsAPIService.CreateReward")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit/accounts/{account_token}/rewards"
	localVarPath = strings.Replace(localVarPath, "{"+"account_token"+"}", url.PathEscape(parameterValueToString(r.accountToken, "accountToken")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rewardCreateReq == nil {
		return localVarReturnValue, nil, reportError("rewardCreateReq is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.rewardCreateReq
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
