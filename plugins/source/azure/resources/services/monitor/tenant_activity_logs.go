package monitor

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func TenantActivityLogs() *schema.Table {
	return &schema.Table{
		Name:        "azure_monitor_tenant_activity_logs",
		Resolver:    fetchTenantActivityLogs,
		Description: "https://learn.microsoft.com/en-us/rest/api/monitor/tenant-activity-logs/list?tabs=HTTP#eventdata",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_monitor_tenant_activity_logs", client.Namespacemicrosoft_insights),
		Transform:   transformers.TransformWithStruct(&armmonitor.EventData{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionID},
	}
}

func fetchTenantActivityLogs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armmonitor.NewTenantActivityLogsClient(cl.Creds, cl.Options)
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
