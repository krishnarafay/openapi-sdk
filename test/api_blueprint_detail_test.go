/*
Kubernetes Operations APIs

Testing BlueprintDetailApiService

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

func Test_rafaysdkgov3_BlueprintDetailApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BlueprintDetailApiService ApisInfraK8smgmtIoV3ProjectsProjectBlueprintsNameGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string
		var name string

		resp, httpRes, err := apiClient.BlueprintDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectBlueprintsNameGet(context.Background(), project, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
