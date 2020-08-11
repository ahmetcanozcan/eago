var require = (function () {
  function require(str) {
    if (str.substr(0, 2) == "./") {
      str = __dirname + str.substr(2);
    }
    return __require(str);
  }
  return require;
})();
