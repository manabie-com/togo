/* 
 * DTO - Data Transfer Object
 * Handles data passing from controller to useCase
 * Also Handles input validation
 */
import { Error } from '../error';

export interface TodoRequest {
	userId: number;
	userType: string;
	task: string;
	done: boolean;
	createdAt: Date;
	validate(): void;
}

export const todoRequest: TodoRequest = {
	userId: 0,
	userType: '',
	task: '',
	done: false,
	createdAt: new Date(),
	validate() {
		if (!this.userId || this.userId == 0) Error.exec('Not signed in', 401)
		if (!this.task || this.task == '') Error.exec('task is empty or not valid', 400)
	}
}
