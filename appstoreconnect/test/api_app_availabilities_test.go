/*
App Store Connect API

Testing AppAvailabilitiesApiService

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

func Test_appstoreconnect_AppAvailabilitiesApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AppAvailabilitiesApiService AppAvailabilitiesAvailableTerritoriesGetToManyRelated", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppAvailabilitiesApi.AppAvailabilitiesAvailableTerritoriesGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AppAvailabilitiesApiService AppAvailabilitiesCreateInstance", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AppAvailabilitiesApi.AppAvailabilitiesCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AppAvailabilitiesApiService AppAvailabilitiesGetInstance", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppAvailabilitiesApi.AppAvailabilitiesGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AppAvailabilitiesApiService AppAvailabilitiesV2CreateInstance", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AppAvailabilitiesApi.AppAvailabilitiesV2CreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AppAvailabilitiesApiService AppAvailabilitiesV2GetInstance", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppAvailabilitiesApi.AppAvailabilitiesV2GetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AppAvailabilitiesApiService AppAvailabilitiesV2TerritoryAvailabilitiesGetToManyRelated", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppAvailabilitiesApi.AppAvailabilitiesV2TerritoryAvailabilitiesGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
