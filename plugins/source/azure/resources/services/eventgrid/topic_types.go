package eventgrid

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func TopicTypes() *schema.Table {
	return &schema.Table{
		Name:        "azure_eventgrid_topic_types",
		Resolver:    fetchTopicTypes,
		Description: "https://learn.microsoft.com/en-us/rest/api/eventgrid/controlplane-version2022-06-15/topic-types/list?tabs=HTTP#topictypeinfo",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_eventgrid_topic_types", client.Namespacemicrosoft_eventgrid),
		Transform:   transformers.TransformWithStruct(&armeventgrid.TopicTypeInfo{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionIDPK},
	}
}

func fetchTopicTypes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armeventgrid.NewTopicTypesClient(cl.Creds, cl.Options)
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
