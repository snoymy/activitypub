package activitypub

import (
	"errors"
	"github.com/buger/jsonparser"
	"time"
	"unsafe"
)

// CanReceiveActivities Types
const (
	ApplicationType  ActivityVocabularyType = "Application"
	GroupType        ActivityVocabularyType = "Group"
	OrganizationType ActivityVocabularyType = "Organization"
	PersonType       ActivityVocabularyType = "Person"
	ServiceType      ActivityVocabularyType = "Service"
)

var ActorTypes = ActivityVocabularyTypes{
	ApplicationType,
	GroupType,
	OrganizationType,
	PersonType,
	ServiceType,
}

// CanReceiveActivities is generally one of the ActivityStreams Actor Types, but they don't have to be.
// For example, a Profile object might be used as an actor, or a type from an ActivityStreams extension.
// Actors are retrieved like any other Object in ActivityPub.
// Like other ActivityStreams objects, actors have an id, which is a URI.
type CanReceiveActivities Item

// Actor is generally one of the ActivityStreams actor Types, but they don't have to be.
// For example, a Profile object might be used as an actor, or a type from an ActivityStreams extension.
// Actors are retrieved like any other Object in ActivityPub.
// Like other ActivityStreams objects, actors have an id, which is a URI.
type Actor struct {
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
	// A reference to an [ActivityStreams] OrderedCollection comprised of all the messages received by the actor;
	// see 5.2 Inbox.
	Inbox Item `jsonld:"inbox,omitempty"`
	// An [ActivityStreams] OrderedCollection comprised of all the messages produced by the actor;
	// see 5.1 Outbox.
	Outbox Item `jsonld:"outbox,omitempty"`
	// A link to an [ActivityStreams] collection of the actors that this actor is following;
	// see 5.4 Following Collection
	Following Item `jsonld:"following,omitempty"`
	// A link to an [ActivityStreams] collection of the actors that follow this actor;
	// see 5.3 Followers Collection.
	Followers Item `jsonld:"followers,omitempty"`
	// A link to an [ActivityStreams] collection of objects this actor has liked;
	// see 5.5 Liked Collection.
	Liked Item `jsonld:"liked,omitempty"`
	// A short username which may be used to refer to the actor, with no uniqueness guarantees.
	PreferredUsername NaturalLanguageValues `jsonld:"preferredUsername,omitempty,collapsible"`
	// A json object which maps additional (typically server/domain-wide) endpoints which may be useful either
	// for this actor or someone referencing this actor.
	// This mapping may be nested inside the actor document as the value or may be a link
	// to a JSON-LD document with these properties.
	Endpoints *Endpoints `jsonld:"endpoints,omitempty"`
	// A list of supplementary Collections which may be of interest.
	Streams   []ItemCollection `jsonld:"streams,omitempty"`
	PublicKey PublicKey        `jsonld:"publicKey,omitempty"`
}

// GetID returns the ID corresponding to the current Actor
func (a Actor) GetID() ID {
	return a.ID
}

// GetLink returns the IRI corresponding to the current Actor
func (a Actor) GetLink() IRI {
	return IRI(a.ID)
}

// GetType returns the type of the current Actor
func (a Actor) GetType() ActivityVocabularyType {
	return a.Type
}

// IsLink validates if currentActivity Pub Actor is a Link
func (a Actor) IsLink() bool {
	return false
}

// IsObject validates if currentActivity Pub Actor is an Object
func (a Actor) IsObject() bool {
	return true
}

// IsCollection returns false for Actor Objects
func (a Actor) IsCollection() bool {
	return false
}

// PublicKey holds the ActivityPub compatible public key data
type PublicKey struct {
	ID           ID           `jsonld:"id,omitempty"`
	Owner        ObjectOrLink `jsonld:"owner,omitempty"`
	PublicKeyPem string       `jsonld:"publicKeyPem,omitempty"`
}

