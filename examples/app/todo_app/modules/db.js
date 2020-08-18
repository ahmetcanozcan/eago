import { Postgres } from "psql";

const username = process.env.DB_USER || "postgres";
const password = process.env.DB_PASS;
const dbname = process.env.DB_NAME || "test";
const host = process.env.DB_HOST || "localhost";
const port = process.env.DB_PORT || 5432;

const connectionURI =
  `host=${host} port=${port} user=${username}` +
  ` password=${password} dbname=${dbname} sslmode=disable`;

const db = new Postgres(connectionURI);

if (typeof db === "undefined") {
  console.error("Can not connect to DB");
  process.exit();
}
export default db;
