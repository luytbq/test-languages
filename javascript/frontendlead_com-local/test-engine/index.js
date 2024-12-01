import * as fs from "fs";
import * as vm from "vm";


// Get the paths for the user code and test case from the environment variables
const userCodePath = process.env.USER_CODE_PATH;
const testCasePath = process.env.TEST_CASE_PATH;

if (!userCodePath || !testCasePath) {
  console.error(
    JSON.stringify({
      error: "USER_CODE_PATH and TEST_CASE_PATH environment variables are required.",
    })
  );
  process.exit(1);
}

const importJasmine = import('./node_modules/jasmine/lib/jasmine.js', {with: {type: 'module'}});
// const importJasmineCore = import('./node_modules/jasmine-core/lib/jasmine-core.js', {with: {type: 'module'}});
const importUserCode = import(userCodePath, {with: {type: 'module'}});

importJasmine.then((module) => {
  const jasmine = new module.default();

  // console.log(jasmine.env.it);

  globalThis.jasmine = jasmine;
  globalThis.describe = jasmine.env.describe;
  globalThis.it = jasmine.env.it;
  globalThis.expect = jasmine.env.expect;

  const result = vm.runInThisContext(  "it('test it', (done) => { expect('1').to.deep.equal('2'); done(); })");
  console.log(result.getFullName);
  // vm.runInThisContext(  "console.log('abcbac')");
})

// importJasmineCore.then((jasmineCore) => {
//   importUserCode.then((userCodeModule) => {
//     // Make Jasmine's functions globally available
//     // const jasmine = jasmineCore.default || jasmineCore; // Adjust for how the library exports

//     // console.log(jasmineCore.default);
//     console.log(jasmineCore.default);

//     // globalThis.jasmine = jasmineCore.default;
//     // globalThis.describe = jasmine.describe;
//     // globalThis.it = jasmine.it;
//     // globalThis.expect = jasmine.expect;

//     // const functionName = userCodeModule.default.name;
//     // globalThis[functionName] = userCodeModule.default

//     // const testScript = fs.readFileSync(testCasePath, "utf-8");

//     // // vm.runInThisContext(testScript, { filename: testScriptName });
//     // vm.runInThisContext(testScript);
//   })
// });