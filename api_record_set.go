
/*
 * STACKIT DNS API
 *
 * This api provides dns
 *
 * API version: 1.0
 * Contact: stackit-dns@mail.schwarz
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type RecordSetApiService service
/*
RecordSetApiService All get selected RRSets
All RRSet
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectId project id
 * @param zoneId zone id
 * @param optional nil or *RecordSetApiV1ProjectsProjectIdZonesZoneIdRrsetsGetOpts - Optional Parameters:
     * @param "Page" (optional.Int32) -  page
     * @param "PageSize" (optional.Int32) -  page size
     * @param "NameEq" (optional.String) -  filter name equal
     * @param "NameLike" (optional.String) -  filter name like
     * @param "TypeEq" (optional.String) -  filter type
     * @param "ActiveEq" (optional.Bool) -  filter active equal
     * @param "CreationStartedGt" (optional.String) -  filter creation started greater with utc timestamp
     * @param "CreationStartedLt" (optional.String) -  filter creation started lesser with utc timestamp
     * @param "CreationStartedGte" (optional.String) -  filter creation started greater equal with utc timestamp
     * @param "CreationStartedLte" (optional.String) -  filter creation started lesser equal with utc timestamp
     * @param "CreationFinishedGt" (optional.String) -  filter creation finished greater with utc timestamp
     * @param "CreationFinishedLt" (optional.String) -  filter creation finished lesser with utc timestamp
     * @param "CreationFinishedGte" (optional.String) -  filter creation finished greater equal with utc timestamp
     * @param "CreationFinishedLte" (optional.String) -  filter creation finished lesser equal with utc timestamp
     * @param "UpdateStartedGt" (optional.String) -  filter update started greater with utc timestamp
     * @param "UpdateStartedLt" (optional.String) -  filter update started lesser with utc timestamp
     * @param "UpdateStartedGte" (optional.String) -  filter update started greater equal with utc timestamp
     * @param "UpdateStartedLte" (optional.String) -  filter update started lesser equal with utc timestamp
     * @param "UpdateFinishedGt" (optional.String) -  filter update finished greater with utc timestamp
     * @param "UpdateFinishedLt" (optional.String) -  filter update finished lesser with utc timestamp
     * @param "UpdateFinishedGte" (optional.String) -  filter update finished greater equal with utc timestamp
     * @param "UpdateFinishedLte" (optional.String) -  filter update finished lesser equal with utc timestamp
     * @param "OrderByName" (optional.String) -  order by name
     * @param "OrderByCreationStarted" (optional.String) -  order by creationStarted
     * @param "OrderByCreationFinished" (optional.String) -  order by creationFinished
     * @param "OrderByUpdateStarted" (optional.String) -  order by updateStarted
     * @param "OrderByUpdateFinished" (optional.String) -  order by updateFinished
@return RrsetResponseRrSetAll
*/

type RecordSetApiV1ProjectsProjectIdZonesZoneIdRrsetsGetOpts struct {
    Page optional.Int32
    PageSize optional.Int32
    NameEq optional.String
    NameLike optional.String
    TypeEq optional.String
    ActiveEq optional.Bool
    CreationStartedGt optional.String
    CreationStartedLt optional.String
    CreationStartedGte optional.String
    CreationStartedLte optional.String
    CreationFinishedGt optional.String
    CreationFinishedLt optional.String
    CreationFinishedGte optional.String
    CreationFinishedLte optional.String
    UpdateStartedGt optional.String
    UpdateStartedLt optional.String
    UpdateStartedGte optional.String
    UpdateStartedLte optional.String
    UpdateFinishedGt optional.String
    UpdateFinishedLt optional.String
    UpdateFinishedGte optional.String
    UpdateFinishedLte optional.String
    OrderByName optional.String
    OrderByCreationStarted optional.String
    OrderByCreationFinished optional.String
    OrderByUpdateStarted optional.String
    OrderByUpdateFinished optional.String
}

