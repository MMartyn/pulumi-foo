// Copyright 2016-2023, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"

	"github.com/pulumi/pulumi-go-provider/middleware/schema"

	"github.com/mmartyn/pulumi-foo/provider/pkg/bar"
	"github.com/mmartyn/pulumi-foo/provider/pkg/baz"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

const Name string = "foo"

func Provider() p.Provider {
	// We tell the provider what resources it needs to support.
	// In this case, a single custom resource.
	return infer.Provider(infer.Options{
		Components: []infer.InferredComponent{
			infer.Component[*bar.Bar, bar.BarArgs, *bar.BarState](),
			infer.Component[*baz.Baz, baz.BazArgs, *baz.BazState](),
		},
		Metadata: schema.Metadata{
			LanguageMap: map[string]any{
				"csharp": map[string]any{
					"packageReferences": map[string]string{
						"Pulumi":     "3.*",
						"Pulumi.Aws": "6.*",
					},
				},
				"go": map[string]any{
					"generateResourceContainerTypes": true,
					"importBasePath":                 "github.com/mmartyn/pulumi-foo/sdk/go/foo",
				},
				"nodejs": map[string]any{
					"dependencies": map[string]string{
						"@pulumi/aws": "^6.6.0",
					},
				},
				"python": map[string]any{
					"requires": map[string]string{
						"pulumi-aws": ">=6.6.1,<7.0.0",
					},
				},
			},
		},
	})
}
