import { Request, Response } from 'express';
import UserUsecase from '../usecase/user';
import { userRequest } from '../dto/userRequest';

const userController = (uc: UserUsecase | any) => ({
	create: async (req: Request | any, res: Response | any) => {
		try {
			userRequest.username = req.body.username;
			userRequest.password = req.body.password;

			const user = await uc.create(userRequest);
			return res.status(200).json(user);
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
	},
	upgrade: async (req: any, res: Response) => {
		try {
			userRequest.id = req.user.id;

			const user = await uc.upgrade(userRequest);
			res.status(200).json(user);
		} catch (e: any) {
			res.status(e.code).json(e.message);
		}
	}
})

export default userController;
