package elasticbeanstalk

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elasticbeanstalk/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ConfigurationSettings() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticbeanstalk_configuration_settings",
		Description: `https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ConfigurationSettingsDescription.html`,
		Resolver:    fetchElasticbeanstalkConfigurationSettings,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticbeanstalk"),
		Transform:   transformers.TransformWithStruct(models.ConfigurationSettingsDescriptionWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
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
				Name:     "environment_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}
