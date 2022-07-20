import { TodoRequest } from '../dto/todoRequest';

interface TodoUsecase {
	create(t: TodoRequest): Promise<boolean>
}

export default TodoUsecase;
