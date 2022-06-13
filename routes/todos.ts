import express, { Request, Response } from "express";
import { TodoRepositoryFactory } from "../factories/TodoRepositoryFactory";
import { UserRepositoryFactory } from "../factories/UserRepositoryFactory";

let router = express.Router();

router.route("/").post(async (req: Request, res: Response) => {
  try {
    const params = req.body;
    const todoRepository = await TodoRepositoryFactory.createInstance();
    const userRepository = await UserRepositoryFactory.createInstance();
    const { userId, task } = params;

    if (!userId) {
      throw new Error("User ID is missing");
    }

    if (!task) {
      throw new Error("Task is missing");
    }

    // Find user based on the given userId
    const user = await userRepository.getUserById(userId);

    if (!user) {
      throw new Error("User not found");
    } else {
      // Check if user has reached daily limit
      const tasks = await todoRepository.getCurrentTasksByUserId(userId);
      const { dailyMaximumTasks } = user;

      if (tasks >= dailyMaximumTasks) {
        throw new Error("Maximum daily limit reached");
      }
    }

    const todo = await todoRepository.saveTodo(params);

    res.json({
      status: "success",
      data: todo,
    });
  } catch (e) {
    return res.json({
      status: "error",
      data: `${e}`,
    });
  }
});

export default router;
