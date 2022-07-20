import { Response } from 'express';
import TodoUsecase from '../usecase/todo';
import { todoRequest } from '../dto/todoRequest';

const todoController = (tc: TodoUsecase) => ({
	create: async(req: any, res: Response) => {
		try {
			todoRequest.task = req.body.task;
			todoRequest.userId = req.user.id;
			const todo = await tc.create(todoRequest);
			res.status(201).json(todo);
		} catch (e: any) {
			res.status(e.code).json(e.message);
		}
	},
});

export default todoController;
