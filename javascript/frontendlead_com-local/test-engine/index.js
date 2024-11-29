const fs = require("fs");
const path = require("path");
const Jasmine = require("jasmine");

// Get the paths for the user code and test case from the environment variables
const userCodePath = process.env.USER_CODE_PATH;
const testCasePath = process.env.TEST_CASE_PATH;

console.log(process.env)

if (!userCodePath || !testCasePath) {
  console.error(
    JSON.stringify({
      error: "USER_CODE_PATH and TEST_CASE_PATH environment variables are required.",
    })
  );
  process.exit(1);
}

// Read the user code and test case from the specified files
let userCode, testCaseCode;
try {
  userCode = fs.readFileSync(userCodePath, "utf8");
  testCaseCode = fs.readFileSync(testCasePath, "utf8");
} catch (err) {
  console.error(
    JSON.stringify({
      error: `Failed to read user code or test case: ${err.message}`,
    })
  );
  process.exit(1);
}

// Save the user code and test case to temporary files to import them
const userCodeFilePath = path.join(__dirname, "userCode.js");
const testCaseFilePath = path.join(__dirname, "testCase.js");
fs.writeFileSync(userCodeFilePath, userCode);
fs.writeFileSync(testCaseFilePath, testCaseCode);

// Import the user's function and the test case
let userFunction, testCase;
try {
  userFunction = require(userCodeFilePath);
  testCase = require(testCaseFilePath);
} catch (err) {
  console.error(
    JSON.stringify({
      error: `Failed to import user code or test case: ${err.message}`,
    })
  );
  process.exit(1);
}

// Set up Jasmine and run the test case
const jasmine = new Jasmine();

jasmine.describe('User Provided Test', () => {
  jasmine.it('should run the user-defined test case', (done) => {
    // Run the test case
    try {
      testCase(userFunction);
      done();
    } catch (err) {
      console.error(
        JSON.stringify({
          error: `Error running the test case: ${err.message}`,
        })
      );
      process.exit(1);
    }
  });
});

// Execute the Jasmine test suite
jasmine.execute();
