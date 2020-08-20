request.body.json = (function () {
  textFunction = request.body.text;
  return function () {
    return JSON.parse(textFunction());
  };
})();

request.body.readLine = (function () {
  var sent = false;
  return function () {
    if (sent) return undefined;
    return request.body.text();
  };
})();
