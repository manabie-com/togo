import express from "express";
import todos from "./routes/todos";

const app = express();

app.use("/api/todos", todos);

app.listen(5000, async (error) => {
  if (error) {
    console.log(`Error: ${error}`);
  }

  console.log(`Server is listening on port 5000`);
});

export default app;
