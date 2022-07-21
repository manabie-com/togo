import { DataSource } from 'typeorm';
import userUsecase from './userUsecase';
import { UserResponse, userResponse } from '../dto/userResponse';
import { AuthUserResponse, authUserResponse } from '../dto/authResponse';
import { userRequest } from '../dto/userRequest';
import { authRequest } from '../dto/authRequest';

describe('Handles Users usecases', () => {
	/* Setup requirements */
	process.env.JWT_SECRET = 'secret';
	const Response = <Type>(data: Type): Promise<Type> => new Promise((res, rej) => res(data));
	const db = new DataSource({
		type: 'mysql',
		host: process.env.DB_HOST,
		port: 3306,
		username: process.env.DB_USER,
		password: process.env.DB_PASSWORD,
		database: process.env.DB_NAME,
	});
	const username = 'Test';
	const password = 'Test';
	const userId = 1;
	userRequest.id = userId;
	userRequest.username = username;
	userRequest.password = password;

	authRequest.username = username;
	authRequest.password = password;

	authUserResponse.password = '$2b$10$79Y6ynm2DygLMj7GDzLekes2iv0wYIZ5O8/RomrPHjzZ1au5RVU4a';
	authUserResponse.username = username;
	authUserResponse.type = 'BASIC';
	authUserResponse.id = userId;

	/* Defined Expectations */
	const creationExpectation = userResponse;
	const authUserExpectation = authUserResponse;
	const upgradeExpectation = true;

	/* Mock repository */
	const userRepository = (db: DataSource) => ({
		create: () => Response<UserResponse>(creationExpectation),
		authenticate: () => Response<AuthUserResponse>(authUserExpectation),
		upgrade: () => Response<boolean>(upgradeExpectation),
	})

	/* Tests */
	test('Handles user creation', async () => {
		const result = await userUsecase(userRepository(db)).create(userRequest)
		expect(result).toEqual(creationExpectation);
	})
	test('Handle user authentication', async () => {
		const result = await userUsecase(userRepository(db)).authenticate(authRequest);
		expect(() => result).not.toThrow();
	})
	test('Handles user upgrade', async () => {
		const result = await userUsecase(userRepository(db)).upgrade(userRequest)
		expect(result).toEqual(upgradeExpectation);
	})
});
