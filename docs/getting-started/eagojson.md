# Eago JSON

an Eago project can be configured by using `eago.json` in root of the project directory

Default `eago.json`

```js
{
  "name": "tt",
  "version": "0.1.0",
  "package": false,
  "port": 3000,
  "notFound": "404.html",
  "dependincies": {},
  "devDependincies": {}
}
```

## name

- type : `string`

name of the project. For package projects, name should be full pacakge name(`github.com/<organization name>/<repo name>`)

```javascript
{
  "name":"hello-world"
}
```

## version

- type `string`
- default : `1.0.0`

version number of the project

```javascript
{
  "version":"1.0.0"
}
```

## author

- type `string`

author of the project. indicate the mail address between `<` and `>`

```javascript
{
  "version":"Dave Bowman <hal900@hal.com>"
}
```

## package

- type `boolean`
- default :`false`

```javascript
{
  "package" : false
}
```

## eagoEnv

- type `string`
- default `development`

represents Eago environment. it can be `production` or `development`.

In production mode:

- store logs in a file as json instead of printing stdout.
- silence debug logs

```javascript
{
  "eagoEnv" : "production"
}
```

## port

- type `number`
- default `3000`

```javascript
{
  "port" : 8080
}
```

## staticPath

- type `string`

the root directory of static file server. `staticPath` is a path that relative to `files` directory. For example, `public` refers `<project dir>/files/public/` directory.

```javascript
{
  "staticPath" : "public"
}
```

## notFound

- type `string`
- default `404.html`

the file path of he not found page. if Eago does not find a notFound file, response a not found text.`noFound` is a path that relative to `files` directory. For example, `public/404.html` refers `<project dir>/files/public/404.html` file.

```javascript
{
  "notFound" : "my404.html"
}
```
