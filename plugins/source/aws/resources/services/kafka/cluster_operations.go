// Code generated by codegen; DO NOT EDIT.

package kafka

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ClusterOperations() *schema.Table {
	return &schema.Table{
		Name:        "aws_kafka_cluster_operations",
		Description: `https://docs.aws.amazon.com/msk/1.0/apireference/clusters-clusterarn-operations.html`,
		Resolver:    fetchKafkaClusterOperations,
		Multiplex:   client.ServiceAccountRegionMultiplexer("kafka"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OperationArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "cluster_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveKafkaTags("OperationArn"),
			},
			{
				Name:     "client_request_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClientRequestId"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "end_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("EndTime"),
			},
			{
				Name:     "error_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ErrorInfo"),
			},
			{
				Name:     "operation_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OperationState"),
			},
			{
				Name:     "operation_steps",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OperationSteps"),
			},
			{
				Name:     "operation_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OperationType"),
			},
			{
				Name:     "source_cluster_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SourceClusterInfo"),
			},
			{
				Name:     "target_cluster_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TargetClusterInfo"),
			},
		},
	}
}
