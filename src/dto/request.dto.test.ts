import { authRequest } from './authRequest';
import { userRequest } from './userRequest';
import { todoRequest } from './todoRequest';

describe('Validation for fields', () => {
	test('Should throw an error when passing empty fields', () => {
		expect(() => authRequest.validate()).toThrow();
	})

	test('Should throw an error when passing empty fields', () => {
		expect(() => userRequest.validate()).toThrow();
	})

	test('Should throw an error when passing empty fields', () => {
		expect(() => todoRequest.validate()).toThrow();
	})
})
