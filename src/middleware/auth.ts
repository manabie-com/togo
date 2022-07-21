/*
 * Handles authorization for protected routes
 */
import { Response, NextFunction } from 'express';
import { jwt } from '../util/jwt';

const authorize = (req: any, res: Response, next: NextFunction) => {
	try {
	const authHeader = req.headers.authorization;
	const token = authHeader?.split(' ')[1];

	const user = jwt.verify(token);

	req.user = user;
	
	next();
	} catch (e: any) {
		res.status(403).json({ message: 'Authorization required!'});
	}
}

export default authorize;
