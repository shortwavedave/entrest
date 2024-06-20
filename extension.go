// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package entrest

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ogen-go/ogen"
)

// Ensure that Extension implements the entc.Extension interface.
var _ entc.Extension = (*Extension)(nil)

type Extension struct {
	entc.DefaultExtension

	config *Config
}

func NewExtension(config *Config) (*Extension, error) {
	if config == nil {
		config = &Config{}
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &Extension{config: config}, nil
}

func (e *Extension) Hooks() []gen.Hook {
	return []gen.Hook{
		func(next gen.Generator) gen.Generator {
			return gen.GenerateFunc(func(g *gen.Graph) error {
				if !e.config.DisablePatchJSONTag {
					err := e.patchJSONTag(g)
					if err != nil {
						return err
					}
				}
				return next.Generate(g)
			})
		},
		func(next gen.Generator) gen.Generator {
			return gen.GenerateFunc(func(g *gen.Graph) error {
				spec, err := e.Generate(g)
				if err != nil {
					return err
				}

				err = e.writeSpec(g, spec)
				if err != nil {
					return err
				}
				return next.Generate(g)
			})
		},
	}
}

func (e *Extension) Generate(g *gen.Graph) (*ogen.Spec, error) {
	// Validate all annotations first.
	err := ValidateAnnotations(g.Nodes...)
	if err != nil {
		return nil, fmt.Errorf("failed to validate annotations: %w", err)
	}

	spec := e.config.Spec
	if spec == nil {
		spec = ogen.NewSpec()
	}

	if e.config.PreGenerateHook != nil {
		err = e.config.PreGenerateHook(g, spec)
		if err != nil {
			return nil, err
		}
	}

	// If they weren't provided, set some defaults which are required by OpenAPI,
	// as well as most code-generators.
	if spec.OpenAPI == "" {
		spec.OpenAPI = OpenAPIVersion
	}
	if spec.Info.Title == "" {
		spec.Info.Title = "EntGo Rest API"
	}
	if spec.Info.Version == "" {
		spec.Info.Version = "1.0.0"
	}

	var specs []*ogen.Spec
	var tspec *ogen.Spec

	for _, t := range g.Nodes {
		ta := GetAnnotation(t)

		if ta.GetSkip(e.config) {
			continue
		}

		for _, op := range ta.GetOperations(e.config) {
			tspec, err = GetSpecType(t, op)
			if err != nil {
				panic(err)
			}
			specs = append(specs, tspec)
		}

		for _, edge := range t.Edges {
			ea := GetAnnotation(edge)

			if ea.GetSkip(e.config) || !ea.GetEdgeEndpoint(e.config) {
				continue
			}

			ops := ta.GetOperations(e.config)

			if edge.Unique && slices.Contains(ops, OperationRead) {
				tspec, err = GetSpecEdge(t, edge, OperationRead)
			}
			if !edge.Unique && slices.Contains(ops, OperationList) {
				tspec, err = GetSpecEdge(t, edge, OperationList)
			}

			if err != nil {
				panic(err)
			}
			specs = append(specs, tspec)
		}
	}

	err = MergeSpecOverlap(spec, specs...)
	if err != nil {
		panic(err)
	}

	if len(spec.Paths) == 0 {
		return nil, errors.New("no schemas contain any operations, thus no spec paths can be generated")
	}

	if e.config.PostGenerateHook != nil {
		err = e.config.PostGenerateHook(g, spec)
		if err != nil {
			return nil, err
		}
	}

	addGlobalErrorResponses(spec, e.config.GlobalErrorResponses)
	addGlobalRequestHeaders(spec, e.config.GlobalRequestHeaders)
	addGlobalResponseHeaders(spec, e.config.GlobalResponseHeaders)

	return spec, nil
}

func (e *Extension) writeSpec(g *gen.Graph, spec *ogen.Spec) error {
	if e.config.PreWriteHook != nil {
		if err := e.config.PreWriteHook(spec); err != nil {
			return err
		}
	}

	if e.config.Writer == nil {
		dir := filepath.Join(g.Target, "rest")

		err := os.MkdirAll(dir, 0o750)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}

		f, err := os.OpenFile(filepath.Join(dir, "openapi.json"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o640)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer f.Close()

		e.config.Writer = f
	}

	enc := json.NewEncoder(e.config.Writer)
	enc.SetIndent("", "    ")
	return enc.Encode(spec)
}

func (e *Extension) Annotations() []entc.Annotation {
	return []entc.Annotation{e.config}
}

func (e *Extension) patchJSONTag(g *gen.Graph) error {
	for _, node := range g.Nodes {
		for _, field := range node.Fields {
			if field.StructTag == `json:"-"` {
				continue
			}
			field.StructTag = fmt.Sprintf("json:%q", field.Name)
		}
	}
	return nil
}
