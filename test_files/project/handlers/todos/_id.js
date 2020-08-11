import http from "http";

const id = request.params["id"];
const url = `https://jsonplaceholder.typicode.com/todos/${id}`;

const res = http.request("GET", url);
response.write(res.body);
