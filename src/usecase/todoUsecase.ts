import TodoRepo from '../repo/todo';
import { TodoRequest } from '../dto/todoRequest';
import { Todo } from '../entity/Todo';

const todoUsecase = (repo: TodoRepo) => ({
	create: async (t: TodoRequest) => {
		t.validate();

		const todo = new Todo()

		todo.userId = t.userId;
		todo.task = t.task;
		todo.done = t.done;

		const todoResponse = await repo.create(todo);
		return todoResponse;
	},
});

export default todoUsecase;
