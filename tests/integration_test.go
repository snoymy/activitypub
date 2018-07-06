package tests

import (
	"testing"

	"github.com/mariusor/activitypub.go/activitypub"
	"github.com/mariusor/activitypub.go/jsonld"

	"strings"
)

func TestAcceptSerialization(t *testing.T) {
	obj := activitypub.AcceptNew("https://localhost/myactivity", nil)
	obj.Name = make(activitypub.NaturalLanguageValue, 1)
	obj.Name["en"] = "test"
	obj.Name["fr"] = "teste"

	jsonld.Ctx = &jsonld.Context{URL: "https://www.w3.org/ns/activitystreams"}

	data, err := jsonld.Marshal(obj)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if !strings.Contains(string(data), string(jsonld.Ctx.URL)) {
		t.Errorf("Could not find context url %#v in output %s", jsonld.Ctx.URL, data)
	}
	if !strings.Contains(string(data), string(obj.ID)) {
		t.Errorf("Could not find id %#v in output %s", string(obj.ID), data)
	}
	if !strings.Contains(string(data), string(obj.Name["en"])) {
		t.Errorf("Could not find name %#v in output %s", string(obj.Name["en"]), data)
	}
	if !strings.Contains(string(data), string(obj.Name["fr"])) {
		t.Errorf("Could not find name %#v in output %s", string(obj.Name["fr"]), data)
	}
	if !strings.Contains(string(data), string(obj.Type)) {
		t.Errorf("Could not find activity type %#v in output %s", obj.Type, data)
	}
}

func TestCreateActivityHTTPSerialization(t *testing.T) {
	id := activitypub.ObjectID("test_object")
	obj := activitypub.AcceptNew(id, nil)
	obj.Name["en"] = "Accept New"

	baseURI := string(activitypub.ActivityBaseURI)
	jsonld.Ctx = &jsonld.Context{
		URL: jsonld.Ref(baseURI + string(obj.Type)),
	}
	data, err := jsonld.Marshal(obj)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(string(data), string(jsonld.Ctx.URL)) {
		t.Errorf("Could not find context url %#v in output %s", jsonld.Ctx.URL, data)
	}
	if !strings.Contains(string(data), string(obj.ID)) {
		t.Errorf("Could not find id %#v in output %s", string(obj.ID), data)
	}
	if !strings.Contains(string(data), string(obj.Name["en"])) {
		t.Errorf("Could not find name %#v in output %s", string(obj.Name["en"]), data)
	}
	if !strings.Contains(string(data), string(obj.Type)) {
		t.Errorf("Could not find activity type %#v in output %s", obj.Type, data)
	}
}
