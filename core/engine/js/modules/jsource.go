package modules

// This file created from a script. don't modify!

import "github.com/robertkrimen/otto"

func fsPipeValue(vm *otto.Otto) otto.Value {
  o, err := vm.Object(`
(function () {
  return function pipe(reader, writer) {
    var data = reader.read();
    while (typeof data !== "undefined") {
      writer.write(data);
      data = reader.read();
    }
  };
})();

  `)
  if err != nil {
    panic(err)
  }
  return o.Value()
}
