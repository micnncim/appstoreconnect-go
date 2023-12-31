/*
App Store Connect API

Testing GameCenterAchievementsApiService

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

func Test_appstoreconnect_GameCenterAchievementsApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GameCenterAchievementsApiService GameCenterAchievementsCreateInstance", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GameCenterAchievementsApi.GameCenterAchievementsCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test GameCenterAchievementsApiService GameCenterAchievementsDeleteInstance", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.GameCenterAchievementsApi.GameCenterAchievementsDeleteInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test GameCenterAchievementsApiService GameCenterAchievementsGetInstance", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterAchievementsApi.GameCenterAchievementsGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test GameCenterAchievementsApiService GameCenterAchievementsGroupAchievementGetToOneRelated", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterAchievementsApi.GameCenterAchievementsGroupAchievementGetToOneRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test GameCenterAchievementsApiService GameCenterAchievementsGroupAchievementGetToOneRelationship", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterAchievementsApi.GameCenterAchievementsGroupAchievementGetToOneRelationship(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test GameCenterAchievementsApiService GameCenterAchievementsGroupAchievementUpdateToOneRelationship", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.GameCenterAchievementsApi.GameCenterAchievementsGroupAchievementUpdateToOneRelationship(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test GameCenterAchievementsApiService GameCenterAchievementsLocalizationsGetToManyRelated", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterAchievementsApi.GameCenterAchievementsLocalizationsGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test GameCenterAchievementsApiService GameCenterAchievementsReleasesGetToManyRelated", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterAchievementsApi.GameCenterAchievementsReleasesGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test GameCenterAchievementsApiService GameCenterAchievementsUpdateInstance", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterAchievementsApi.GameCenterAchievementsUpdateInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
