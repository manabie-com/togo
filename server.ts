import express from "express";
import todos from "./routes/todos";
import db from "./providers/typeorm";

const app = express();

// Middleware
app.use(express.json());

app.use("/api/todos", todos);

app.listen(5000, async (error) => {
  const connection = db.getConnection();

  if (!connection) {
    await db.initialize();
  }

  if (error) {
    console.log(`Error: ${error}`);
  }

  console.log(`Server is listening on port 5000`);
});

export default app;
