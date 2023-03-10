/*
Kubernetes Operations APIs

Testing ContainerRegistryApplyApiService

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

func Test_rafaysdkgov3_ContainerRegistryApplyApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContainerRegistryApplyApiService ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistriesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string

		resp, httpRes, err := apiClient.ContainerRegistryApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistriesPost(context.Background(), project).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerRegistryApplyApiService ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistrysPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string

		resp, httpRes, err := apiClient.ContainerRegistryApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistrysPost(context.Background(), project).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
