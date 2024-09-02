// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package entrest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"slices"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ogen-go/ogen"
)

var _ entc.Annotation = (*Config)(nil)

// Config holds the main configuration for this extension.
type Config struct {
	isValidated bool

	// Spec is an optional default spec to merge all generated endpoints/schemas/etc
	// into, which will allow you to specify API info, servers, security schemes, etc.
	Spec *ogen.Spec

	// SpecFromPath is similar to [Config.Spec], but instead of providing a spec directly,
	// it will read the spec (json) from the provided path. If you have a combination
	// of auto-generated endpoints from this extension, plus a bunch of your own endpoints,
	// this can make it very easy to layer each of the specs on top of each other, as it
	// can be a bit tedious to use [Config.Spec] directly.
	SpecFromPath string

	// DisablePagination disables pagination support for all schemas by default.
	// It scan still be enabled on a per-schema basis with annotations.
	DisablePagination bool

	// MinItemsPerPage controls the default minimum number of items per page, for
	// paginated calls. This can be overridden on a per-schema basis with annotations.
	MinItemsPerPage int

	// MaxItemsPerPage controls the default maximum number of items per page, for
	// paginated calls. This can be overridden on a per-schema basis with annotations.
	MaxItemsPerPage int

	// ItemsPerPage controls the default number of items per page, for paginated calls.
	// This can be overridden on a per-schema basis with annotations.
	ItemsPerPage int

	// DefaultEagerLoad enables eager loading of all edges by default. This can be
	// overridden on a per-edge basis with annotations. If edges load a lot of data
	// or are expensive, this can be a performance hit and isn't recommended.
	DefaultEagerLoad bool

	// DisableEagerLoadNonPagedOpt disables the optimization which automatically disables
	// the pagination for edge endpoints where the edge was also eager-loaded. The idea for
	// the optimization is that if the edge is also eager-loaded, then the amount of data
	// isn't large enough to justify the additional overhead of pagination, so we can
	// disable it.
	DisableEagerLoadNonPagedOpt bool

	// DisableEagerLoadedEndpoints disables the generation of dedicated endpoints for
	// edges which are also eager-loaded. This can be useful to reduce the number of
	// endpoints generated, but does mean that callers would have to always call the
	// entity which eager loads the edge, rather than only fetching the edge itself.
	// This can be overridden on a per-edge basis with annotations.
	//
	// Example: Given a schema with users and pets, and an edge on pets called "owner",
	// pointing to user, if you configure owner to be eager-loaded (so any time you query
	// a pet, you also get the owner), setting this to true will then disable the
	// /pets/{id}/owner endpoint (idea being that you could just call /pets/{id} and
	// get the owner from that response).
	DisableEagerLoadedEndpoints bool

	// EagerLoadLimit controls the default maximum number of results that can be
	// eager-loaded for a given edge. The default, when not specified, is 1000. The limit
	// can be disabled by setting the value to -1. The intent of this option is to
	// provide a safe default cap on the number of eager-loaded results, to prevent
	// potential abuse, denial-of-service/resource-exhaustion, etc.
	//
	// This can be overridden on a per-edge basis with annotations.
	EagerLoadLimit int

	// AddEdgesToTags enables the addition of edge fields to the "tags" field in the
	// OpenAPI spec. This is helpful to see if querying a specific entity also returns
	// the thing you're looking for, though can be very noisy for large schemas. Note
	// that edge endpoints (e.g. /users/{id}/pets) will still have both "User" and "Pet"
	// in the tags, this only affects eager-loaded edges.
	AddEdgesToTags bool

	// DefaultFilterID enables the default filter for ID fields, which applies
	// [FilterGroupEqualExact] and [FilterGroupArray] to the ID field. This is helpful
	// if you don't explicitly declare your "id" field in your schema (as it is handled
	// by default by ent).
	DefaultFilterID bool

	// DefaultOperations is a list of operations to generate by default. If nil,
	// all operations will be generated by default (unless excluded with annotations).
	DefaultOperations []Operation

	// GlobalRequestHeaders are headers to add to every request, which can be optional
	// (e.g. X-Request-Id or X-Correlation-ID), or required (e.g. API version). Note
	// that these should not include anything related to authentication -- use the
	// security schemes instead via [Config.Spec].
	GlobalRequestHeaders RequestHeaders

	// GlobalResponseHeaders are headers to add to every response, recommended for headers
	// like X-Ratelimit-Limit, X-Ratelimit-Remaining, X-Ratelimit-Reset, etc.
	GlobalResponseHeaders ResponseHeaders

	// GlobalErrorResponses are status code -> response mappings for errors, which are
	// added to all path operations. Note that some status codes are excluded on specific
	// operations (e.g. 404 on list, 409 on non-create/update, etc). If not specified,
	// a default set of responses will be generated which can be used with entrest's
	// built-in auto-generated HTTP handlers (see below). Defaults to [DefaultErrorResponses].
	GlobalErrorResponses ErrorResponses

	// Handler enables the generation of HTTP handlers for the specified server/routing
	// library. If this is disabled, no Go code will be generated, and only the OpenAPI
	// spec will be generated.
	Handler HTTPHandler

	// StrictMutate if set to true, will cause a 400 "Bad Request" response if an unknown
	// field is provided to the update/create/etc functions. This is useful for ensuring
	// that all fields are provided, and that the client is not attempting to provide
	// fields that are not defined in the schema.
	StrictMutate bool

	// ListNotFound if set to true, will cause a 404 "Not Found" response if a list endpoint
	// (with any filtering as part of the request) returns no results. This is technically
	// "more correct" according to the RFC, but some prefer to return a 200 "OK". In either
	// case, the body of the response would still be the typical pagination or list object,
	// with the "content" field being an empty array.
	ListNotFound bool

	// DisableSpecHandler disables the generation of an OpenAPI spec handler (e.g.
	// /openapi.json). Disabling this will also disable embedding the spec into the
	// binary/rest generated library.
	DisableSpecHandler bool

	// AllowClientIDs, when enabled, allows the built-in "id" field as part of a "Create"
	// payload for entity creation, allowing the client to supply UUIDs as primary keys
	// and for idempotency.
	AllowClientUUIDs bool

	// DisablePatchJSONTag disables a ent generation hook that patches the JSON tag of all
	// fields in the schema, removing the usage of omitempty. This helps ensure that fields
	// that have default values and/or aren't required, still get returned in JSON response
	// bodies. Skips over fields which are json-excluded (e.g. sensitive data).
	DisablePatchJSONTag bool

	// WithTesting enables the generation of a resttest package, which contains a
	// set of helpers for testing the generated REST API.
	WithTesting bool

	// PreHook is a hook that runs before the spec is generated. This is useful for
	// things like adding global security schemes, or adding global request headers,
	// if you're unable to provide the [Config.Spec] field for some reason.
	PreGenerateHook func(g *gen.Graph, spec *ogen.Spec) error `json:"-"`

	// PostHook is a hook that runs after the spec is generated, but before we run global
	// writers (headers, error codes, etc) as well as before we write the spec to disk.
	// Recommended for adding additional paths so they can also receive the global headers,
	// error codes, etc.
	PostGenerateHook func(g *gen.Graph, spec *ogen.Spec) error `json:"-"`

	// PreWriteHook is similar to PostGenerateHook, except it is run directly before
	// writing to disk, after the entire spec has been resolved.
	PreWriteHook func(spec *ogen.Spec) error `json:"-"`

	// Writer is an optional writer to write the spec to. If not provided, the spec
	// will be written to the filesystem under "<ent>/rest/openapi.json".
	Writer io.Writer `json:"-"`
}

