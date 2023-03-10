/*
Kubernetes Operations APIs

Testing AddonDeleteApiService

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

func Test_rafaysdkgov3_AddonDeleteApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AddonDeleteApiService ApisInfraK8smgmtIoV3ProjectsProjectAddonsNameDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string
		var name string

		httpRes, err := apiClient.AddonDeleteApi.ApisInfraK8smgmtIoV3ProjectsProjectAddonsNameDelete(context.Background(), project, name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
