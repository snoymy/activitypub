package activitypub

import (
	"bytes"
	"errors"
	"time"
	"unsafe"
)

const CollectionOfItems ActivityVocabularyType = "ItemCollection"

var CollectionTypes = ActivityVocabularyTypes{
	CollectionOfItems,
	CollectionType,
	OrderedCollectionType,
	CollectionPageType,
	OrderedCollectionPageType,
}

type CollectionInterface interface {
	ObjectOrLink
	Collection() ItemCollection
	Append(ob Item) error
	Count() uint
	Contains(IRI) bool
}

// Collection is a subtype of Activity Pub Object that represents ordered or unordered sets of Activity Pub Object or Link instances.
type Collection struct {
	// ID provides the globally unique identifier for anActivity Pub Object or Link.
	ID ID `jsonld:"id,omitempty"`
	// Type identifies the Activity Pub Object or Link type. Multiple values may be specified.
	Type ActivityVocabularyType `jsonld:"type,omitempty"`
	// Name a simple, human-readable, plain-text name for the object.
	// HTML markup MUST NOT be included. The name MAY be expressed using multiple language-tagged values.
	Name NaturalLanguageValues `jsonld:"name,omitempty,collapsible"`
	// Attachment identifies a resource attached or related to an object that potentially requires special handling.
	// The intent is to provide a model that is at least semantically similar to attachments in email.
	Attachment Item `jsonld:"attachment,omitempty"`
	// AttributedTo identifies one or more entities to which this object is attributed. The attributed entities might not be Actors.
	// For instance, an object might be attributed to the completion of another activity.
	AttributedTo Item `jsonld:"attributedTo,omitempty"`
	// Audience identifies one or more entities that represent the total population of entities
	// for which the object can considered to be relevant.
	Audience ItemCollection `jsonld:"audience,omitempty"`
	// Content or textual representation of the Activity Pub Object encoded as a JSON string.
	// By default, the value of content is HTML.
	// The mediaType property can be used in the object to indicate a different content type.
	// (The content MAY be expressed using multiple language-tagged values.)
	Content NaturalLanguageValues `jsonld:"content,omitempty,collapsible"`
	// Context identifies the context within which the object exists or an activity was performed.
	// The notion of "context" used is intentionally vague.
	// The intended function is to serve as a means of grouping objects and activities that share a
	// common originating context or purpose. An example could be all activities relating to a common project or event.
	Context Item `jsonld:"context,omitempty"`
	// MediaType when used on an Object, identifies the MIME media type of the value of the content property.
	// If not specified, the content property is assumed to contain text/html content.
	MediaType MimeType `jsonld:"mediaType,omitempty"`
	// EndTime the date and time describing the actual or expected ending time of the object.
	// When used with an Activity object, for instance, the endTime property specifies the moment
	// the activity concluded or is expected to conclude.
	EndTime time.Time `jsonld:"endTime,omitempty"`
	// Generator identifies the entity (e.g. an application) that generated the object.
	Generator Item `jsonld:"generator,omitempty"`
	// Icon indicates an entity that describes an icon for this object.
	// The image should have an aspect ratio of one (horizontal) to one (vertical)
	// and should be suitable for presentation at a small size.
	Icon Item `jsonld:"icon,omitempty"`
	// Image indicates an entity that describes an image for this object.
	// Unlike the icon property, there are no aspect ratio or display size limitations assumed.
	Image Item `jsonld:"image,omitempty"`
	// InReplyTo indicates one or more entities for which this object is considered a response.
	InReplyTo Item `jsonld:"inReplyTo,omitempty"`
	// Location indicates one or more physical or logical locations associated with the object.
	Location Item `jsonld:"location,omitempty"`
	// Preview identifies an entity that provides a preview of this object.
	Preview Item `jsonld:"preview,omitempty"`
	// Published the date and time at which the object was published
	Published time.Time `jsonld:"published,omitempty"`
	// Replies identifies a Collection containing objects considered to be responses to this object.
	Replies Item `jsonld:"replies,omitempty"`
	// StartTime the date and time describing the actual or expected starting time of the object.
	// When used with an Activity object, for instance, the startTime property specifies
	// the moment the activity began or is scheduled to begin.
	StartTime time.Time `jsonld:"startTime,omitempty"`
	// Summary a natural language summarization of the object encoded as HTML.
	// *Multiple language tagged summaries may be provided.)
	Summary NaturalLanguageValues `jsonld:"summary,omitempty,collapsible"`
	// Tag one or more "tags" that have been associated with an objects. A tag can be any kind of Activity Pub Object.
	// The key difference between attachment and tag is that the former implies association by inclusion,
	// while the latter implies associated by reference.
	Tag ItemCollection `jsonld:"tag,omitempty"`
	// Updated the date and time at which the object was updated
	Updated time.Time `jsonld:"updated,omitempty"`
	// URL identifies one or more links to representations of the object
	URL LinkOrIRI `jsonld:"url,omitempty"`
	// To identifies an entity considered to be part of the public primary audience of an Activity Pub Object
	To ItemCollection `jsonld:"to,omitempty"`
	// Bto identifies anActivity Pub Object that is part of the private primary audience of this Activity Pub Object.
	Bto ItemCollection `jsonld:"bto,omitempty"`
	// CC identifies anActivity Pub Object that is part of the public secondary audience of this Activity Pub Object.
	CC ItemCollection `jsonld:"cc,omitempty"`
	// BCC identifies one or more Objects that are part of the private secondary audience of this Activity Pub Object.
	BCC ItemCollection `jsonld:"bcc,omitempty"`
	// Duration when the object describes a time-bound resource, such as an audio or video, a meeting, etc,
	// the duration property indicates the object's approximate duration.
	// The value must be expressed as an xsd:duration as defined by [ xmlschema11-2],
	// section 3.3.6 (e.g. a period of 5 seconds is represented as "PT5S").
	Duration time.Duration `jsonld:"duration,omitempty"`
	// This is a list of all Like activities with this object as the object property, added as a side effect.
	// The likes collection MUST be either an OrderedCollection or a Collection and MAY be filtered on privileges
	// of an authenticated user or as appropriate when no authentication is given.
	Likes Item `jsonld:"likes,omitempty"`
	// This is a list of all Announce activities with this object as the object property, added as a side effect.
	// The shares collection MUST be either an OrderedCollection or a Collection and MAY be filtered on privileges
	// of an authenticated user or as appropriate when no authentication is given.
	Shares Item `jsonld:"shares,omitempty"`
	// Source property is intended to convey some sort of source from which the content markup was derived,
	// as a form of provenance, or to support future editing by clients.
	// In general, clients do the conversion from source to content, not the other way around.
	Source Source `jsonld:"source,omitempty"`
	// In a paged Collection, indicates the page that contains the most recently updated member items.
	Current ObjectOrLink `jsonld:"current,omitempty"`
	// In a paged Collection, indicates the furthest preceeding page of items in the collection.
	First ObjectOrLink `jsonld:"first,omitempty"`
	// In a paged Collection, indicates the furthest proceeding page of the collection.
	Last ObjectOrLink `jsonld:"last,omitempty"`
	// A non-negative integer specifying the total number of objects contained by the logical view of the collection.
	// This number might not reflect the actual number of items serialized within the Collection object instance.
	TotalItems uint `jsonld:"totalItems"`
	// Identifies the items contained in a collection. The items might be ordered or unordered.
	Items ItemCollection `jsonld:"items,omitempty"`
}

