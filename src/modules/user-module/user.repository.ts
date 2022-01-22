import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Constants } from '../../utils/constants';
import { Repository, Connection } from 'typeorm';
import { RegisterDTO } from './dto/register.dto';
import { UserEntity } from './entities/user.entity';
import { FindUserDTO } from './dto/find.dto';
import { ListUser } from './interfaces/listUser.interface';
import { TaskEntity } from './entities/task.entity';
import { TaskDTO } from './dto/task.dto';
import { ConfigureMaxTaskDTO } from './dto/configure_max_task.dto';


@Injectable()
export class UserRepository {
    constructor(
        @InjectRepository(UserEntity)
        private readonly userRepository: Repository<UserEntity>,
        @InjectRepository(TaskEntity)
        private readonly taskRepository: Repository<TaskEntity>,
        private readonly connection: Connection,

    ) { }

    async register(data: RegisterDTO, isAdmin: boolean): Promise<UserEntity> {
        const newUser = this.userRepository.create({
            email: data.email,
            password: data.password,
            fullName: data.fullName,
            role: isAdmin ? Constants.ADMIN_ROLE : Constants.USER_ROLE
        });
        try {
            const result = await this.userRepository.save(newUser);
            return result;
        } catch (e) {
            return null;
        }
    }

    async findUserByEmail(email: string): Promise<UserEntity> {
        return await this.userRepository.findOne({ email })
    }

    async find(query: FindUserDTO): Promise<ListUser> {
        // Get count
        let queryBuilder = this.userRepository.createQueryBuilder("user")
        const [
            total,
            listUsers
        ] = await Promise.all([
            queryBuilder.getCount(),
            queryBuilder.take(query.perPage)
                .skip(query.perPage * (query.pageIndex - 1))
                .orderBy("user.id", "ASC")
                .getMany()
        ]);
        return {
            listUsers: listUsers.map(e => {
                delete e.password;
                return e;
            }),
            pageIndex: query.pageIndex,
            totalPage: Math.ceil(total / query.perPage)
        }
    }

    private _createTaskEntity(userId: number, task: TaskDTO): TaskDTO {
        const userEntity = new UserEntity({
            id: userId
        })
        return this.taskRepository.create({
            name: task.name,
            note: task.note,
            owner: userEntity,
            created_at: new Date()
        });
    }

    async addNewTask(userId: number, task: TaskDTO): Promise<boolean> {
        try {
            const date = new Date();
            const dateString = date.toISOString().split('T')[0]
            // Start transaction
            await this.connection.transaction(async manager => {
                await manager.update(UserEntity, { id: userId }, {
                    task_left: () => `IF (DATE(lasted_date_task) = '${dateString}', task_left - 1, max_task_per_day - 1)`,
                    lasted_date_task: date
                });
                await manager.save(this._createTaskEntity(userId, task));
            });
            return true
        } catch (e) {
            console.log(e);
            return false
        }
    }

    private async _getTotalTaskToday(userId: number): Promise<number> {
        const date = new Date();
        const dateString = date.toISOString().split('T')[0]
        return await this.taskRepository.createQueryBuilder()
            .where("ownerId = :userId", { userId })
            .andWhere("DATE(created_at) = :dateString", {
                dateString
            })
            .getCount();
    }
    async configureMaxTask(config: ConfigureMaxTaskDTO): Promise<boolean> {
        try {
            // Get total task in this day
            const totalTaskToday = await this._getTotalTaskToday(config.id);
            const taskLeft = Math.max(config.maxTask - totalTaskToday, 0);
            await this.userRepository.createQueryBuilder()
                .update(UserEntity)
                .set({
                    task_left: taskLeft,
                    max_task_per_day: config.maxTask
                })
                .where("id = :id", { id: config.id })
                .execute();
            return true;
        } catch (e) {
            return false;
        }
    }
}