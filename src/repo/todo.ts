import { Todo } from '../entity/Todo';

interface TodoRepo {
	create(t: Todo): Promise<boolean>
}

export default TodoRepo;
