import { DataSource } from 'typeorm';
import { Todo } from '../entity/Todo';
import { Error } from '../error';

const todoRepo = (db: DataSource) => ({
	create: async (t: Todo) => {
		try {
			await db.manager.save(t);
			return true;
		} catch (e: any) {
			console.log(e);
			Error.exec('Database error', 500);
			return false;
		}
	},
})

export default todoRepo;
