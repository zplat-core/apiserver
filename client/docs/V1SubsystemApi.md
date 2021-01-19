# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SubsystemsGet**](V1SubsystemApi.md#V1SubsystemsGet) | **Get** /v1/subsystems | 获取所有子系统
[**V1SubsystemsIdDelete**](V1SubsystemApi.md#V1SubsystemsIdDelete) | **Delete** /v1/subsystems/{id} | 删除子系统
[**V1SubsystemsIdPut**](V1SubsystemApi.md#V1SubsystemsIdPut) | **Put** /v1/subsystems/{id} | 修改子系统
[**V1SubsystemsPost**](V1SubsystemApi.md#V1SubsystemsPost) | **Post** /v1/subsystems | 添加子系统

# **V1SubsystemsGet**
> InlineResponse200 V1SubsystemsGet(ctx, )
获取所有子系统

获取所有子系统

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SubsystemsIdDelete**
> HttputilJsonResponse V1SubsystemsIdDelete(ctx, id)
删除子系统

删除子系统

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID | 

### Return type

[**HttputilJsonResponse**](httputil.JSONResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SubsystemsIdPut**
> HttputilJsonResponse V1SubsystemsIdPut(ctx, body, id)
修改子系统

修改子系统

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BindBodySubsystem**](BindBodySubsystem.md)| 参数 | 
  **id** | **string**| ID | 

### Return type

[**HttputilJsonResponse**](httputil.JSONResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SubsystemsPost**
> HttputilJsonResponse V1SubsystemsPost(ctx, body)
添加子系统

添加子系统

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BindBodySubsystem**](BindBodySubsystem.md)| 参数 | 

### Return type

[**HttputilJsonResponse**](httputil.JSONResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

