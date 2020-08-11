import SQL from "psql";
import { generateDbURI } from "./util";
const sql = new SQL();

const connectionURI = generateDbURI(
  "localhost",
  5432,
  "test",
  process.env["DB_USR"],
  process.env["DB_PASS"]
);

const res = sql.connect(connectionURI);

console.log("connection res", res);

export default sql;
