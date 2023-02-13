/*
Kubernetes Operations APIs

Testing OverrideDetailApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package rafaysdkgov3

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_rafaysdkgov3_OverrideDetailApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OverrideDetailApiService ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameArtifactsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string
		var name string

		resp, httpRes, err := apiClient.OverrideDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameArtifactsGet(context.Background(), project, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OverrideDetailApiService ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string
		var name string

		resp, httpRes, err := apiClient.OverrideDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameGet(context.Background(), project, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
