/*
 * Concrete implementation for
 * Todo Repository
 */
import { DataSource } from 'typeorm';
import { Todo } from '../entity/Todo';
import { User } from '../entity/User';
import { Error } from '../error';

const todoRepo = (db: DataSource) => ({
	create: async (t: Todo) => {
		try {
			await db.manager.save(t);
			return true;
		} catch (e: any) {
			console.log(e);
			Error.exec('Database error', 500);
			return false;
		}
	},
	getAll: async () => {
		try {
			const repo = db.getRepository(Todo);
			const todos = await repo.find();
			return todos;
		} catch (e: any) {
			const result: Todo[] = [];
			Error.exec('Database error', 500);
			return result;
		}
	},
	getAllByUser: async (u: User) => {
		try {
			const repo = db.getRepository(Todo);
			const todos = await repo.findBy({ userId: u.id });
			return todos;
		} catch (e: any) {
			const result: Todo[] = [];
			Error.exec('Database error', 500);
			return result;
		}
	},
	getTodayByUser: async (u: User) => {
		try {
			const repo = db.getRepository(Todo);
			const todos = await repo.createQueryBuilder('todo')
			.where('todo.createdAt >= CURDATE() AND todo.createdAt <= DATE_ADD(CURDATE(), INTERVAL 1 DAY)')
			.andWhere('todo.userId = :id', {id: u.id})
			.getRawMany();
			return todos;

		} catch (e: any) {
			const result: Todo[] = [];
			Error.exec('Database error', 500);
			return result;
		}
	}
})

export default todoRepo;
