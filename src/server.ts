import "reflect-metadata";
import express, { Express, Request, Response } from 'express';
import bodyParser from 'body-parser';
import dotenv from 'dotenv';
dotenv.config();

import authorize from './middleware/auth';
import userController from './controller/userController';
import userUsecase from './usecase/userUsecase';
import userRepo from './repo/userRepo';

import todoController from './controller/todoController';
import todoUsecase from './usecase/todoUsecase';
import todoRepo from './repo/todoRepo';

import { db } from './db';

const app: Express = express();
const port = process.env.PORT || 3009;

app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());

app.get('/', (req: Request, res: Response) => {
	res.send('Welcome Togo');
})

app.post('/user', userController(userUsecase(userRepo(db))).create);
app.post('/authenticate', userController(userUsecase(userRepo(db))).authenticate);
app.post('/todo', authorize, todoController(todoUsecase(todoRepo(db))).create);

app.listen(port, () => {
	console.log(`Server running on port: ${port}`);
})