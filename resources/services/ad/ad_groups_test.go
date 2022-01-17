package ad_test

import (
	"encoding/json"
	"github.com/cloudquery/cq-provider-msgraph/resources/provider"
	"github.com/cloudquery/cq-provider-msgraph/resources/services/ad"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/cloudquery/faker/v3"

	"github.com/cloudquery/cq-provider-msgraph/client"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
)

func createADGroupsTestServer(t *testing.T) (*msgraph.GraphServiceRequestBuilder, error) {
	mux := httprouter.New()
	mux.GET("/v1.0/groups", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		groups := []msgraph.Group{
			*fakeGroup(t),
		}

		value, err := json.Marshal(groups)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}

		resp := msgraph.Paging{
			NextLink: "",
			Value:    value,
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

	ts := httptest.NewTLSServer(mux)
	u, _ := url.Parse(ts.URL)
	client := client.CreateTestClient(u.Host)
	svc := msgraph.NewClient(&client)
	return svc, nil
}

func fakeGroup(t *testing.T) *msgraph.Group {
	var group msgraph.Group
	if err := faker.FakeDataSkipFields(&group, []string{
		"Conversations",
		"Threads",
		"CalendarView",
		"Events",
		"Calendar",
		"Drive",
		"Drives",
		"Sites",
		"Onenote",
	}); err != nil {
		t.Fatal(err)
	}

	group.Threads = []msgraph.ConversationThread{fakeConversationThread(t)}
	group.Conversations = []msgraph.Conversation{fakeConversation(t)}
	group.Calendar = fakeCalendar(t)
	group.CalendarView = []msgraph.Event{fakeEvent(t)}
	group.Events = []msgraph.Event{fakeEvent(t)}
	group.Drive = fakeDrive(t)
	group.Drives = []msgraph.Drive{*fakeDrive(t)}
	group.Sites = []msgraph.Site{*fakeSite(t)}
	group.Onenote = fakeOnenote(t)
	return &group
}

func fakeOnenote(t *testing.T) *msgraph.Onenote {
	e := msgraph.Onenote{}
	if err := faker.FakeDataSkipFields(&e, []string{
		"Notebooks",
		"Sections",
		"SectionGroups",
		"Pages",
	}); err != nil {
		t.Fatal(err)
	}

	return &e
}
func fakeSite(t *testing.T) *msgraph.Site {
	e := msgraph.Site{}
	if err := faker.FakeDataSkipFields(&e, []string{
		"BaseItem",
		"Drive",
		"Drives",
		"Items",
		"Item",
		"Lists",
		"Sites",
		"Onenote",
		"Analytics",
	}); err != nil {
		t.Fatal(err)
	}
	if err := faker.FakeDataSkipFields(&e.BaseItem, []string{
		"CreatedByUser",
		"LastModifiedByUser",
	}); err != nil {
		t.Fatal(err)
	}
	return &e
}
func fakeDrive(t *testing.T) *msgraph.Drive {
	e := msgraph.Drive{}
	if err := faker.FakeDataSkipFields(&e, []string{
		"BaseItem",
		"Special",
		"Items",
		"List",
		"Root",
	}); err != nil {
		t.Fatal(err)
	}
	if err := faker.FakeDataSkipFields(&e.BaseItem, []string{
		"CreatedByUser",
		"LastModifiedByUser",
	}); err != nil {
		t.Fatal(err)
	}
	return &e
}

func fakeEvent(t *testing.T) msgraph.Event {
	e := msgraph.Event{}
	if err := faker.FakeData(&e.OutlookItem); err != nil {
		t.Fatal(err)
	}
	e.Calendar = fakeCalendar(t)
	return e
}

func fakeCalendar(t *testing.T) *msgraph.Calendar {
	e := msgraph.Calendar{}
	if err := faker.FakeDataSkipFields(&e, []string{
		"Events",
		"CalendarView",
	}); err != nil {
		t.Fatal(err)
	}
	return &e
}

func fakeConversationThread(t *testing.T) msgraph.ConversationThread {
	e := msgraph.ConversationThread{}
	if err := faker.FakeDataSkipFields(&e, []string{
		"Posts",
	}); err != nil {
		t.Fatal(err)
	}
	return e
}

func fakeConversation(t *testing.T) msgraph.Conversation {
	e := msgraph.Conversation{}
	if err := faker.FakeDataSkipFields(&e, []string{
		"Threads",
	}); err != nil {
		t.Fatal(err)
	}
	e.Threads = []msgraph.ConversationThread{fakeConversationThread(t)}
	return e
}

func TestADGroups(t *testing.T) {
	resource := providertest.ResourceTestData{
		Table: ad.AdGroups(),
		Config: client.Config{
			//Subscriptions: []string{"testProject"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			graph, err := createADGroupsTestServer(t)
			if err != nil {
				return nil, err
			}
			c := client.NewMsgraphClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), client.TestTenantId, graph)
			return c, nil
		},
	}
	providertest.TestResource(t, provider.Provider, resource)
}
