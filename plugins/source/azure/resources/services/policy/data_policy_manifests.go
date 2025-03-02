package policy

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DataPolicyManifests() *schema.Table {
	return &schema.Table{
		Name:        "azure_policy_data_policy_manifests",
		Resolver:    fetchDataPolicyManifests,
		Description: "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy@v0.6.0#DataPolicyManifest",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_policy_data_policy_manifests", client.Namespacemicrosoft_authorization),
		Transform:   transformers.TransformWithStruct(&armpolicy.DataPolicyManifest{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionIDPK},
	}
}

func fetchDataPolicyManifests(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armpolicy.NewDataPolicyManifestsClient(cl.Creds, cl.Options)
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