func (p *PublicKey) UnmarshalJSON(data []byte) error {
	if id, err := jsonparser.GetString(data, "id"); err == nil {
		p.ID = ID(id)
	} else {
		return err
	}
	if o, err := jsonparser.GetString(data, "owner"); err == nil {
		p.Owner = IRI(o)
	} else {
		return err
	}
	if pub, err := jsonparser.GetString(data, "publicKeyPem"); err == nil {
		p.PublicKeyPem = pub
	} else {
		return err
	}
	return nil
}
type (
	// Application describes a software application.
	Application = Actor

	// Group represents a formal or informal collective of Actors.
	Group = Actor

	// Organization represents an organization.
	Organization = Actor

	// Person represents an individual person.
	Person = Actor

	// Service represents a service of any kind.
	Service = Actor
)

// ActorNew initializes an CanReceiveActivities type actor
func ActorNew(id ID, typ ActivityVocabularyType) *Actor {
	if !ActorTypes.Contains(typ) {
		typ = ActorType
	}

	a := Actor{ID: id, Type: typ}
	a.Name = NaturalLanguageValuesNew()
	a.Content = NaturalLanguageValuesNew()
	a.Summary = NaturalLanguageValuesNew()
	in := OrderedCollectionNew(ID("test-inbox"))
	out := OrderedCollectionNew(ID("test-outbox"))
	liked := OrderedCollectionNew(ID("test-liked"))

	a.Inbox = in
	a.Outbox = out
	a.Liked = liked
	a.PreferredUsername = NaturalLanguageValuesNew()

	return &a
}

// ApplicationNew initializes an Application type actor
func ApplicationNew(id ID) *Application {
	a := ActorNew(id, ApplicationType)
	o := Application(*a)
	return &o
}

// GroupNew initializes a Group type actor
func GroupNew(id ID) *Group {
	a := ActorNew(id, GroupType)
	o := Group(*a)
	return &o
}

// OrganizationNew initializes an Organization type actor
func OrganizationNew(id ID) *Organization {
	a := ActorNew(id, OrganizationType)
	o := Organization(*a)
	return &o
}

// PersonNew initializes a Person type actor
func PersonNew(id ID) *Person {
	a := ActorNew(id, PersonType)
	o := Person(*a)
	return &o
}

// ServiceNew initializes a Service type actor
func ServiceNew(id ID) *Service {
	a := ActorNew(id, ServiceType)
	o := Service(*a)
	return &o
}

func (a *Actor) Recipients() ItemCollection {
	rec, _ := ItemCollectionDeduplication(&a.To, &a.Bto, &a.CC, &a.BCC, &a.Audience)
	return rec
}

func (a *Actor) Clean() {
	a.BCC = nil
	a.Bto = nil
}

func (a *Actor) UnmarshalJSON(data []byte) error {
	if ItemTyperFunc == nil {
		ItemTyperFunc = JSONGetItemByType
	}
	a.ID = JSONGetID(data)
	a.Type = JSONGetType(data)
	a.Name = JSONGetNaturalLanguageField(data, "name")
	a.Content = JSONGetNaturalLanguageField(data, "content")
	a.Summary = JSONGetNaturalLanguageField(data, "summary")
	a.Context = JSONGetItem(data, "context")
	a.URL = JSONGetURIItem(data, "url")
	a.MediaType = MimeType(JSONGetString(data, "mediaType"))
	a.Generator = JSONGetItem(data, "generator")
	a.AttributedTo = JSONGetItem(data, "attributedTo")
	a.Attachment = JSONGetItem(data, "attachment")
	a.Location = JSONGetItem(data, "location")
	a.Published = JSONGetTime(data, "published")
	a.StartTime = JSONGetTime(data, "startTime")
	a.EndTime = JSONGetTime(data, "endTime")
	a.Duration = JSONGetDuration(data, "duration")
	a.Icon = JSONGetItem(data, "icon")
	a.Preview = JSONGetItem(data, "preview")
	a.Image = JSONGetItem(data, "image")
	a.Updated = JSONGetTime(data, "updated")
	inReplyTo := JSONGetItems(data, "inReplyTo")
	if len(inReplyTo) > 0 {
		a.InReplyTo = inReplyTo
	}
	to := JSONGetItems(data, "to")
	if len(to) > 0 {
		a.To = to
	}
	audience := JSONGetItems(data, "audience")
	if len(audience) > 0 {
		a.Audience = audience
	}
	bto := JSONGetItems(data, "bto")
	if len(bto) > 0 {
		a.Bto = bto
	}
	cc := JSONGetItems(data, "cc")
	if len(cc) > 0 {
		a.CC = cc
	}
	bcc := JSONGetItems(data, "bcc")
	if len(bcc) > 0 {
		a.BCC = bcc
	}
	replies := JSONGetItem(data, "replies")
	if replies != nil {
		a.Replies = replies
	}
	tag := JSONGetItems(data, "tag")
	if len(tag) > 0 {
		a.Tag = tag
	}
	a.Likes = JSONGetItem(data, "likes")
	a.Shares = JSONGetItem(data, "shares")
	a.Source = GetAPSource(data)
	a.PreferredUsername = JSONGetNaturalLanguageField(data, "preferredUsername")
	a.Followers = JSONGetItem(data, "followers")
	a.Following = JSONGetItem(data, "following")
	a.Inbox = JSONGetItem(data, "inbox")
	a.Outbox = JSONGetItem(data, "outbox")
	a.Liked = JSONGetItem(data, "liked")
	a.Endpoints = JSONGetActorEndpoints(data, "endpoints")
	a.Streams = JSONGetStreams(data, "streams")
	a.PublicKey = JSONGetPublicKey(data, "publicKey")
	return nil
}

