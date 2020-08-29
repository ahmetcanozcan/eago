const path = require("path");
const fs = require("fs");

const { handleErr } = require("./util");

const modulesPath = path.join(
  __dirname,
  "..",
  "core",
  "engine",
  "js",
  "modules"
);
const modulesSrcPath = path.join(modulesPath, "src");

fs.readdir(
  modulesSrcPath,
  /**
   * @param {Error} err
   * @param {File[]} files
   */
  (err, files) => {
    handleErr(err);
    let content = "";
    // Colllect codes
    files.forEach((file) => {
      let code = fs.readFileSync(path.join(modulesSrcPath, file));
      content += wrapCode(file.split(".")[0], code);
    });
    // Wrap it
    content = wrapContent(content);
    // Write it
    const targetFile = path.join(modulesPath, "jsource.go");
    fs.writeFileSync(targetFile, content, {});
  }
);

function wrapContent(content) {
  return `package modules

// This file created from a script. don't modify!

import "github.com/robertkrimen/otto"
${content}`;
}

/**
 *
 * @param {string} name
 * @param {string} code
 */
function wrapCode(name, code) {
  return `
func ${name}Value(vm *otto.Otto) otto.Value {
  o, err := vm.Object(\`
${code}
  \`)
  if err != nil {
    panic(err)
  }
  return o.Value()
}
`;
}
