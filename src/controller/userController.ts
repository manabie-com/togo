import { Request, Response } from 'express';
import UserUsecase from '../usecase/user';
import userRequest from '../dto/userRequest';

const userController = (uc: UserUsecase) => async (req: Request, res: Response) => {
	try {
	userRequest.username = req.body.username;
	userRequest.password = req.body.password;

	const user = await uc.create(userRequest);
	res.status(200).json(user);
	} catch (e: any) {
		res.status(e.code).json(e.message);
	} finally {
		res.status(500).json({ message: 'Server error 500' });
	}
}

export default userController;
