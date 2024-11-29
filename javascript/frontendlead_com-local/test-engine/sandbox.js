const { VM } = require("vm2");

function executeUserCode(userCode) {
  const vm = new VM({
    timeout: 1000, // 1 second timeout
    sandbox: {
      module: { exports: {} }, // Provide a mock `module.exports` object
    },
  });

  vm.run(userCode); // Execute the user's code in the sandbox
  return vm.sandbox.module.exports; // Return the exported object
}

module.exports = { executeUserCode };
