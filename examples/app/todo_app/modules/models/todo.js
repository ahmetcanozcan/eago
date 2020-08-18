import db from "db";
import { generateID } from "util";

export default class Todo {
  /**
   * @param {string} content
   * @param {number} priority
   */
  constructor(content, priority) {
    this.id = -1;
    this.content = content;
    this.priority = priority;
  }

  static getFromID(id) {
    const res = db.exec(`SELECT * FROM todos WHERE id = ${id}`)[0];
    if (!res) {
      return;
    }
    const todo = new Todo(res.content, res.priority);
    todo.id = id;
    return todo;
  }

  static getAll() {
    return db.exec("SELECT * FROM todos");
  }

  delete() {
    if (this.id == -1) return;
    return db.exec(`DELETE FROM todos WHERE id = ${this.id}`);
  }

  save() {
    if (this.id == -1) {
      // Create
      const q = `INSERT INTO todos VALUES (${generateID()} , '${
        this.content
      }',${this.priority})`;
      console.log(q);
      return db.exec(q);
    } else {
      return db.exec(`UPDATE todos 
      SET content = '${this.content}',
      priority = ${this.priority}
      WHERE id = ${this.id}`);
    }
  }
}
