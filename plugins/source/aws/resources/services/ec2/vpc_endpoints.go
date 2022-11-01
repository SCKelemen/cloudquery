// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func VpcEndpoints() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_vpc_endpoints",
		Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VpcEndpoint.html",
		Resolver:    fetchEc2VpcEndpoints,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveVpcEndpointArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "creation_timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTimestamp"),
			},
			{
				Name:     "dns_entries",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DnsEntries"),
			},
			{
				Name:     "dns_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DnsOptions"),
			},
			{
				Name:     "groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Groups"),
			},
			{
				Name:     "ip_address_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IpAddressType"),
			},
			{
				Name:     "last_error",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LastError"),
			},
			{
				Name:     "network_interface_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("NetworkInterfaceIds"),
			},
			{
				Name:     "owner_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OwnerId"),
			},
			{
				Name:     "policy_document",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyDocument"),
			},
			{
				Name:     "private_dns_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("PrivateDnsEnabled"),
			},
			{
				Name:     "requester_managed",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("RequesterManaged"),
			},
			{
				Name:     "route_table_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("RouteTableIds"),
			},
			{
				Name:     "service_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceName"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "subnet_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SubnetIds"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "vpc_endpoint_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcEndpointId"),
			},
			{
				Name:     "vpc_endpoint_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcEndpointType"),
			},
			{
				Name:     "vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcId"),
			},
		},
	}
}
