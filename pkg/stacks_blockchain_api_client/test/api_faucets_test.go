/*
Stacks Blockchain API

Testing FaucetsAPIService

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

func Test_stacks_blockchain_api_client_FaucetsAPIService(t *testing.T) {

	configuration := stacks_blockchain_api_client.NewConfiguration()
	apiClient := stacks_blockchain_api_client.NewAPIClient(configuration)

	t.Run("Test FaucetsAPIService GetBtcBalance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var address string

		resp, httpRes, err := apiClient.FaucetsAPI.GetBtcBalance(context.Background(), address).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FaucetsAPIService RunFaucetBtc", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FaucetsAPI.RunFaucetBtc(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FaucetsAPIService RunFaucetStx", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FaucetsAPI.RunFaucetStx(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
