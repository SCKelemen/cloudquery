package dns

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"

	"google.golang.org/api/dns/v1"
)

type MockManagedZonesResult struct {
	ManagedZones []*dns.ManagedZone `json:"managedzones,omitempty"`
}

func createManagedZones(mux *httprouter.Router) error {
	var item dns.ManagedZone
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	mux.GET("/*filepath", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &MockManagedZonesResult{
			ManagedZones: []*dns.ManagedZone{&item},
		}
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})
	return nil
}

func TestManagedZones(t *testing.T) {
	client.MockTestRestHelper(t, ManagedZones(), createManagedZones, client.TestOptions{})
}
