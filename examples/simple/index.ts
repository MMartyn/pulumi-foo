import * as foo from "@pulumi/foo";
import "@pulumi/aws";

const first = new foo.bar.Bar("first", {
    bucketName: "first-bucket",
});

const second = new foo.baz.Baz("second", {
    bucket: first.bucket,
}, { dependsOn: first});

export const output = second.bucketTwo;