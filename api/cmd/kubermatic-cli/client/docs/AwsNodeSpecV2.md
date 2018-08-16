# AwsNodeSpecV2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ami** | **string** | ami to use. Will be defaulted to a ami for your selected operating system and region. Only set this when you know what you do. | [optional] [default to null]
**DiskSize** | **int64** | size of the volume in gb. Only one volume will be created | [default to null]
**InstanceType** | **string** |  | [default to null]
**Tags** | **map[string]string** | additional instance tags | [optional] [default to null]
**VolumeType** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


