// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import * as pulumiAws from "@pulumi/aws";

export class Bar extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'foo:bar:Bar';

    /**
     * Returns true if the given object is an instance of Bar.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Bar {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Bar.__pulumiType;
    }

    public /*out*/ readonly bucket!: pulumi.Output<pulumiAws.s3.Bucket>;
    public readonly bucketName!: pulumi.Output<string>;

    /**
     * Create a Bar resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BarArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.bucketName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucketName'");
            }
            resourceInputs["bucketName"] = args ? args.bucketName : undefined;
            resourceInputs["bucket"] = undefined /*out*/;
        } else {
            resourceInputs["bucket"] = undefined /*out*/;
            resourceInputs["bucketName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Bar.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a Bar resource.
 */
export interface BarArgs {
    bucketName: pulumi.Input<string>;
}