import db from "db";

const res = db.exec(`
  CREATE TABLE IF NOT EXISTS todos (
    id INT UNIQUE,
    content VARCHAR(75) NOT NULL,
    priority INT NOT NULL
  )
`);
