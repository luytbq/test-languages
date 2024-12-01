import * as fs from "fs";
import * as vm from "vm";

const filename = "./etc/import.js"
const testScriptName = "./test/test.js"

import(filename, {with: {type: "module"}}).then((module) => {

    // console.log(module.default(4));
    // console.log(module.default.name)
    const functionName = module.default.name;
    globalThis[functionName] = module.default

    const testScript = fs.readFileSync(testScriptName, "utf-8");

    // vm.runInThisContext(testScript, { filename: testScriptName });
    vm.runInThisContext(testScript);
}, (err) => {
    console.error(err);
})