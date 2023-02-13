/*
Kubernetes Operations APIs

Testing PipelineDeleteApiService

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

func Test_rafaysdkgov3_PipelineDeleteApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PipelineDeleteApiService ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string
		var name string

		httpRes, err := apiClient.PipelineDeleteApi.ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameDelete(context.Background(), project, name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
