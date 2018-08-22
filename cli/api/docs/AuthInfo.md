# AuthInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**As** | **string** | Impersonate is the username to imperonate.  The name matches the flag. +optional | [optional] [default to null]
**AsGroups** | **[]string** | ImpersonateGroups is the groups to imperonate. +optional | [optional] [default to null]
**AsUserExtra** | [**map[string][]string**](array.md) | ImpersonateUserExtra contains additional information for impersonated user. +optional | [optional] [default to null]
**AuthProvider** | [***AuthProviderConfig**](AuthProviderConfig.md) |  | [optional] [default to null]
**ClientCertificate** | **string** | ClientCertificate is the path to a client cert file for TLS. +optional | [optional] [default to null]
**ClientCertificateData** | **[]int32** | ClientCertificateData contains PEM-encoded data from a client cert file for TLS. Overrides ClientCertificate +optional | [optional] [default to null]
**ClientKey** | **string** | ClientKey is the path to a client key file for TLS. +optional | [optional] [default to null]
**ClientKeyData** | **[]int32** | ClientKeyData contains PEM-encoded data from a client key file for TLS. Overrides ClientKey +optional | [optional] [default to null]
**Exec** | [***ExecConfig**](ExecConfig.md) |  | [optional] [default to null]
**Extensions** | [**[]NamedExtension**](NamedExtension.md) | Extensions holds additional information. This is useful for extenders so that reads and writes don&#39;t clobber unknown fields +optional | [optional] [default to null]
**Password** | **string** | Password is the password for basic authentication to the kubernetes cluster. +optional | [optional] [default to null]
**Token** | **string** | Token is the bearer token for authentication to the kubernetes cluster. +optional | [optional] [default to null]
**TokenFile** | **string** | TokenFile is a pointer to a file that contains a bearer token (as described above).  If both Token and TokenFile are present, Token takes precedence. +optional | [optional] [default to null]
**Username** | **string** | Username is the username for basic authentication to the kubernetes cluster. +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


