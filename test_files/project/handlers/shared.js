import { get, set } from "shared";

const i = get("count") || 0;

response.write("Count: " + i);

set("count", i + 1);
