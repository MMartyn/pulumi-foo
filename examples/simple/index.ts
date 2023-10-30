import * as foo from "@pulumi/foo";

const random = new foo.Random("my-random", { length: 24 });

export const output = random.result;