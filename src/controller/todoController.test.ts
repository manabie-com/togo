import todoController from './todoController';

describe('Handles Users controller', () => {
	/* Setup requirements */
	const req = {
		body: {
			task: 'Sampe task',
			password: 'test',
		},
		user: { id: 1, type: 'BASIC' },
		params: { userId: 1 },
	}
	const res = {
		status: (code: number) => ({
			json: (data: any) => data,
		}),
	}
	const Response = (data: any) => new Promise((res, rej) => res(data));

	/* Defined Expectations */
	const creationExpectation = { created: true }
	const getTodosExpectation = [{id: 1, task: 'test'}];

	/* Mock usecase */
	const todoUsecase = () => ({
		create: () => Response(creationExpectation),
		getAll: () => Response(getTodosExpectation),
		getAllByUser: () => Response(getTodosExpectation),
	})

	/* Tests */
	test('Handles task creation', async () => {
		const result = await todoController(todoUsecase()).create(req, res)
		expect(result).toEqual(creationExpectation);
	})
	test('Handles get all task', async () => {
		const result = await todoController(todoUsecase()).getAll(req, res)
		expect(result).toEqual(getTodosExpectation);
	})
	test('Handles get all task by user', async () => {
		const result = await todoController(todoUsecase()).getAllByUser(req, res)
		expect(result).toEqual(getTodosExpectation);
	})
});
