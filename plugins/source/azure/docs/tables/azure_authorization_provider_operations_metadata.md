# Table: azure_authorization_provider_operations_metadata

https://learn.microsoft.com/en-us/rest/api/authorization/provider-operations-metadata/list?tabs=HTTP#provideroperationsmetadata

The composite primary key for this table is (**subscription_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id (PK)|String|
|display_name|String|
|id (PK)|String|
|name|String|
|operations|JSON|
|resource_types|JSON|
|type|String|