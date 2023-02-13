/*
Kubernetes Operations APIs

Kubernetes Operations APIs

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rafaysdkgov3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// MeshProfileApplyApiService MeshProfileApplyApi service
type MeshProfileApplyApiService service

type ApiApisServicemeshK8smgmtIoV3ProjectsProjectMeshprofilesPostRequest struct {
	ctx context.Context
	ApiService *MeshProfileApplyApiService
	project string
	meshProfile *MeshProfile
}

// schema of the resource to apply
func (r ApiApisServicemeshK8smgmtIoV3ProjectsProjectMeshprofilesPostRequest) MeshProfile(meshProfile MeshProfile) ApiApisServicemeshK8smgmtIoV3ProjectsProjectMeshprofilesPostRequest {
	r.meshProfile = &meshProfile
	return r
}

func (r ApiApisServicemeshK8smgmtIoV3ProjectsProjectMeshprofilesPostRequest) Execute() (*MeshProfile, *http.Response, error) {
	return r.ApiService.ApisServicemeshK8smgmtIoV3ProjectsProjectMeshprofilesPostExecute(r)
}

/*
ApisServicemeshK8smgmtIoV3ProjectsProjectMeshprofilesPost Method for ApisServicemeshK8smgmtIoV3ProjectsProjectMeshprofilesPost

MeshProfile Apply endpoint

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param project project of the resource
 @return ApiApisServicemeshK8smgmtIoV3ProjectsProjectMeshprofilesPostRequest
*/
func (a *MeshProfileApplyApiService) ApisServicemeshK8smgmtIoV3ProjectsProjectMeshprofilesPost(ctx context.Context, project string) ApiApisServicemeshK8smgmtIoV3ProjectsProjectMeshprofilesPostRequest {
	return ApiApisServicemeshK8smgmtIoV3ProjectsProjectMeshprofilesPostRequest{
		ApiService: a,
		ctx: ctx,
		project: project,
	}
}

// Execute executes the request
//  @return MeshProfile
func (a *MeshProfileApplyApiService) ApisServicemeshK8smgmtIoV3ProjectsProjectMeshprofilesPostExecute(r ApiApisServicemeshK8smgmtIoV3ProjectsProjectMeshprofilesPostRequest) (*MeshProfile, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MeshProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MeshProfileApplyApiService.ApisServicemeshK8smgmtIoV3ProjectsProjectMeshprofilesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apis/servicemesh.k8smgmt.io/v3/projects/{project}/meshprofiles"
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterValueToString(r.project, "project")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.meshProfile == nil {
		return localVarReturnValue, nil, reportError("meshProfile is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/yaml"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/yaml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.meshProfile
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-KEY"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 502 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v MeshProfile
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
