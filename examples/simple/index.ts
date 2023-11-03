import * as foo from "@pulumi/foo";

const first = new foo.bar.Bar("first", {
    bucketName: "first-bucket",
});

const second = new foo.baz.Baz("second", {
    bucket: first.bucket,
});

export const output = second.bucketTwo;