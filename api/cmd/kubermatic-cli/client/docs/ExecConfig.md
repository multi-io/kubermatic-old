# ExecConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | Preferred input version of the ExecInfo. The returned ExecCredentials MUST use the same encoding version as the input. | [optional] [default to null]
**Args** | **[]string** | Arguments to pass to the command when executing it. +optional | [optional] [default to null]
**Command** | **string** | Command to execute. | [optional] [default to null]
**Env** | [**[]ExecEnvVar**](ExecEnvVar.md) | Env defines additional environment variables to expose to the process. These are unioned with the host&#39;s environment, as well as variables client-go uses to pass argument to the plugin. +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


