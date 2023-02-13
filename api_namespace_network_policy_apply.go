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


// NamespaceNetworkPolicyApplyApiService NamespaceNetworkPolicyApplyApi service
type NamespaceNetworkPolicyApplyApiService service

type ApiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPostRequest struct {
	ctx context.Context
	ApiService *NamespaceNetworkPolicyApplyApiService
	project string
	namespaceNetworkPolicy *NamespaceNetworkPolicy
}

// schema of the resource to apply
func (r ApiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPostRequest) NamespaceNetworkPolicy(namespaceNetworkPolicy NamespaceNetworkPolicy) ApiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPostRequest {
	r.namespaceNetworkPolicy = &namespaceNetworkPolicy
	return r
}

func (r ApiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPostRequest) Execute() (*NamespaceNetworkPolicy, *http.Response, error) {
	return r.ApiService.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPostExecute(r)
}

/*
ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPost Method for ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPost

NamespaceNetworkPolicy Apply endpoint

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param project project of the resource
 @return ApiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPostRequest
*/
func (a *NamespaceNetworkPolicyApplyApiService) ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPost(ctx context.Context, project string) ApiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPostRequest {
	return ApiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPostRequest{
		ApiService: a,
		ctx: ctx,
		project: project,
	}
}

// Execute executes the request
//  @return NamespaceNetworkPolicy
func (a *NamespaceNetworkPolicyApplyApiService) ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPostExecute(r ApiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPostRequest) (*NamespaceNetworkPolicy, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NamespaceNetworkPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamespaceNetworkPolicyApplyApiService.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apis/security.k8smgmt.io/v3/projects/{project}/namespacenetworkpolicies"
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterValueToString(r.project, "project")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.namespaceNetworkPolicy == nil {
		return localVarReturnValue, nil, reportError("namespaceNetworkPolicy is required and must be specified")
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
	localVarPostBody = r.namespaceNetworkPolicy
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
			var v NamespaceNetworkPolicy
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

type ApiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPostRequest struct {
	ctx context.Context
	ApiService *NamespaceNetworkPolicyApplyApiService
	project string
	namespaceNetworkPolicy *NamespaceNetworkPolicy
}

// schema of the resource to apply
func (r ApiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPostRequest) NamespaceNetworkPolicy(namespaceNetworkPolicy NamespaceNetworkPolicy) ApiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPostRequest {
	r.namespaceNetworkPolicy = &namespaceNetworkPolicy
	return r
}

func (r ApiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPostRequest) Execute() (*NamespaceNetworkPolicy, *http.Response, error) {
	return r.ApiService.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPostExecute(r)
}

/*
ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPost Method for ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPost

NamespaceNetworkPolicy Apply endpoint

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param project project of the resource
 @return ApiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPostRequest
*/
func (a *NamespaceNetworkPolicyApplyApiService) ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPost(ctx context.Context, project string) ApiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPostRequest {
	return ApiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPostRequest{
		ApiService: a,
		ctx: ctx,
		project: project,
	}
}

// Execute executes the request
//  @return NamespaceNetworkPolicy
func (a *NamespaceNetworkPolicyApplyApiService) ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPostExecute(r ApiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPostRequest) (*NamespaceNetworkPolicy, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NamespaceNetworkPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamespaceNetworkPolicyApplyApiService.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apis/security.k8smgmt.io/v3/projects/{project}/namespacenetworkpolicys"
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterValueToString(r.project, "project")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.namespaceNetworkPolicy == nil {
		return localVarReturnValue, nil, reportError("namespaceNetworkPolicy is required and must be specified")
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
	localVarPostBody = r.namespaceNetworkPolicy
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
			var v NamespaceNetworkPolicy
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
