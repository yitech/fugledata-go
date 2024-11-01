/*
standard public schema

Testing InventoryAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package fugledata

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/yitech/fugledata-go"
)

func Test_fugledata_InventoryAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InventoryAPIService InventoryDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.InventoryAPI.InventoryDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryAPIService InventoryGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InventoryAPI.InventoryGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryAPIService InventoryPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.InventoryAPI.InventoryPatch(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryAPIService InventoryPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.InventoryAPI.InventoryPost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
