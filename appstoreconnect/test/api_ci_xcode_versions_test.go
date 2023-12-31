/*
App Store Connect API

Testing CiXcodeVersionsApiService

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

func Test_appstoreconnect_CiXcodeVersionsApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CiXcodeVersionsApiService CiXcodeVersionsGetCollection", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CiXcodeVersionsApi.CiXcodeVersionsGetCollection(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test CiXcodeVersionsApiService CiXcodeVersionsGetInstance", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.CiXcodeVersionsApi.CiXcodeVersionsGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test CiXcodeVersionsApiService CiXcodeVersionsMacOsVersionsGetToManyRelated", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.CiXcodeVersionsApi.CiXcodeVersionsMacOsVersionsGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
