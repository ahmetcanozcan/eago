import Todo from "models/todo";

const id = request.params.id;
const todo = Todo.getFromID(id);

if (!todo) {
  response.send({ success: false });
  process.exit();
}

const body = request.body.json();

if (body.content) {
  todo.content = body.content;
}

if (body.priority) {
  todo.priority = body.priority;
}

const res = todo.save();

response.send({ success: true });
