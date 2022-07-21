import { hash, compare } from './hasher';

describe('Should hash string and return true when compared to original', () => {
	const plainText = 'my plaint text';
	let hashed = '';

	test('Hash a plaint text', async () => {
		let hashed = await hash(plainText);
		expect(hashed).not.toHaveLength(0);
	})

	test('Hashed string should be equal to original', () => {
		expect(compare(plainText, hashed)).toBeTruthy();
	})
})
