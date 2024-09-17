/*
Stacks Blockchain API

Testing MempoolAPIService

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

func Test_stacks_blockchain_api_client_MempoolAPIService(t *testing.T) {

	configuration := stacks_blockchain_api_client.NewConfiguration()
	configuration.Servers = stacks_blockchain_api_client.ServerConfigurations{configuration.Servers[1]}
	apiClient := stacks_blockchain_api_client.NewAPIClient(configuration)

	t.Run("Test MempoolAPIService GetMempoolFeePriorities", func(t *testing.T) {
		resp, httpRes, err := apiClient.MempoolAPI.GetMempoolFeePriorities(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
