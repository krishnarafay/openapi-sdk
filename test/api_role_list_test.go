/*
Kubernetes Operations APIs

Testing RoleListApiService

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

func Test_rafaysdkgov3_RoleListApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RoleListApiService ApisSystemK8smgmtIoV3RolesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RoleListApi.ApisSystemK8smgmtIoV3RolesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
