import { Request, Response } from 'express';
import UserUsecase from '../usecase/user';
import userRequest from '../dto/userRequest';

const userController = (uc: UserUsecase) => async (req: Request, res: Response) => {
	userRequest.username = req.body.username;
	userRequest.password = req.body.password;

	const user = await uc.create(userRequest);
	res.status(200).json(user);
}

export default userController;
