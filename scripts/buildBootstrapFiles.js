const path = require("path");
const fs = require("fs");

const bootstrapPath = path.join(
  __dirname,
  "..",
  "core",
  "engine",
  "js",
  "bootstrap"
);
const bootstrapSrcPath = path.join(bootstrapPath, "src");

const handleErr = (err) => {
  if (err) {
    console.log(err);
    process.exit(1);
  }
};

["eago", "handler"].forEach((typeName) => {
  const dirPath = path.join(bootstrapSrcPath, typeName);
  fs.readdir(
    dirPath,
    /**
     * @param {Error} err
     * @param {File[]} files
     */
    (err, files) => {
      handleErr(err);

      let code = "";
      // Colllect codes
      files.forEach((file) => {
        code = fs.readFileSync(path.join(dirPath, file));
      });
      // Wrap it
      code = wrapCode(typeName, code);
      // Write it
      const targetFile = path.join(bootstrapPath, typeName + ".go");
      fs.writeFileSync(targetFile, code, {});
    }
  );
});

/**
 *
 * @param {string} name
 * @param {string} code
 */
function wrapCode(name, code) {
  const publicName = name[0].toUpperCase() + name.substr(1);
  return `
 /*
  * This file generated by \`build.js\` script
  */
 
package bootstrap
 
import (
  "github.com/robertkrimen/otto/ast"
  "github.com/robertkrimen/otto/parser"
)
 
// ${publicName}BootstrapProgram :
var ${publicName}BootstrapProgram *ast.Program
 
func init() {
  ${publicName}BootstrapProgram, _ = parser.ParseFile(nil, "", ${name}BootstrapJS, 0)
}
 
var ${name}BootstrapJS string = \`
  ${code}
\`
  
`;
}
