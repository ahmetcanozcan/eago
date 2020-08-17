var describe = (function () {
  function getAssertFunction(testName) {
    return function (cond) {
      if (!cond) {
        console.error("TEST FAILED ", testName);
      } else {
        console.log("TEST PASSED ", testName);
      }
    };
  }

  function getTestFunction(desc) {
    return function testFunction(description, cb) {
      var assert = getAssertFunction(desc + "/" + description);
      cb(assert);
    };
  }

  return function (description, cb) {
    cb(getTestFunction(description));
  };
})();
