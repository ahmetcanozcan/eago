response.send = (function () {
  function send(value) {
    var v = JSON.stringify(value);
    response.write(v);
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
