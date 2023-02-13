/*
Kubernetes Operations APIs

Testing OPAPolicyApplyApiService

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

func Test_rafaysdkgov3_OPAPolicyApplyApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OPAPolicyApplyApiService ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string

		resp, httpRes, err := apiClient.OPAPolicyApplyApi.ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesPost(context.Background(), project).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OPAPolicyApplyApiService ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string

		resp, httpRes, err := apiClient.OPAPolicyApplyApi.ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysPost(context.Background(), project).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}