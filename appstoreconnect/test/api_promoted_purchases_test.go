/*
App Store Connect API

Testing PromotedPurchasesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package appstoreconnect

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	openapiclient "github.com/micnncim/appstoreconnect-go/appstoreconnect"
)

func Test_appstoreconnect_PromotedPurchasesApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PromotedPurchasesApiService PromotedPurchasesCreateInstance", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PromotedPurchasesApi.PromotedPurchasesCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PromotedPurchasesApiService PromotedPurchasesDeleteInstance", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.PromotedPurchasesApi.PromotedPurchasesDeleteInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PromotedPurchasesApiService PromotedPurchasesGetInstance", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.PromotedPurchasesApi.PromotedPurchasesGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PromotedPurchasesApiService PromotedPurchasesPromotionImagesGetToManyRelated", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.PromotedPurchasesApi.PromotedPurchasesPromotionImagesGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PromotedPurchasesApiService PromotedPurchasesUpdateInstance", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.PromotedPurchasesApi.PromotedPurchasesUpdateInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
