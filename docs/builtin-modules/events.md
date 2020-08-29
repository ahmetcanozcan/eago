# Events

the `events` module provide communicatin with events

## emit(name[,payload])

- name `string`
- payload `any`
- returns `any`

emits a task to given event and wait for it response

```javascript
import { emit } from "events";

const comment = {
  author: "HAL",
  content: "This conversation can serve no purpose anymore",
};

const ok = emit("log/comment", comment);
```
