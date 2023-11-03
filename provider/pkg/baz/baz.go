package baz

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
