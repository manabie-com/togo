import { Repository } from 'typeorm';

export class BaseRepository<T> extends Repository<T> {}
