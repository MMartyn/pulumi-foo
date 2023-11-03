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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/pulumi/pulumi-go-provider/middleware/schema"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

const Name string = "foo"

func Provider() p.Provider {
	// We tell the provider what resources it needs to support.
	// In this case, a single custom resource.
	return infer.Provider(infer.Options{
		Components: []infer.InferredComponent{
			infer.Component[*Bar, BarArgs, *BarState](),
			infer.Component[*Baz, BazArgs, *BazState](),
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

type Bar struct{}
type BarArgs struct {
	BucketName string `pulumi:"bucketName"`
}

type BarState struct {
	pulumi.ResourceState
	BarArgs
	// Outputs
	Bucket s3.BucketOutput `pulumi:"bucket" provider:"type=aws@6.1.0:s3%2Fbucket:Bucket"`
}

func (i *Bar) Construct(ctx *pulumi.Context, name, typ string, args BarArgs, opts pulumi.ResourceOption) (*BarState, error) {
	comp := &BarState{BarArgs: args}
	err := ctx.RegisterComponentResource(typ, name, comp, opts)
	if err != nil {
		return nil, err
	}

	// ---- Bucket
	bucket, err := s3.NewBucket(ctx, name, &s3.BucketArgs{
		Bucket: pulumi.String(args.BucketName),
	}, pulumi.Parent(comp))
	if err != nil {
		return nil, err
	}

	comp.Bucket = bucket.ToBucketOutput()

	return comp, nil
}

type Baz struct{}
type BazArgs struct {
	Bucket s3.BucketInput `pulumi:"bucket" provider:"type=aws@6.1.0:s3%2Fbucket:Bucket"`
}

type BazState struct {
	pulumi.ResourceState
	BazArgs
	// Outputs
	BucketTwo s3.BucketOutput `pulumi:"bucketTwo" provider:"type=aws@6.1.0:s3%2Fbucket:Bucket"`
}

func (i *Baz) Construct(ctx *pulumi.Context, name, typ string, args BazArgs, opts pulumi.ResourceOption) (*BazState, error) {
	comp := &BazState{BazArgs: args}
	err := ctx.RegisterComponentResource(typ, name, comp, opts)
	if err != nil {
		return nil, err
	}

	// ---- Bucket Two
	bucket, err := s3.NewBucket(ctx, name, &s3.BucketArgs{
		Bucket: pulumi.Sprintf("%s-%s", args.Bucket.ToBucketOutput().Bucket(), "two"),
	}, pulumi.Parent(comp))
	if err != nil {
		return nil, err
	}

	comp.BucketTwo = bucket.ToBucketOutput()

	return comp, nil
}
