import { jwt } from './jwt';

describe('Should generate token and verify token which contains user data', () => {

	process.env.JWT_SECRET = 'secret';

	const user = {
		name: 'Test',
		age: 10,
	}

	let token = '';

	test('Should issue token with data', () => {
		token = jwt.generate(user);
		expect(token).not.toHaveLength(0)
	})

	test('Should verify token if contains exact data', () => {
		const got: any = jwt.verify(token);
		expect(user).toEqual({ name: got.name, age: got.age });
	})
})
