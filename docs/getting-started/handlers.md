# Handlers

Handler scripts are processed by Eago using an interpreter, which executes the code. Handlers run on its own runtime that provides request and response global objects that provide information and control over the incoming HTTP request. Each handler assigned to a URL path. the file location of a handler represents the URL path of the handler.

## URL Path

If there is no special condition, it's so simple to generate URL paths from file location, just delete `.js` extention and that's it.

- `api/all.js` represents `/api/all`

For creating more advance and usefull URL paths, Eago offers several features

### index.js

`index.js` files refer to parent directory just like `index.html`. For example `api/index.js` refers `api/`

> :warning: Be careful using `index.js`. For example: `api/index.js` and `api.js` refers to the same URL path `api` and it causes ambiguity

### Underscore

The file or directory names that start with `_`, represents the variable part of the URL. For example, `api/_id.js` refers to `api/1` , `api/2`, `api/HAL900`, `api/everything` and so on. The variable part can be readable in Handler script using `request.params`.

```javascript
// handlers/api/_from/_to/_year.js

// GET /api/hal/dave/2001

print(`${from} sent message to ${to} in ${year}`);
// Output : hal sent message to dave in 2001
```

### Method Name

By default, a handler script is executed for all HTTP methods but it can be spesified using `filename.<method name>.js`. For example : `api/index.get.js` represents `GET /api`.

If a handler and method name specified version of it exists at the same time and method of the incoming request is the specified method name, Eago runs method name specified handler otherwise runs default handler.

In this example, Eago runs `index.get.js` for `GET /`. For other request methods, Eago runs`index.js`

```
handlers
  ├── index.get.js
  └── index.js
```

### Example

handlers structure of [example todo app](https://github.com/ahmetcanozcan/eago/tree/master/examples/app/todo_app)

```
.
└── handlers
      └── api
           ├── index.get.js
           ├── index.post.js
           ├── _id.get.js
           ├── _id.put.js
           └── _id.delete.js
```
