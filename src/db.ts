import { DataSource, DataSourceOptions } from 'typeorm';
import { User } from './entity/User';

const opt: DataSourceOptions = {
	type: 'mysql',
	host: 'localhost',
	port: 3306,
	username: 'root',
	password: 'password',
	database: 'togo',
	synchronize: true,
	logging: false,
	dropSchema: true,
	entities: [User],
	migrations: [],
	subscribers: [],
}

export const db = new DataSource(opt);
db.initialize();
