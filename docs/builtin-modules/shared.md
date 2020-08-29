# Shared Data

Eago runs every handler on a goroutine, And In concurrent systems accessing shared data and storing state it's hard. `shared` module provides to store state in a secure and reliable way.

example usage:

```javascript
import shared from "shared";
```

## shared.get(name)

- name `<string>`
- returns `<any>`

the `shared.get()` retuns shared data

## shared.set(name,value)

- name `<string>`
- value `<any>`

the `shared.set()` sets or update variable `name` as `value`

## shared.update(name,cb)

- name `<string>`
- cb `<Function>`
  - val `<any>`
  - returns `<any>`

the `shared.update()` get and sets value in the same time, it's usefull for tasks like incrementinig a value

```javascript
import shared from "shared";

/**
 * This snippet, gets current value of count,
 * increment it and set it as `count` variable.
 */
const count = shared.update("count", (val) => {
  val = val || 0; // if val is undefined set it 0
  return ++val;
});

response.send("Count: " + count);
```
