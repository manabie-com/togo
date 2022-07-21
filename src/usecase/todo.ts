import { TodoRequest } from '../dto/todoRequest';
import { UserRequest } from '../dto/userRequest';
import { Todo } from '../entity/Todo';

interface TodoUsecase {
	create(t: TodoRequest): Promise<boolean>
	getAll(): Promise<Todo[]>
	getAllByUser(u: UserRequest): Promise<Todo[]>
}

export default TodoUsecase;
