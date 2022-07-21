/* 
 * DTO - Data Transfer Object
 * Handles data passing from controller to useCase
 * Also Handles input validation
 */
export interface TodoResponse {
	id: number;
	userId: number;
	task: string;
	done: boolean;
}

export const todoResponse: TodoResponse = {
	id: 0,
	userId: 0,
	task: '',
	done: false,
}
