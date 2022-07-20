import bcrypt from 'bcrypt';

const SALTROUND = 10;

export const hash = async (str: string) => {
	const hashed = await bcrypt.hash(str, SALTROUND);
	return hashed;
}

export const compare = async (str: string, hashed: string)  => {
	const result = await bcrypt.compare(str, hashed);
	return result;
}