type (
	// FollowersCollection is a collection of followers
	FollowersCollection = Followers

	// Followers is a Collection type
	Followers = Collection

	// FollowingCollection is a list of everybody that the actor has followed, added as a side effect.
	// The following collection MUST be either an OrderedCollection or a Collection and MAY
	// be filtered on privileges of an authenticated user or as appropriate when no authentication is given.
	FollowingCollection = Following

	// Following is a type alias for a simple Collection
	Following = Collection
)

// CollectionNew initializes a new Collection
func CollectionNew(id ID) *Collection {
	c := Collection{ID: id, Type: CollectionType}
	c.Name = NaturalLanguageValuesNew()
	c.Content = NaturalLanguageValuesNew()
	c.Summary = NaturalLanguageValuesNew()
	return &c
}

// OrderedCollectionNew initializes a new OrderedCollection
func OrderedCollectionNew(id ID) *OrderedCollection {
	o := OrderedCollection{ID: id, Type: OrderedCollectionType}
	o.Name = NaturalLanguageValuesNew()
	o.Content = NaturalLanguageValuesNew()

	return &o
}

// GetID returns the ID corresponding to the Collection object
func (c Collection) GetID() ID {
	return c.ID
}

// GetType returns the Collection's type
func (c Collection) GetType() ActivityVocabularyType {
	return c.Type
}

// IsLink returns false for a Collection object
func (c Collection) IsLink() bool {
	return false
}

// IsObject returns true for a Collection object
func (c Collection) IsObject() bool {
	return true
}

// IsCollection returns true for Collection objects
func (c Collection) IsCollection() bool {
	return true
}

// GetLink returns the IRI corresponding to the Collection object
func (c Collection) GetLink() IRI {
	return IRI(c.ID)
}

// Collection returns the Collection's items
func (c Collection) Collection() ItemCollection {
	return c.Items
}

// Append adds an element to a Collection
func (c *Collection) Append(ob Item) error {
	c.Items = append(c.Items, ob)
	return nil
}

// Count returns the maximum between the length of Items in collection and its TotalItems property
func (c *Collection) Count() uint {
	if c.TotalItems > 0 {
		return c.TotalItems
	}
	return uint(len(c.Items))
}

