package support

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:        "azure_support_services",
		Resolver:    fetchServices,
		Description: "https://learn.microsoft.com/en-us/rest/api/support/services/list?tabs=HTTP#service",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_support_services", client.Namespacemicrosoft_support),
		Transform:   transformers.TransformWithStruct(&armsupport.Service{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionIDPK},
	}
}

func fetchServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsupport.NewServicesClient(cl.Creds, cl.Options)
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
