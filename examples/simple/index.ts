import * as foo from "@pulumi/foo";

const random = new foo.Random("random", {
    length: 10,
});


export const output = random.result;