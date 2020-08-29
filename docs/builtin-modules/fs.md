# File System

The `fs` module provides an API for interacting with the files in `files` folder. In `fs` operations, can only acces `files` folder in project directory. Almost every `fs` operation use `filepaths` and

To use this module:

```javascript
import fs from "fs";
```

## Class: fs.Info

- isDir `<boolean>` is a directory or not. equivalent of : `this.size === 0`
- size `<number>` size of the file. if it's a directory, size is `0`

## Class: fs.Reader

the `fs.Reader` class allows to read operation on given file.

> :warning: WARNING: all instance of `fs.Reader` has to be closed. for more secure usage see [fs.readFile()](#fsreadFilepathcb)

- **close()**
  closes reader

- **read()**

  - data `<string>`
  - returns `<string>` | `<undefined>`

  reads data from source. If there is no data to read, returns `undefined`

- **setBufferSize(size)**

  - size `<number>` size of the buffer in byte

  sets the buffer size. default buffer size is `1024` bytes

```javascript
import fs from "fs";
const file = new fs.Reader("info.txt");
file.setBufferSize(512);
```

- **constructor(path)**

  - path `<string>`

```javascript
import fs from "fs";

const file = new fs.Reader("info.txt");

let line = file.readLine();
while (typeof line !== undefined) {
  print(line);
  line = file.readLine();
}

// Ensure to close the Reader
file.close();
```

## Class: fs.Writer

the `fs.Writer` class allows to write operation on given file

- **write(data)**

  - data `<string>`

  writes `data` to writer target

- **close()**
  closes writer

- **constructor(path)**

  - path `<string>`

```javascript
import fs from "fs";

const file = new fs.Writer("info_cp.txt");

fs.readFile("info.txt", (line) => {
  file.write(line + "\n");
});
```

## fs.mkdir(path)

- path `<string>`

the `fs.mkdir()` creates directory on given path with mode `0o777`. if directory can not be created throws an error

```javascript
import fs from "fs";

try {
  fs.mkdir("uploads");
} catch (err) {
  console.error("Can not create the folder", err);
}
```

## fs.pipe(reader,writer)

- reader `<fs.Reader>`
- writer `<fs.Writer>` | `<response>`

the `fs.pipe()` pipes a reader to a writer

```javascript
import fs from "fs";

// Pipe 404.html to HTTP client
fs.pipe(new fs.Reader("404.html"), response);

// Copy a file
fs.pipe(new fs.Reader("info.txt"), new fs.Writer("info_cp.txt"));
```

## fs.readDir(path)

- path `<string>`
- returns `<string[]>`

the `fs.readDir()` reads directory return its content

```javascript
import fs from "fs";

const files = fs.readDir("assets");
files.forEach((file) => {
  print(file);
});
```

## fs.readFile(path,cb)

- path `<string>`
- cb `<Function>`
  - line `<string>`

the `fs.readFile()` reads file line by line and call `cb` function for each line

```javascript
import fs from "fs";

response.set("Content-Type", "text/plain");
fs.readFile("info.txt", response.writeln); // Send info.txt file to HTTP client

fs.readFile("info.txt", print); // print whole file to stdout

let lineCount = 0;

fs.readFile("info.txt", (line) => {
  lineCount++;
});

print(lineCount);
```

## fs.stat(path)

- path `<string>`
- returns `<fs.Info>`

the `fs.stat()` returns information about file

```javascript
let files = [];
fs.readDir("images").forEach((file) => {
  const stat = fs.stat(file);
  if (!file.isDir) {
    files.push(file);
  }
});
```

## fs.writeFile(path,data)

- path `<string>`
- data `<string>`

the `fs.writeFile()` writes `data` to file in given `path`

```javascript
import fs from "fs";

fs.writeFile("info.txt", "this conversation can serve no purpose anymore");
```
