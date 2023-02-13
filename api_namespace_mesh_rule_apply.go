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


// NamespaceMeshRuleApplyApiService NamespaceMeshRuleApplyApi service
type NamespaceMeshRuleApplyApiService service

type ApiApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPostRequest struct {
	ctx context.Context
	ApiService *NamespaceMeshRuleApplyApiService
	project string
	namespaceMeshRule *NamespaceMeshRule
}

// schema of the resource to apply
func (r ApiApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPostRequest) NamespaceMeshRule(namespaceMeshRule NamespaceMeshRule) ApiApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPostRequest {
	r.namespaceMeshRule = &namespaceMeshRule
	return r
}

func (r ApiApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPostRequest) Execute() (*NamespaceMeshRule, *http.Response, error) {
	return r.ApiService.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPostExecute(r)
}

/*
ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPost Method for ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPost

NamespaceMeshRule Apply endpoint

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param project project of the resource
 @return ApiApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPostRequest
*/
func (a *NamespaceMeshRuleApplyApiService) ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPost(ctx context.Context, project string) ApiApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPostRequest {
	return ApiApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPostRequest{
		ApiService: a,
		ctx: ctx,
		project: project,
	}
}

// Execute executes the request
//  @return NamespaceMeshRule
func (a *NamespaceMeshRuleApplyApiService) ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPostExecute(r ApiApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPostRequest) (*NamespaceMeshRule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NamespaceMeshRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamespaceMeshRuleApplyApiService.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apis/servicemesh.k8smgmt.io/v3/projects/{project}/namespacemeshrules"
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterValueToString(r.project, "project")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.namespaceMeshRule == nil {
		return localVarReturnValue, nil, reportError("namespaceMeshRule is required and must be specified")
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
	localVarPostBody = r.namespaceMeshRule
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
			var v NamespaceMeshRule
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
