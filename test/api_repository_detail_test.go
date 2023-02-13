/*
Kubernetes Operations APIs

Testing RepositoryDetailApiService

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

func Test_rafaysdkgov3_RepositoryDetailApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoryDetailApiService ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesNameArtifactsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string
		var name string

		resp, httpRes, err := apiClient.RepositoryDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesNameArtifactsGet(context.Background(), project, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoryDetailApiService ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesNameGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string
		var name string

		resp, httpRes, err := apiClient.RepositoryDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesNameGet(context.Background(), project, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoryDetailApiService ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysNameArtifactsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string
		var name string

		resp, httpRes, err := apiClient.RepositoryDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysNameArtifactsGet(context.Background(), project, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoryDetailApiService ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysNameGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string
		var name string

		resp, httpRes, err := apiClient.RepositoryDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysNameGet(context.Background(), project, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}