import { Injectable } from '@nestjs/common';
import { EntityRepository } from 'typeorm/decorator/EntityRepository';

import { User } from '../entities/user.entity';
import { BaseRepository } from './base.repository';

@Injectable()
@EntityRepository(User)
export class UserRepository extends BaseRepository<User> {}