func (c *Config) Validate() error {
	if c.isValidated {
		return nil
	}

	if c.Spec != nil && c.SpecFromPath != "" {
		return errors.New("Config.Spec and Config.SpecFromPath cannot be provided at the same time")
	}

	if c.MinItemsPerPage < 1 {
		c.MinItemsPerPage = defaultMinItemsPerPage
	}

	if c.MaxItemsPerPage < 1 {
		c.MaxItemsPerPage = defaultMaxItemsPerPage
	}

	if c.MaxItemsPerPage < c.MinItemsPerPage {
		c.MaxItemsPerPage = c.MinItemsPerPage
	}

	if c.ItemsPerPage < 1 {
		c.ItemsPerPage = defaultItemsPerPage
	}

	if c.ItemsPerPage < c.MinItemsPerPage {
		c.ItemsPerPage = c.MinItemsPerPage
	}

	if c.ItemsPerPage > c.MaxItemsPerPage {
		c.ItemsPerPage = c.MaxItemsPerPage
	}

	if c.EagerLoadLimit < -1 {
		c.EagerLoadLimit = -1
	}

	if c.EagerLoadLimit == 0 {
		c.EagerLoadLimit = 1000
	}

	if c.DefaultOperations == nil {
		c.DefaultOperations = AllOperations
	}

	if len(c.GlobalErrorResponses) == 0 {
		c.GlobalErrorResponses = DefaultErrorResponses
	}

	for k := range c.GlobalErrorResponses {
		if k < 400 {
			return fmt.Errorf("error response defined with status code %d, which is not an HTTP error code", k)
		}
	}

	if c.Handler != HandlerNone && !slices.Contains(AllSupportedHTTPHandlers, c.Handler) {
		return fmt.Errorf("unsupported handler provided: %s", c.Handler)
	}

	if c.Handler == HandlerNone && c.WithTesting {
		c.WithTesting = false
	}

	c.isValidated = true
	return nil
}

func (c Config) Name() string {
	return "RestConfig"
}

func (c *Config) Decode(o any) error {
	buf, err := json.Marshal(o)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, c) //nolint:musttag
}

// GetConfig returns the rest config for the given graph. If the graph does not
// contain the config (extension was not loaded), this will panic.
func GetConfig(gc *gen.Config) *Config {
	c := &Config{}

	if gc == nil || gc.Annotations == nil || gc.Annotations[c.Name()] == nil {
		panic("nil config")
	}

	err := c.Decode(gc.Annotations[c.Name()])
	if err != nil {
		panic(fmt.Sprintf("failed to decode config: %v", err))
	}

	err = c.Validate()
	if err != nil {
		panic(fmt.Sprintf("failed to validate config: %v", err))
	}
	return c
}
