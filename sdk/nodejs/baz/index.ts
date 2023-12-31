// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { BazArgs } from "./baz";
export type Baz = import("./baz").Baz;
export const Baz: typeof import("./baz").Baz = null as any;
utilities.lazyLoad(exports, ["Baz"], () => require("./baz"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "foo:baz:Baz":
                return new Baz(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("foo", "baz", _module)
