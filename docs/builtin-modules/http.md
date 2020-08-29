# HTTP

the `HTTP` module provides http/https communication

## Class: http.RequestBody

extends [`<fs.Reader>`](builtin-modules/fs#class-fsReader)

- text() `<Function>` reads wholoe body as `<string>`
- json() `<Function>` reads wholoe body as `<Object>`

## http.request(url[,options])

- url `<string>`
- options `<Object>`
  - body `<string>` the ody of the request
  - headers `<Object>` the object that contains request headers as key value map
  - method `<string>` name of the HTTP request method. **Default:** `'GET'`
- returns <Object>
  - body `<http.RequestBody>` body of the response
  - headers `<Object>` headers of the response
  - status `<number>` status code of the response
  - redirected `<boolean>` indicates response is the result of a redirect or not
  - url `<string>` the url of response
