## Install

Eago can be installed by using go toolchain

```
go get github.com/ahmetcanozcan/eago
```

## First Project

In a few steps, you can create your first project easily using eago cli.

```
  eago new project <project name>
  cd <project name>
```

## Writing Content

after creating the project, you can see the files and folders in project directory

- `eago.json` contains information about the eago project such as dependincies, static file server path, not found page, port number, etc.

- `start.js` the script that runs before application server

- `handlers/` the directory that contains all scripts that handles http requests, and response that requests

- `events/` the directory that contains all event scripts

- `modules/` the directory that contains reusable scripts

- `packages/` the directory that contains third-party bundles of modules. pacakges can be installed using `eago install`

- `files/` Eago does not allow to access the whole file system of the machine. you can only read and write files from that folder

## Hello World

Now, let's write a basic hello world application.

First create a handler file called `index.get.js` in `handlers` folder.

```javascript
// handlers/index.get.js

response.send("HELLO WORLD");
```

Then start the application and open http://localhost:3000 on browser

```
eago start
```
