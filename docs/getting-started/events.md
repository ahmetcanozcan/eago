# Events

an `Event` queues the asynchronous task emits and executes the tasks by one by one synchronously.

You can emit tasks to an `event` by using [`events`](builtin-modules/events.md) module

```javascript
// events/index/count.js

let count = 0;

listen((payload) => {
  count++;
  return count;
});
```

```javascript
// handlers/index.get.js
import events from "events";

const count = events.emit("index/count");

response.send(`Request Count: ${count} `);
```
