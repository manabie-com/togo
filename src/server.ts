import "reflect-metadata";
import express, { Express, Request, Response } from 'express';
import bodyParser from 'body-parser';
import dotenv from 'dotenv';

import userController from './controller/userController';
import userUsecase from './usecase/userUsecase';
import userRepo from './repo/userRepo';

import { db } from './db';

dotenv.config();
const app: Express = express();
const port = process.env.PORT || 3009;

app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());

app.get('/', (req: Request, res: Response) => {
	res.send('Welcome Togo');
})

app.post('/user', userController(userUsecase(userRepo(db))));

app.listen(port, () => {
	console.log(`Server running on port: ${port}`);
})
