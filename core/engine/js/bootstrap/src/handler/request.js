request.body.json = (function () {
  textFunction = request.body.text;
  return function () {
    return JSON.parse(textFunction());
  };
})();

request.body.read = (function () {
  var sent = false;
  return function () {
    if (sent) return undefined;
    sent = true;
    return request.body.text();
  };
})();
