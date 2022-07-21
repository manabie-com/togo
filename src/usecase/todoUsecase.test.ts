import { DataSource } from 'typeorm';
import todoUsecase from './todoUsecase';
import { todoRequest } from '../dto/todoRequest';
import { userRequest } from '../dto/userRequest';

describe('Handles Users usecases', () => {
	/* Setup requirements */
	const Response = <Type>(data: Type): Promise<Type> => new Promise((res, rej) => res(data));
	const db = new DataSource({
		type: 'mysql',
		host: process.env.DB_HOST,
		port: 3306,
		username: process.env.DB_USER,
		password: process.env.DB_PASSWORD,
		database: process.env.DB_NAME,
	});
	type Todo = {id: number; userId: number; task: string, done: boolean, createdAt: Date}
	todoRequest.task = 'My task';
	todoRequest.userId = 1;
	todoRequest.userType = 'BASIC';
	userRequest.id = 1;

	/* Defined Expectations */
	const creationExpectation = true;
	const getTodosExpectation: Todo[] = [{id: 1, userId: 1, task: 'Test', done: false, createdAt: new Date()}];
	const getTodosExceedExpectation: Todo[] = new Array(5).map(() => getTodosExpectation[0]);

	/* Mock repository */
	const todoRepository = (db: DataSource) => ({
		create: () => Response<boolean>(creationExpectation),
		getAll: () => Response<Todo[]>(getTodosExpectation),
		getAllByUser: () => Response<Todo[]>(getTodosExceedExpectation),
		getTodayByUser: () => Response<Todo[]>(getTodosExpectation),
	})

	/* Tests */
	test('Handles task creation', async () => {
		const result = await todoUsecase(todoRepository(db)).create(todoRequest);
		expect(result).toEqual(creationExpectation);
	})
	test('Should throw error when task creation exceeds limit', () => {
		expect(async() => await todoUsecase(todoRepository(db)).create(todoRequest));
	})
	test('Handles get all task', async () => {
		const result = await todoUsecase(todoRepository(db)).getAll();
		expect(result).toEqual(getTodosExpectation);
	})
	test('Handles get all task by user', async () => {
		const result = await todoUsecase(todoRepository(db)).getAllByUser(userRequest)
		expect(result).toEqual(getTodosExceedExpectation);
	})
});
