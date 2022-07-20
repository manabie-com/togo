import { Error } from '../error';

export interface TodoRequest {
	userId: number;
	task: string;
	done: boolean;
	createdAt: Date;
	validate(): void;
}

export const todoRequest: TodoRequest = {
	userId: 0,
	task: '',
	done: false,
	createdAt: new Date(),
	validate() {
		if (!this.userId || this.userId == 0) Error.exec('Not signed in', 401)
		if (!this.task || this.task == '') Error.exec('task is empty or not valid', 400)
	}
}
