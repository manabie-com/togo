import { Request, Response } from 'express';
import TodoUsecase from '../usecase/todo';
import { todoRequest } from '../dto/todoRequest';
import { userRequest } from '../dto/userRequest';

/* 
 * Route Handler also holds Usecase interface
 * to easily create a mock during test
 */
const todoController = (tc: TodoUsecase | any) => ({
	create: async(req: any, res: Response | any) => {
		try {
			todoRequest.task = req.body.task;
			todoRequest.userId = req.user.id;
			todoRequest.userType = req.user.type;
			const todo = await tc.create(todoRequest);
			return res.status(201).json(todo);
		} catch (e: any) {
			res.status(e.code).json(e.message);
		}
	},
	getAll: async(req: Request | any, res: Response | any) => {
		try {
			const todos = await tc.getAll();
			return res.status(200).json(todos);
		} catch (e: any) {
			res.status(e.code).json(e.message);
		}
	},
	getAllByUser: async(req: Request | any, res: Response | any) => {
		try {
			userRequest.id = Number(req.params.userId);
			const todos = await tc.getAllByUser(userRequest);
			return res.status(200).json(todos);
		} catch (e: any) {
			res.status(e.code).json(e.message);
		}
	},
});

export default todoController;
