# Table: azure_eventgrid_topic_types

https://learn.microsoft.com/en-us/rest/api/eventgrid/controlplane-version2022-06-15/topic-types/list?tabs=HTTP#topictypeinfo

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
|type|String|