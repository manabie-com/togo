import { Injectable } from '@nestjs/common';
import { EntityRepository } from 'typeorm/decorator/EntityRepository';

import { Task } from '../entities/task.entity';
import { BaseRepository } from './base.repository';

@Injectable()
@EntityRepository(Task)
export class TaskRepository extends BaseRepository<Task> {}
