# ClusterStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiserverCert** | [***KeyCert**](KeyCert.md) |  | [optional] [default to null]
**ApiserverSshKey** | [***RsaKeys**](RSAKeys.md) |  | [optional] [default to null]
**ErrorMessage** | **string** | ErrorMessage contains a defauled error message in case the controller encountered an error. Will be reset if the error was resolved | [optional] [default to null]
**ErrorReason** | [***ClusterStatusError**](ClusterStatusError.md) |  | [optional] [default to null]
**Health** | [***ClusterHealth**](ClusterHealth.md) |  | [optional] [default to null]
**KubeletCert** | [***KeyCert**](KeyCert.md) |  | [optional] [default to null]
**LastUpdated** | [***Time**](Time.md) |  | [optional] [default to null]
**NamespaceName** | **string** | NamespaceName defines the namespace the control plane of this cluster is deployed in | [optional] [default to null]
**Phase** | [***ClusterPhase**](ClusterPhase.md) |  | [optional] [default to null]
**RootCA** | [***KeyCert**](KeyCert.md) |  | [optional] [default to null]
**ServiceAccountKey** | [***Bytes**](Bytes.md) |  | [optional] [default to null]
**UserEmail** | **string** | UserEmail contains the email of the owner of this cluster | [optional] [default to null]
**UserName** | **string** | UserName contains the name of the owner of this cluster | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


