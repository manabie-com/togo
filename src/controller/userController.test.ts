import userController from './userController';

describe('Handles Users controller', () => {
	/* Setup requirements */
	const req = {
		body: {
			username: 'Test',
			password: 'test',
		},
		user: { id: 1 },
	}
	const res = {
		status: (code: number) => ({
			json: (data: any) => data,
		}),
	}
	const Response = (data: any) => new Promise((res, rej) => res(data));

	/* Defined Expectations */
	const creationExpectation = {
		id: 1,
		username: 'Test',
		type: 'BASIC',
	}
	const authenticationExpectation = { token: 'xxxxxxxxxxxxxxxxxxx' }
	const upgradeResponse = true;

	/* Mock usecase */
	const userUsecase = () => ({
		create: () => Response(creationExpectation),
		authenticate: () => Response(authenticationExpectation),
		upgrade: () => Response(upgradeResponse),
	})

	/* Tests */
	test('Handles user creation', async () => {
		const result = await userController(userUsecase()).create(req, res)
		expect(result).toEqual(creationExpectation);
	})
	test('Handles user authentication', async () => {
		const result = await userController(userUsecase()).authenticate(req, res)
		expect(result).toEqual(authenticationExpectation);
	})
	test('Handles user upgrade', async () => {
		const result = await userController(userUsecase()).upgrade(req, res)
		expect(result).toEqual(upgradeResponse);
	})
});
