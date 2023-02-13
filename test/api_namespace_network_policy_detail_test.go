/*
Kubernetes Operations APIs

Testing NamespaceNetworkPolicyDetailApiService

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

func Test_rafaysdkgov3_NamespaceNetworkPolicyDetailApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NamespaceNetworkPolicyDetailApiService ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesNameGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string
		var name string

		resp, httpRes, err := apiClient.NamespaceNetworkPolicyDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesNameGet(context.Background(), project, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespaceNetworkPolicyDetailApiService ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysNameGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string
		var name string

		resp, httpRes, err := apiClient.NamespaceNetworkPolicyDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysNameGet(context.Background(), project, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
