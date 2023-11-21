package bar

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NestedObject struct {
	ReqProp string   `pulumi:"reqProp"`
	OptProp []string `pulumi:"optProp,optional"`
}

type Bar struct{}
type BarArgs struct {
	BucketName string         `pulumi:"bucketName"`
	NestedObjs []NestedObject `pulumi:"nestedObjs"`
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

	for _, nestedObj := range args.NestedObjs {
		fmt.Println("ReqProp:", nestedObj.ReqProp)
		fmt.Println("OptProp:", nestedObj.OptProp)
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
