request.body.json = (function () {
  textFunction = request.body.text;
  return function () {
    return JSON.parse(textFunction());
  };
})();
