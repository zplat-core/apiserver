# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1UserGet**](V1UsersApi.md#V1UserGet) | **Get** /v1/user | 当前登录用户信息
[**V1UserPasswordPut**](V1UsersApi.md#V1UserPasswordPut) | **Put** /v1/user/password | 修改登录用户密码
[**V1UserProfilePut**](V1UsersApi.md#V1UserProfilePut) | **Put** /v1/user/profile | 修改个人信息
[**V1UsersEmailPatch**](V1UsersApi.md#V1UsersEmailPatch) | **Patch** /v1/users/{email} | 更新一项用户信息
[**V1UsersGet**](V1UsersApi.md#V1UsersGet) | **Get** /v1/users | 用户列表
[**V1UsersPost**](V1UsersApi.md#V1UsersPost) | **Post** /v1/users | 用户注册
[**V1UsersUsernameGet**](V1UsersApi.md#V1UsersUsernameGet) | **Get** /v1/users/{username} | 用户查询

# **V1UserGet**
> InlineResponse200 V1UserGet(ctx, )
当前登录用户信息

获取已登录用户的详细信息

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

# **V1UserPasswordPut**
> HttputilJsonResponse V1UserPasswordPut(ctx, body)
修改登录用户密码

修改登录用户密码

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BindBodyUserPassword**](BindBodyUserPassword.md)| 参数 | 

### Return type

[**HttputilJsonResponse**](httputil.JSONResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UserProfilePut**
> HttputilJsonResponse V1UserProfilePut(ctx, body)
修改个人信息

更新用户的个人信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BindBodyUserProfile**](BindBodyUserProfile.md)| 参数 | 

### Return type

[**HttputilJsonResponse**](httputil.JSONResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersEmailPatch**
> HttputilJsonResponse V1UsersEmailPatch(ctx, body, email)
更新一项用户信息

用于账户激活和密码重置

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BindBodyUserPatch**](BindBodyUserPatch.md)| 参数 | 
  **email** | **string**| 邮箱 | 

### Return type

[**HttputilJsonResponse**](httputil.JSONResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersGet**
> InlineResponse2001 V1UsersGet(ctx, optional)
用户列表

获取用户列表信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***V1UsersApiV1UsersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1UsersApiV1UsersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.String**|  | 
 **pageNo** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersPost**
> InlineResponse2002 V1UsersPost(ctx, body)
用户注册

注册一个用户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BindBodyUser**](BindBodyUser.md)| 参数 | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUsernameGet**
> InlineResponse2003 V1UsersUsernameGet(ctx, username)
用户查询

获取一个用户的公开信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| 用户名 | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

