# PostgreSQL

## Class : psql.Postgres

- ### exec(query)

  - query `<string>`
  - returns `<Object>` | `<number>`

  executes sql query on psql server. the `exec()` method returns several type of result depends on given query type

```javascript
import { Postgres } from "psql";

const db = new Postgres(connectionURI);

// `db.exec` returns array of key value map
// for this example, the result can be:
// [
//   { author: "John", content: "Nice work!" },
//   { author: "HAL", content: "this conversation can serve no purpose anymore" },
// ];
let res = db.exec("SELECT author,content FROM messages");

// for create and update operations `db.exec` returns count of effected lines
res = db.exec(
  "INSERT INTO messages (author,content) VALUES ('Dave','Open the Pod bay doors')"
);
```

- ### close()

  closes PostgreSQL server connection
