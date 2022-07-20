import { DataSource, DataSourceOptions } from 'typeorm';
import { User } from './entity/User';
import { Todo } from './entity/Todo';

const opt: DataSourceOptions = {
	type: 'mysql',
	host: process.env.DB_HOST,
	port: 3306,
	username: process.env.DB_USER,
	password: process.env.DB_PASSWORD,
	database: process.env.DB_NAME,
	synchronize: true,
	logging: false,
	dropSchema: true,
	entities: [User, Todo],
	migrations: [],
	subscribers: [],
}

export const db = new DataSource(opt);
db.initialize();
