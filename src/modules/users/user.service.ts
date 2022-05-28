/* eslint-disable @typescript-eslint/no-unused-vars */
import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository, EntityManager, ObjectLiteral } from 'typeorm';
import { User } from './entities/user.entity';
import { BaseService } from '@modules/common/services/base.service';
import { CreateUserDto } from './dto/create-user.dto';
import { UpdateUserDto } from './dto/update-user.dto';
import { PermissionStatus } from '@modules/permissions/permission.status.enum';

const alias = 'users';

@Injectable()
export class UserService extends BaseService {
  constructor(@InjectRepository(User) private readonly repo: Repository<User>) {
    super(['role', 'role.permissions', 'userTaskConfigs']);
  }

  getRepository(manager?: EntityManager): Repository<User> {
    return manager ? manager.getRepository(User) : this.repo;
  }

  async findOne(conditions: ObjectLiteral, manager?: EntityManager): Promise<User> {
    const repo = this.getRepository(manager);

    return await repo.findOne({ where: conditions, relations: this.relations });
  }

  async findByUsername(username: string, manager?: EntityManager): Promise<User> {
    return await this.findOne({ username, deletedAt: null }, manager);
  }

  async findById(id: string, manager?: EntityManager): Promise<User> {
    return await this.findOne({ id, deletedAt: null }, manager);
  }

  async findByEmail(email: string, manager?: EntityManager): Promise<User | undefined> {
    return await this.findOne({ email: email, deletedAt: null }, manager);
  }

  async checkPermission(userId: string, resource: string, action: string, manager?: EntityManager): Promise<boolean> {
    const repo = this.getRepository(manager);
    const conds: string[] = [
      'users.id = :userId',
      'permissions.resource = :resource',
      'permissions.action = :action',
      'permissions.status = :status',
    ];

    const params = {
      userId,
      resource,
      action,
      status: PermissionStatus.Active,
    };

    const hasPermission = await repo
      .createQueryBuilder(alias)
      .innerJoin('users.role', 'role', 'role.deleted_at is null')
      .innerJoin('role.permissions', 'permissions', 'permissions.deleted_at is null')
      .where(conds.join(' AND '), params)
      .getCount();

    return hasPermission > 0;
  }

  create(createUserDto: CreateUserDto) {
    return 'This action adds a new user';
  }

  findAll() {
    return `This action returns all users`;
  }

  update(id: number, updateUserDto: UpdateUserDto) {
    return `This action updates a #${id} user`;
  }

  remove(id: number) {
    return `This action removes a #${id} user`;
  }
}
