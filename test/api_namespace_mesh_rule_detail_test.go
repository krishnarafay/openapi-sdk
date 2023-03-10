/*
Kubernetes Operations APIs

Testing NamespaceMeshRuleDetailApiService

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

func Test_rafaysdkgov3_NamespaceMeshRuleDetailApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NamespaceMeshRuleDetailApiService ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesNameArtifactsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string
		var name string

		resp, httpRes, err := apiClient.NamespaceMeshRuleDetailApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesNameArtifactsGet(context.Background(), project, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespaceMeshRuleDetailApiService ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesNameGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string
		var name string

		resp, httpRes, err := apiClient.NamespaceMeshRuleDetailApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesNameGet(context.Background(), project, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
