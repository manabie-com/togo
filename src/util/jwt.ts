import Jwt from 'jsonwebtoken';

export const jwt = {
	generate: (user: any) => {
		return Jwt.sign(user, process.env.JWT_SECRET as string);
	},
	verify: (token: string) => {
		return Jwt.verify(token, process.env.JWT_SECRET as string);
	}
}
