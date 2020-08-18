import Todo from "models/todo";

const res = Todo.getAll();

if (res) {
  response.send({ success: true, data: res });
} else {
  response.send({ success: false });
}
