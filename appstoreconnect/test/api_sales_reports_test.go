/*
App Store Connect API

Testing SalesReportsApiService

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

func Test_appstoreconnect_SalesReportsApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SalesReportsApiService SalesReportsGetCollection", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SalesReportsApi.SalesReportsGetCollection(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
