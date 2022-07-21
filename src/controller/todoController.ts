import { Request, Response } from 'express';
import TodoUsecase from '../usecase/todo';
import { todoRequest } from '../dto/todoRequest';
import { userRequest } from '../dto/userRequest';

const todoController = (tc: TodoUsecase) => ({
	create: async(req: any, res: Response) => {
		try {
			todoRequest.task = req.body.task;
			todoRequest.userId = req.user.id;
			todoRequest.userType = req.user.type;
			const todo = await tc.create(todoRequest);
			res.status(201).json(todo);
		} catch (e: any) {
			res.status(e.code).json(e.message);
		}
	},
	getAll: async(req: Request, res: Response) => {
		try {
			const todos = await tc.getAll();
			res.status(200).json(todos);
		} catch (e: any) {
			res.status(e.code).json(e.message);
		}
	},
	getAllByUser: async(req: Request, res: Response) => {
		try {
			userRequest.id = Number(req.params.userId);
			const todos = await tc.getAllByUser(userRequest);
			res.status(200).json(todos);
		} catch (e: any) {
			res.status(e.code).json(e.message);
		}
	},
});

export default todoController;