// Contains verifies if Collection array contains the received one
func (c Collection) Contains(r IRI) bool {
	if len(c.Items) == 0 {
		return false
	}
	for _, iri := range c.Items {
		if r.Equals(iri.GetLink(), false) {
			return true
		}
	}
	return false
}

// UnmarshalJSON
func (c *Collection) UnmarshalJSON(data []byte) error {
	if ItemTyperFunc == nil {
		ItemTyperFunc = JSONGetItemByType
	}
	c.ID = JSONGetID(data)
	c.Type = JSONGetType(data)
	c.Name = JSONGetNaturalLanguageField(data, "name")
	c.Content = JSONGetNaturalLanguageField(data, "content")
	c.Summary = JSONGetNaturalLanguageField(data, "summary")
	c.Context = JSONGetItem(data, "context")
	c.URL = JSONGetURIItem(data, "url")
	c.MediaType = MimeType(JSONGetString(data, "mediaType"))
	c.Generator = JSONGetItem(data, "generator")
	c.AttributedTo = JSONGetItem(data, "attributedTo")
	c.Attachment = JSONGetItem(data, "attachment")
	c.Location = JSONGetItem(data, "location")
	c.Published = JSONGetTime(data, "published")
	c.StartTime = JSONGetTime(data, "startTime")
	c.EndTime = JSONGetTime(data, "endTime")
	c.Duration = JSONGetDuration(data, "duration")
	c.Icon = JSONGetItem(data, "icon")
	c.Preview = JSONGetItem(data, "preview")
	c.Image = JSONGetItem(data, "image")
	c.Updated = JSONGetTime(data, "updated")
	inReplyTo := JSONGetItems(data, "inReplyTo")
	if len(inReplyTo) > 0 {
		c.InReplyTo = inReplyTo
	}
	to := JSONGetItems(data, "to")
	if len(to) > 0 {
		c.To = to
	}
	audience := JSONGetItems(data, "audience")
	if len(audience) > 0 {
		c.Audience = audience
	}
	bto := JSONGetItems(data, "bto")
	if len(bto) > 0 {
		c.Bto = bto
	}
	cc := JSONGetItems(data, "cc")
	if len(cc) > 0 {
		c.CC = cc
	}
	bcc := JSONGetItems(data, "bcc")
	if len(bcc) > 0 {
		c.BCC = bcc
	}
	replies := JSONGetItem(data, "replies")
	if replies != nil {
		c.Replies = replies
	}
	tag := JSONGetItems(data, "tag")
	if len(tag) > 0 {
		c.Tag = tag
	}

	c.TotalItems = uint(JSONGetInt(data, "totalItems"))
	c.Items = JSONGetItems(data, "items")

	c.Current = JSONGetItem(data, "current")
	c.First = JSONGetItem(data, "first")
	c.Last = JSONGetItem(data, "last")

	return nil
}

// MarshalJSON
func (c Collection) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	notEmpty := false
	b.Write([]byte{'{'})

	OnObject(c, func(o *Object) error {
		notEmpty = writeObject(&b, *o)
		return nil
	})
	writeComma := func() { b.WriteString(",") }
	writeCommaIfNotEmpty := func(notEmpty bool) {
		if notEmpty {
			writeComma()
		}
	}
	if c.Current != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(&b, "current", c.Current) || notEmpty
	}
	if c.First != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(&b, "first", c.First) || notEmpty
	}
	if c.Last != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemProp(&b, "last", c.Last) || notEmpty
	}
	if c.Items != nil {
		writeCommaIfNotEmpty(notEmpty)
		notEmpty = writeItemCollectionProp(&b, "items", c.Items) || notEmpty
	}
	writeCommaIfNotEmpty(notEmpty)
	notEmpty = writeIntProp(&b, "totalItems", int(c.TotalItems)) || notEmpty

	if notEmpty {
		b.Write([]byte{'}'})
		return b.Bytes(), nil
	}
	return nil, nil
}

// ToCollection
func ToCollection(it Item) (*Collection, error) {
	switch i := it.(type) {
	case *Collection:
		return i, nil
	case Collection:
		return &i, nil
	case *CollectionPage:
		return (*Collection)(unsafe.Pointer(i)), nil
	case CollectionPage:
		return (*Collection)(unsafe.Pointer(&i)), nil
	}
	return nil, errors.New("unable to convert to collection")
}

// FollowingNew initializes a new Following
func FollowingNew() *Following {
	id := ID("following")

	i := Following{ID: id, Type: CollectionType}
	i.Name = NaturalLanguageValuesNew()
	i.Content = NaturalLanguageValuesNew()
	i.Summary = NaturalLanguageValuesNew()

	i.TotalItems = 0

	return &i
}
