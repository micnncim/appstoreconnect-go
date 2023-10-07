/*
App Store Connect API

Testing InAppPurchaseAvailabilitiesApiService

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

func Test_appstoreconnect_InAppPurchaseAvailabilitiesApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InAppPurchaseAvailabilitiesApiService InAppPurchaseAvailabilitiesAvailableTerritoriesGetToManyRelated", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.InAppPurchaseAvailabilitiesApi.InAppPurchaseAvailabilitiesAvailableTerritoriesGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test InAppPurchaseAvailabilitiesApiService InAppPurchaseAvailabilitiesCreateInstance", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.InAppPurchaseAvailabilitiesApi.InAppPurchaseAvailabilitiesCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test InAppPurchaseAvailabilitiesApiService InAppPurchaseAvailabilitiesGetInstance", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.InAppPurchaseAvailabilitiesApi.InAppPurchaseAvailabilitiesGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
