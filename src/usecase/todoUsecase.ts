import TodoRepo from '../repo/todo';
import { TodoRequest } from '../dto/todoRequest';
import { UserRequest } from '../dto/userRequest';
import { Todo } from '../entity/Todo';
import { User } from '../entity/User';
import { CONSTANTS } from '../util/constant';
import { Error } from '../error';

const todoUsecase = (repo: TodoRepo) => ({
	create: async (t: TodoRequest) => {
		t.validate();

		const user = new User();
		user.id = t.userId;

		const userResponse = await repo.getTodayByUser(user);

		if (userResponse.length >= CONSTANTS.ACCOUNT_TYPE.GETDAILYLIMIT(t.userType)) {
			Error.exec('You have reached your daily limit', 400)
		}

		const todo = new Todo()

		todo.userId = t.userId;
		todo.task = t.task;
		todo.done = t.done;

		const todoResponse = await repo.create(todo);
		return todoResponse;
	},
	getAll: async() => {
		const todosResponse = await repo.getAll();
		return todosResponse;
	},
	getAllByUser: async(u: UserRequest) => {
		const user = new User();
		user.id = u.id;

		const todosResponse = await repo.getAllByUser(user);
		return todosResponse;
	}
});

export default todoUsecase;