func (a *RecordSetApiService) V1ProjectsProjectIdZonesZoneIdRrsetsGet(ctx context.Context, projectId string, zoneId string, localVarOptionals *RecordSetApiV1ProjectsProjectIdZonesZoneIdRrsetsGetOpts) (RrsetResponseRrSetAll, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RrsetResponseRrSetAll
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/projects/{projectId}/zones/{zoneId}/rrsets"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", fmt.Sprintf("%v", projectId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"zoneId"+"}", fmt.Sprintf("%v", zoneId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("pageSize", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NameEq.IsSet() {
		localVarQueryParams.Add("name[eq]", parameterToString(localVarOptionals.NameEq.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NameLike.IsSet() {
		localVarQueryParams.Add("name[like]", parameterToString(localVarOptionals.NameLike.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TypeEq.IsSet() {
		localVarQueryParams.Add("type[eq]", parameterToString(localVarOptionals.TypeEq.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ActiveEq.IsSet() {
		localVarQueryParams.Add("active[eq]", parameterToString(localVarOptionals.ActiveEq.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CreationStartedGt.IsSet() {
		localVarQueryParams.Add("creationStarted[gt]", parameterToString(localVarOptionals.CreationStartedGt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CreationStartedLt.IsSet() {
		localVarQueryParams.Add("creationStarted[lt]", parameterToString(localVarOptionals.CreationStartedLt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CreationStartedGte.IsSet() {
		localVarQueryParams.Add("creationStarted[gte]", parameterToString(localVarOptionals.CreationStartedGte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CreationStartedLte.IsSet() {
		localVarQueryParams.Add("creationStarted[lte]", parameterToString(localVarOptionals.CreationStartedLte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CreationFinishedGt.IsSet() {
		localVarQueryParams.Add("creationFinished[gt]", parameterToString(localVarOptionals.CreationFinishedGt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CreationFinishedLt.IsSet() {
		localVarQueryParams.Add("creationFinished[lt]", parameterToString(localVarOptionals.CreationFinishedLt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CreationFinishedGte.IsSet() {
		localVarQueryParams.Add("creationFinished[gte]", parameterToString(localVarOptionals.CreationFinishedGte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CreationFinishedLte.IsSet() {
		localVarQueryParams.Add("creationFinished[lte]", parameterToString(localVarOptionals.CreationFinishedLte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UpdateStartedGt.IsSet() {
		localVarQueryParams.Add("updateStarted[gt]", parameterToString(localVarOptionals.UpdateStartedGt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UpdateStartedLt.IsSet() {
		localVarQueryParams.Add("updateStarted[lt]", parameterToString(localVarOptionals.UpdateStartedLt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UpdateStartedGte.IsSet() {
		localVarQueryParams.Add("updateStarted[gte]", parameterToString(localVarOptionals.UpdateStartedGte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UpdateStartedLte.IsSet() {
		localVarQueryParams.Add("updateStarted[lte]", parameterToString(localVarOptionals.UpdateStartedLte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UpdateFinishedGt.IsSet() {
		localVarQueryParams.Add("updateFinished[gt]", parameterToString(localVarOptionals.UpdateFinishedGt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UpdateFinishedLt.IsSet() {
		localVarQueryParams.Add("updateFinished[lt]", parameterToString(localVarOptionals.UpdateFinishedLt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UpdateFinishedGte.IsSet() {
		localVarQueryParams.Add("updateFinished[gte]", parameterToString(localVarOptionals.UpdateFinishedGte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UpdateFinishedLte.IsSet() {
		localVarQueryParams.Add("updateFinished[lte]", parameterToString(localVarOptionals.UpdateFinishedLte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrderByName.IsSet() {
		localVarQueryParams.Add("orderBy[name]", parameterToString(localVarOptionals.OrderByName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrderByCreationStarted.IsSet() {
		localVarQueryParams.Add("orderBy[creationStarted]", parameterToString(localVarOptionals.OrderByCreationStarted.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrderByCreationFinished.IsSet() {
		localVarQueryParams.Add("orderBy[creationFinished]", parameterToString(localVarOptionals.OrderByCreationFinished.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrderByUpdateStarted.IsSet() {
		localVarQueryParams.Add("orderBy[updateStarted]", parameterToString(localVarOptionals.OrderByUpdateStarted.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrderByUpdateFinished.IsSet() {
		localVarQueryParams.Add("orderBy[updateFinished]", parameterToString(localVarOptionals.OrderByUpdateFinished.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v RrsetResponseRrSetAll
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v SerializerMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 502 {
			var v SerializerMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
RecordSetApiService Post record set
Post record set
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body record set to create
 * @param projectId project id
 * @param zoneId zone id
@return RrsetResponseRrSet
*/
func (a *RecordSetApiService) V1ProjectsProjectIdZonesZoneIdRrsetsPost(ctx context.Context, body RrsetRrSetPost, projectId string, zoneId string) (RrsetResponseRrSet, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RrsetResponseRrSet
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/projects/{projectId}/zones/{zoneId}/rrsets"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", fmt.Sprintf("%v", projectId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"zoneId"+"}", fmt.Sprintf("%v", zoneId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 202 {
			var v RrsetResponseRrSet
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v SerializerMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v SerializerMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 502 {
			var v SerializerMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
RecordSetApiService Delete a record set
Delete a record set
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectId project id
 * @param zoneId zone id
 * @param rrSetId record set id
@return SerializerMessage
*/
func (a *RecordSetApiService) V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdDelete(ctx context.Context, projectId string, zoneId string, rrSetId string) (SerializerMessage, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue SerializerMessage
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/projects/{projectId}/zones/{zoneId}/rrsets/{rrSetId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", fmt.Sprintf("%v", projectId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"zoneId"+"}", fmt.Sprintf("%v", zoneId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rrSetId"+"}", fmt.Sprintf("%v", rrSetId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 202 {
			var v SerializerMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v SerializerMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v SerializerMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 502 {
			var v SerializerMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
RecordSetApiService Get a single rrset
Get rrset
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectId project id
 * @param zoneId zone id
 * @param rrSetId record set id
@return RrsetResponseRrSet
*/
func (a *RecordSetApiService) V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdGet(ctx context.Context, projectId string, zoneId string, rrSetId string) (RrsetResponseRrSet, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RrsetResponseRrSet
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/projects/{projectId}/zones/{zoneId}/rrsets/{rrSetId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", fmt.Sprintf("%v", projectId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"zoneId"+"}", fmt.Sprintf("%v", zoneId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rrSetId"+"}", fmt.Sprintf("%v", rrSetId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v RrsetResponseRrSet
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v SerializerMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 502 {
			var v SerializerMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
RecordSetApiService Patch updates a record set
Patch record set
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body record set to patch
 * @param projectId project id
 * @param zoneId zone id
 * @param rrSetId record set id
@return RrsetResponseRrSet
*/
func (a *RecordSetApiService) V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdPatch(ctx context.Context, body RrsetRrSetPatch, projectId string, zoneId string, rrSetId string) (RrsetResponseRrSet, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Patch")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RrsetResponseRrSet
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/projects/{projectId}/zones/{zoneId}/rrsets/{rrSetId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", fmt.Sprintf("%v", projectId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"zoneId"+"}", fmt.Sprintf("%v", zoneId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rrSetId"+"}", fmt.Sprintf("%v", rrSetId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 202 {
			var v RrsetResponseRrSet
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v SerializerMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v SerializerMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 502 {
			var v SerializerMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
RecordSetApiService PatchRecords updates a record in a rrset
PatchRecords rrset updates a record in a rrset
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body rrset to update
 * @param projectId project id
 * @param zoneId zone id
 * @param rrSetId record set id
@return RrsetResponseRrSet
*/
func (a *RecordSetApiService) V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdRecordsPatch(ctx context.Context, body RrsetRecordPatch, projectId string, zoneId string, rrSetId string) (RrsetResponseRrSet, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Patch")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RrsetResponseRrSet
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/projects/{projectId}/zones/{zoneId}/rrsets/{rrSetId}/records"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", fmt.Sprintf("%v", projectId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"zoneId"+"}", fmt.Sprintf("%v", zoneId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rrSetId"+"}", fmt.Sprintf("%v", rrSetId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 202 {
			var v RrsetResponseRrSet
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v SerializerMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v SerializerMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 502 {
			var v SerializerMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
RecordSetApiService Restore record set
Restore record set
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectId project id
 * @param zoneId zone id
 * @param rrSetId record set id
@return SerializerMessage
*/
func (a *RecordSetApiService) V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdRestoresPost(ctx context.Context, projectId string, zoneId string, rrSetId string) (SerializerMessage, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue SerializerMessage
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/projects/{projectId}/zones/{zoneId}/rrsets/{rrSetId}/restores"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", fmt.Sprintf("%v", projectId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"zoneId"+"}", fmt.Sprintf("%v", zoneId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rrSetId"+"}", fmt.Sprintf("%v", rrSetId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 202 {
			var v SerializerMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v SerializerMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 502 {
			var v SerializerMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}