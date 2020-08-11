import { readFile } from "fs";

export function writeFile(w, filename) {
  readFile(filename, (line) => {
    w.write(`${line}\n`);
  });
}

export function generateDbURI(host, port, dbname, user, password) {
  return (
    `host= ${host} port=${port} user=${user}` +
    ` password=${password} dbname=${dbname} sslmode=disable`
  );
}
