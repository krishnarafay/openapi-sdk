/*
Kubernetes Operations APIs

Testing ChargebackShareDetailApiService

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

func Test_rafaysdkgov3_ChargebackShareDetailApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ChargebackShareDetailApiService ApisSystemK8smgmtIoV3ChargebacksharesNameGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.ChargebackShareDetailApi.ApisSystemK8smgmtIoV3ChargebacksharesNameGet(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
