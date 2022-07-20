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
