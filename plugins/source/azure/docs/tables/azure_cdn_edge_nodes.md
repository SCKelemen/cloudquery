# Table: azure_cdn_edge_nodes

https://learn.microsoft.com/en-us/rest/api/cdn/edge-nodes/list?tabs=HTTP#edgenode

The composite primary key for this table is (**subscription_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id (PK)|String|
|properties|JSON|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|