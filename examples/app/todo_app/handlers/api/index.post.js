import Todo from "models/todo";

const body = request.body.json();

const todo = new Todo(body.content, body.priority);

try {
  const res = todo.save();
} catch (e) {
  response.send({ success: false });
  console.error(err);
}

response.send({ success: true });
