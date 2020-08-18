import Todo from "models/todo";

const id = request.params.id;
const todo = Todo.getFromID(id);

response.send({ success: typeof todo != "undefined", data: todo });
