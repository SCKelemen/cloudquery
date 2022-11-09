-- https://hub.steampipe.io/mods/turbot/kubernetes_compliance/queries/statefulset_immutable_container_filesystem
WITH stateful_set_containers AS (SELECT uid, value AS container 
                               FROM k8s_apps_stateful_sets
                               CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers') AS value)

INSERT INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                                resource_name, status)
select uid                              AS resource_id,
       :'execution_time'::timestamp                          AS execution_time,
       :'framework'                                          AS framework,
       :'check_id'                                           AS check_id,
       'Stateful containers root file system is read-only' AS title,
       context                                               AS context,
       namespace                                             AS namespace,
       name                                                  AS resource_name,
       CASE
           WHEN
               (SELECT COUNT(*) FROM stateful_set_containers WHERE stateful_set_containers.uid = k8s_apps_stateful_sets.uid AND
                stateful_set_containers.container->'securityContext' ->> 'readOnlyRootFilesystem' IS DISTINCT FROM 'true') > 0
               THEN 'fail'
           ELSE 'pass'
           END                                               AS status
FROM k8s_apps_stateful_sets;