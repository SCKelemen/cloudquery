# Table: datadog_roles

The composite primary key for this table is (**account_name**, **id**).

## Relations

The following tables depend on datadog_roles:
  - [datadog_role_permissions](datadog_role_permissions.md)
  - [datadog_role_users](datadog_role_users.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_name (PK)|String|
|id (PK)|String|
|attributes|JSON|
|relationships|JSON|
|type|String|
|additional_properties|JSON|