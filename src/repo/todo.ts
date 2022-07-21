/*
 * Repository Interface which defines repository signature
 */
import { Todo } from '../entity/Todo';
import { User } from '../entity/User';

interface TodoRepo {
	create(t: Todo): Promise<boolean>
	getAll(): Promise<Todo[]>
	getAllByUser(u: User): Promise<Todo[]>
	getTodayByUser(u: User): Promise<Todo[]>
}

export default TodoRepo;
