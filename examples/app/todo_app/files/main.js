const baseURL = "http://localhost:3000/api";

new Vue({
  el: "#app",
  data: {
    todos: [],
    title: "TODO LIST",
    content: "",
    priority: "",
  },
  created() {
    this.getTodos();
  },
  computed: {
    listedTodos() {
      return this.todos;
    },
  },
  methods: {
    getTodos() {
      fetch(`${baseURL}/`)
        .then((res) => res.json())
        .then((json) => {
          if (json.success) {
            this.todos = json.data;
          }
        });
    },
    deleteTodo(id) {
      fetch(`${baseURL}/${id}`, {
        method: "delete",
      })
        .then((res) => res.json())
        .then((json) => {
          if (json.success) {
            this.getTodos();
          }
        });
    },
    incPriority(id, inc) {
      const t =
        parseInt(this.todos.filter((todo) => todo.id == id)[0].priority) + inc;
      fetch(`${baseURL}/${id}`, {
        method: "put",
        body: JSON.stringify({ priority: t }),
      })
        .then((res) => res.json())
        .then((json) => {
          if (json.success) {
            this.getTodos();
          }
        });
    },
    addTodo() {
      fetch(`${baseURL}/`, {
        method: "post",
        body: JSON.stringify({
          content: this.content,
          priority: this.priority,
        }),
      })
        .then((res) => res.json())
        .then((json) => {
          if (json.success) {
            this.getTodos();
            this.content = "";
            this.priority = "";
          }
        });
    },
  },
});
