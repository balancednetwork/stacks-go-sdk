/*
Stacks Blockchain API

Testing FungibleTokensAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package stacks_blockchain_api_client

import (
	"context"
	"testing"

	"github.com/icon-project/stacks-go-sdk/pkg/stacks_blockchain_api_client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_stacks_blockchain_api_client_FungibleTokensAPIService(t *testing.T) {

	configuration := stacks_blockchain_api_client.NewConfiguration()
	apiClient := stacks_blockchain_api_client.NewAPIClient(configuration)

	t.Run("Test FungibleTokensAPIService GetFtHolders", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var token string

		resp, httpRes, err := apiClient.FungibleTokensAPI.GetFtHolders(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
