package advisor

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func RecommendationMetadata() *schema.Table {
	return &schema.Table{
		Name:        "azure_advisor_recommendation_metadata",
		Resolver:    fetchRecommendationMetadata,
		Description: "https://learn.microsoft.com/en-us/rest/api/advisor/recommendation-metadata/list?tabs=HTTP#metadataentity",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_advisor_recommendation_metadata", client.Namespacemicrosoft_advisor),
		Transform:   transformers.TransformWithStruct(&armadvisor.MetadataEntity{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionIDPK},
	}
}

func fetchRecommendationMetadata(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armadvisor.NewRecommendationMetadataClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
