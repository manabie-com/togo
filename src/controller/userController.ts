import { Request, Response } from 'express';
import UserUsecase from '../usecase/user';
import { userRequest } from '../dto/userRequest';

const userController = (uc: UserUsecase) => ({
	create: async (req: Request, res: Response) => {
		try {
			userRequest.username = req.body.username;
			userRequest.password = req.body.password;

			const user = await uc.create(userRequest);
			res.status(200).json(user);
		} catch (e: any) {
			res.status(e.code).json(e.message);
		}
	},
	authenticate: async (req: Request, res: Response) => {
		try {
			userRequest.username = req.body.username;
			userRequest.password = req.body.password;

			const token = await uc.authenticate(userRequest);

			res.status(200).json(token);
		} catch (e: any) {
			res.status(e.code).json(e.message);
		}
	}
})

export default userController;
