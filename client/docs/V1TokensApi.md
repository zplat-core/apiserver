# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TokensDelete**](V1TokensApi.md#V1TokensDelete) | **Delete** /v1/tokens | 退出登录
[**V1TokensPost**](V1TokensApi.md#V1TokensPost) | **Post** /v1/tokens | 登录/密码重置

# **V1TokensDelete**
> HttputilJsonResponse V1TokensDelete(ctx, )
退出登录

用户状态登出

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**HttputilJsonResponse**](httputil.JSONResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TokensPost**
> HttputilJsonResponse V1TokensPost(ctx, body)
登录/密码重置

用于账户登录和申请密码重置

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BindBodyToken**](BindBodyToken.md)| 参数 | 

### Return type

[**HttputilJsonResponse**](httputil.JSONResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

