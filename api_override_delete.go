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


// OverrideDeleteApiService OverrideDeleteApi service
type OverrideDeleteApiService service

type ApiApisAppsK8smgmtIoV3ProjectsProjectOverridesNameDeleteRequest struct {
	ctx context.Context
	ApiService *OverrideDeleteApiService
	project string
	name string
}

func (r ApiApisAppsK8smgmtIoV3ProjectsProjectOverridesNameDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameDeleteExecute(r)
}

/*
ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameDelete Method for ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameDelete

Override Delete endpoint

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param project project of the resource
 @param name name of the resource
 @return ApiApisAppsK8smgmtIoV3ProjectsProjectOverridesNameDeleteRequest
*/
func (a *OverrideDeleteApiService) ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameDelete(ctx context.Context, project string, name string) ApiApisAppsK8smgmtIoV3ProjectsProjectOverridesNameDeleteRequest {
	return ApiApisAppsK8smgmtIoV3ProjectsProjectOverridesNameDeleteRequest{
		ApiService: a,
		ctx: ctx,
		project: project,
		name: name,
	}
}

// Execute executes the request
func (a *OverrideDeleteApiService) ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameDeleteExecute(r ApiApisAppsK8smgmtIoV3ProjectsProjectOverridesNameDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OverrideDeleteApiService.ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apis/apps.k8smgmt.io/v3/projects/{project}/overrides/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterValueToString(r.project, "project")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 502 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