// Endpoints a json object which maps additional (typically server/domain-wide)
// endpoints which may be useful either for this actor or someone referencing this actor.
// This mapping may be nested inside the actor document as the value or may be a link to
// a JSON-LD document with these properties.
type Endpoints struct {
	// UploadMedia Upload endpoint URI for this user for binary data.
	UploadMedia Item `jsonld:"uploadMedia,omitempty"`
	// OauthAuthorizationEndpoint Endpoint URI so this actor's clients may access remote ActivityStreams objects which require authentication
	// to access. To use this endpoint, the client posts an x-www-form-urlencoded id parameter with the value being
	// the id of the requested ActivityStreams object.
	OauthAuthorizationEndpoint Item `jsonld:"oauthAuthorizationEndpoint,omitempty"`
	// OauthTokenEndpoint If OAuth 2.0 bearer tokens [RFC6749] [RFC6750] are being used for authenticating client to server interactions,
	// this endpoint specifies a URI at which a browser-authenticated user may obtain a new authorization grant.
	OauthTokenEndpoint Item `jsonld:"oauthTokenEndpoint,omitempty"`
	// ProvideClientKey  If OAuth 2.0 bearer tokens [RFC6749] [RFC6750] are being used for authenticating client to server interactions,
	// this endpoint specifies a URI at which a client may acquire an access token.
	ProvideClientKey Item `jsonld:"provideClientKey,omitempty"`
	// SignClientKey If Linked Data Signatures and HTTP Signatures are being used for authentication and authorization,
	// this endpoint specifies a URI at which browser-authenticated users may authorize a client's public
	// key for client to server interactions.
	SignClientKey Item `jsonld:"signClientKey,omitempty"`
	// SharedInbox An optional endpoint used for wide delivery of publicly addressed activities and activities sent to followers.
	// SharedInbox endpoints SHOULD also be publicly readable OrderedCollection objects containing objects addressed to the
	// Public special collection. Reading from the sharedInbox endpoint MUST NOT present objects which are not addressed to the Public endpoint.
	SharedInbox Item `jsonld:"sharedInbox,omitempty"`
}

// UnmarshalJSON
func (e *Endpoints) UnmarshalJSON(data []byte) error {
	e.OauthAuthorizationEndpoint = JSONGetItem(data, "oauthAuthorizationEndpoint")
	e.OauthTokenEndpoint = JSONGetItem(data, "oauthTokenEndpoint")
	e.UploadMedia = JSONGetItem(data, "uploadMedia")
	e.ProvideClientKey = JSONGetItem(data, "provideClientKey")
	e.SignClientKey = JSONGetItem(data, "signClientKey")
	e.SharedInbox = JSONGetItem(data, "sharedInbox")
	return nil
}

// ToActor
func ToActor(it Item) (*Actor, error) {
	switch i := it.(type) {
	case *Actor:
		return i, nil
	case Actor:
		return &i, nil
	case *Object:
		return (*Actor)(unsafe.Pointer(i)), nil
	case Object:
		return (*Actor)(unsafe.Pointer(&i)), nil
	}
	return nil, errors.New("unable to convert object")
}