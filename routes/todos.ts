import express, { Request, Response } from "express";
import { TodoRepositoryFactory } from "../factories/TodoRepositoryFactory";

let router = express.Router();

router.route("/").post(async (req: Request, res: Response) => {
  try {
    const params = req.body;
    const todoRepository = await TodoRepositoryFactory.createInstance();
    const todo = await todoRepository.saveTodo(params);

    res.json(todo);
  } catch (e) {
    console.log(`Error: ${e}`);
    return res.sendStatus(500);
  }
});

export default router;
