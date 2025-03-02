package eventbridge

import (
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Archives() *schema.Table {
	return &schema.Table{
		Name:        "aws_eventbridge_archives",
		Description: `https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Archive.html`,
		Resolver:    fetchEventbridgeArchives,
		Multiplex:   client.ServiceAccountRegionMultiplexer("events"),
		Transform:   transformers.TransformWithStruct(&types.Archive{}),
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
				Resolver: resolveArchiveArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
