/*
Stacks 2.0+ RPC API

Testing AccountsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package rpc_client

import (
	"context"
	"testing"

	"github.com/icon-project/stacks-go-sdk/pkg/rpc_client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_rpc_client_AccountsAPIService(t *testing.T) {

	configuration := rpc_client.NewConfiguration()
	apiClient := rpc_client.NewAPIClient(configuration)

	t.Run("Test AccountsAPIService GetAccountInfo", func(t *testing.T) {
		principal := "ST15C893XJFJ6FSKM020P9JQDB5T7X6MQTXMBPAVH"

		ctx := context.WithValue(context.Background(), rpc_client.ContextServerIndex, 1)

		resp, httpRes, err := apiClient.AccountsAPI.GetAccountInfo(ctx, principal).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
