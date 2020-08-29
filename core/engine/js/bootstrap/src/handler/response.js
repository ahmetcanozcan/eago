response.send = (function () {
  function send(value) {
    if (typeof value === "string") {
      response.setHeader("Content-Type", "text/html");
      response.write(value);
    } else {
      response.setHeader("Content-Type", "application/json");
      var v = JSON.stringify(value);
      response.write(v);
    }
  }
  return send;
})();

var __reponse__status = response.status;

response.status = function () {
  function status(status) {
    __reponse__status(status);
    return response;
  }
  return status;
};

response.writeln = function (msg) {
  response.write(msg + "\n");
};

response.redirect = (function () {
  var redirectFunc = response.__redirect;
  return function (url, code) {
    var code = code || 200;
    redirectFunc(url, code);
  };
})();
