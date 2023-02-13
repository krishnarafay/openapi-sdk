/*
Kubernetes Operations APIs

Testing NamespaceListApiService

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

func Test_rafaysdkgov3_NamespaceListApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NamespaceListApiService ApisInfraK8smgmtIoV3ProjectsProjectNamespacesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string

		resp, httpRes, err := apiClient.NamespaceListApi.ApisInfraK8smgmtIoV3ProjectsProjectNamespacesGet(context.Background(), project).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}