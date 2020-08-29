# Base Runtime

In Eago, scripts can be executed on several types of runtimes. And each runtime has its own unique built-in properties on its own global scope. Every runtime is inherited from base runtime and has properties of it.

## print(msg)

- msg `<any>`

prints `msg` to stdout without time info and tag

## process

the `process` object is a global object that provides information about and control over the current Eago runtime

### process.env

- `<Object>`

the `process.env` returns an object containing the user environment

An example of `process.env` looks like:

```javascript
{
  TERM: 'xterm-256color',
  PATH: '~/.bin/:/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/bin',
  PWD:"/Users/hal9k",
  HOME:"/Users/hal9k",
  EDITOR: 'vscode',
}
```

example usage:

```javascript
const secret = process.env.JWT_SECRET;
console.log(secret);
```

### process.exit([code])

- code `<integer>` The exit code. **Default:** 0

`process.exit()` terminates the current runtime. if the runtime is handler runtime, only terminate handler runtime but the program will be still running. Otherwise the program will be terminated.

Terminating runtime with error code :

```javascript
process.exit(1);
```

Calling `process.exit()` will force the runtime to exit as quickly as possible even if there are still some operations that have not yet completed fully, including I/O operations to stdout and stderr.

## console {docsify-ignore-all}

- `<Object>`

`console` is a logs info,error and warning to standard output streams or a file.

### console.error([data][,...args]) {docsify-ignore}

- data `<any>`
- args `<any>`

prints to `stderr` with newline,time information, and error tag

```javascript
const code = 5;

console.error("exit with code %d", code);
// prints: [2020-08-19T01:14:00+03:00] ERR exit with code 5

console.error("exit with code", code);
// prints: [2020-08-19T01:14:00+03:00] ERR exit with code 5
```

### console.info([data][,...args]) {docsify-ignore}

- data `<any>`
- args `<any>`

The `console.info()` function is an alias for [console.log()](<#console.log([data][,...args])>).

### console.log([data][,...args]) {docsify-ignore}

- data `<any>`
- args `<any>`

prints to `stdout` with newline,time information, and info tag

```javascript
console.log("log message");
// prints: [2020-08-19T01:14:00+03:00] INFO log message
```

### console.warn([data][,...args]) {docsify-ignore}

- data `<any>`
- args `<any>`

prints to `stderr` with newline,time information, and warning tag

```javascript
console.warn("warning, warning");
// prints: [2020-08-19T01:14:00+03:00] WARN warning, warning
```

## Other ECMAScript 2015 standard built-in object

Eago has almost every es2015 standard built-in objects. For more information check http://www.ecma-international.org/ecma-262/6.0/#sec-ecmascript-standard-built-in-objects
