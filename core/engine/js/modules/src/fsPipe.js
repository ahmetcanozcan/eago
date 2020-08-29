(function () {
  return function pipe(reader, writer) {
    var data = reader.read();
    while (typeof data !== "undefined") {
      writer.write(data);
      data = reader.read();
    }
  };
})();
