<!--TODO: -->

## request

- `<Object>`

The `request` object represents the HTTP request and has properties for the request parameters, body, HTTP headers, and so on.

### request.body

- `<Object>`

the `reqeust.body` represent raw data submitted in the request body. it can be parsed as JSON, text or multipart/form-data.

#### request.body.read()

Reads line from request body and updates reading offset to otherline. If there is no line to read, returns `undefined`

```javascript
// This code snippet writes whole body to stdout
let line = request.body.read();
while (typeof line !== "undefined") {
  print(line);
  line = request.body.readLine();
}
```

it can be used with [`fs.pipe()`](builtin-modules/fs.md#fspipereaderwriter)

```javascript
import fs from "fs";

const file = new fs.Writer("uploads/img.png");

fs.pipe(request.body, file);
```

#### request.body.json()

Parses reqeust body as JSON. if request body is not a json, throws an error

```javascript
// Body: {"message":""Hello World"}
const body = request.body.json();
response.send(body.message); // Sends : "Hello World"
```

#### request.body.text()

Reads request body as a raw text

```javascript
// Body: {"message":""Hello World"}
const body = request.body.text();
response.send(body.message); // Sends undefined
response.send(body); // send whole body data
```

### request.header(name)

- name : `<string>`

the `request.header()` returns header value of given name.

```javascript
// Request with X-Auth-Token : secretToken
const token = request.header("X-Auth-Token");
response.send(token); // sends secretToken
```

### request.method

- `<string>`

the `request.method` is the HTTP method of the request: GET, POST, PUT, and so on.

### request.params

- `<Object>`

The `request.params` is an object that contains properties mapped to named route parameters.
For example, if your file path of your handler is /user/\_name.js, it represent /user/\_name path,and then the `name` property available as `request.param.name`

```javascript
// GET user/ahmet
response.send(`Wellcome ${request.params.name}`); // Sends : Wellcome ahmet
```

### request.query

- `<Object>`

The `request.query` is an object that contains properties mapped to named query parameters.

```javascript
// GET /greet?name=ahmet
response.send(`Wellcome ${request.query.name}`); // Sends : Wellcome ahmet
```

## response

The `response` object represents the HTTP response that an Eago application sends

### response.end()

the `response.end()` closes the HTTP connection without sending data

### response.redirect(path[,status])

the `response.redirect()` redirects to the `path` with `status` code

```javascript
response.redirect("/foo/bar");
response.redirect("/404.html", 404);
```

### response.send([data])

- data `<any>`

the `response.send()` sends HTTP response and sets `Content-Type` header for sending value. after sending data, closes connection

```javascript
// Sets Content-Type "text/html" and send Hello World
response.send("Hello World");

// Sets Content-Type "application/json" and send data
const data = { success: true, message: "Hello World" };
response.send(data);
```

- data `<any>`

the `response.send()` sends HTTP response and sets `Content-Type` header for sending value.

```javascript
// Sets Content-Type "text/html" and send Hello World
response.send("Hello World");

// Sets Content-Type "application/json" and send data
const data = { success: true, message: "Hello World" };
response.send(data);
```

### response.setHeader(field,value)

- field `<string>`
- value `<string>`

the `response.setHeader()` sets header `field` of HTTP response to `value`

```javascript
response.set("Content-Type", "text/html");
```

### response.status(code)

- code `<number>`

the `response.status()` sets status code of repsonse to `code`. and returns `response` object.

```javascript
response.status(200).end();

response.status(404).send("Not Found");
```

### response.write(data)

- data `<string>`

the `response.write()` writes `data` to HTTP response

### response.writeln(data)

- data `<string>`

Equvilient of:

```javascript
(data) => {
  response.write(data + "\n");
};
```
