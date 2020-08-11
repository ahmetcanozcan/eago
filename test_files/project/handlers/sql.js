import db from "db";

const res = db.exec("SELECT * FROM comments");

response.write(JSON.stringify(res));
