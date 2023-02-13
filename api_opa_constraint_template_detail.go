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


// OPAConstraintTemplateDetailApiService OPAConstraintTemplateDetailApi service
type OPAConstraintTemplateDetailApiService service

type ApiApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameArtifactsGetRequest struct {
	ctx context.Context
	ApiService *OPAConstraintTemplateDetailApiService
	project string
	name string
}

func (r ApiApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameArtifactsGetRequest) Execute() (*OPAConstraintTemplate, *http.Response, error) {
	return r.ApiService.ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameArtifactsGetExecute(r)
}

/*
ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameArtifactsGet Method for ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameArtifactsGet

OPAConstraintTemplate Detail endpoint

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param project project of the resource
 @param name name of the resource
 @return ApiApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameArtifactsGetRequest
*/
func (a *OPAConstraintTemplateDetailApiService) ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameArtifactsGet(ctx context.Context, project string, name string) ApiApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameArtifactsGetRequest {
	return ApiApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameArtifactsGetRequest{
		ApiService: a,
		ctx: ctx,
		project: project,
		name: name,
	}
}

// Execute executes the request
//  @return OPAConstraintTemplate
func (a *OPAConstraintTemplateDetailApiService) ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameArtifactsGetExecute(r ApiApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameArtifactsGetRequest) (*OPAConstraintTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OPAConstraintTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OPAConstraintTemplateDetailApiService.ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameArtifactsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apis/opa.k8smgmt.io/v3/projects/{project}/opaconstrainttemplates/{name}/artifacts"
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
		if localVarHTTPResponse.StatusCode == 404 {
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
			var v OPAConstraintTemplate
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

type ApiApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameGetRequest struct {
	ctx context.Context
	ApiService *OPAConstraintTemplateDetailApiService
	project string
	name string
}

func (r ApiApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameGetRequest) Execute() (*OPAConstraintTemplate, *http.Response, error) {
	return r.ApiService.ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameGetExecute(r)
}

/*
ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameGet Method for ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameGet

OPAConstraintTemplate Detail endpoint

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param project project of the resource
 @param name name of the resource
 @return ApiApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameGetRequest
*/
func (a *OPAConstraintTemplateDetailApiService) ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameGet(ctx context.Context, project string, name string) ApiApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameGetRequest {
	return ApiApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameGetRequest{
		ApiService: a,
		ctx: ctx,
		project: project,
		name: name,
	}
}

// Execute executes the request
//  @return OPAConstraintTemplate
func (a *OPAConstraintTemplateDetailApiService) ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameGetExecute(r ApiApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameGetRequest) (*OPAConstraintTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OPAConstraintTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OPAConstraintTemplateDetailApiService.ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apis/opa.k8smgmt.io/v3/projects/{project}/opaconstrainttemplates/{name}"
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
		if localVarHTTPResponse.StatusCode == 404 {
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
			var v OPAConstraintTemplate
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
