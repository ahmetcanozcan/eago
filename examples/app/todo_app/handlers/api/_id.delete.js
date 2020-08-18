import Todo from "models/todo";

const id = request.params.id;

const todo = Todo.getFromID(id);

if (!todo) {
  response.status(300);
  response.send({ succes: false, message: "Todo not found" });
} else {
  const res = todo.delete();
  response.send({ success: true });
}
