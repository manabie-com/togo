import { Test, TestingModule } from '@nestjs/testing';
import { CreateTaskDTO } from './../context';
import { TaskService } from './../task.service';
import { PriorityEnum, StatusEnum } from './../../../common';
import { Task } from '../task.entity';

import { getRepositoryToken } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { UserService } from './../../user/user.service';
import { User } from './../../user/user.entity';
import { BadRequestException } from '@nestjs/common';
import { set, reset } from 'mockdate';

const createTaskDto: CreateTaskDTO = {
  assignee_id: 1,
  title: 'title task',
  description: 'description task',
  priority: PriorityEnum.Medium,
  status: StatusEnum.Todo,
};

const updateTaskDto: CreateTaskDTO = {
  assignee_id: 1,
  title: 'title task updated',
  description: 'description task updated',
  priority: PriorityEnum.High,
  status: StatusEnum.InProgess,
};

const date = new Date('2020-04-13T18:09:12.451Z');
set(date); // Any request to Date will return this date

const created_time = date;
const updated_time = date;

const tasks: Task[] = [
  {
    id: 1,
    assignee_id: 1,
    title: 'title task',
    description: 'description task',
    priority: PriorityEnum.Medium,
    status: StatusEnum.Todo,
    created_time: created_time,
    updated_time: updated_time,
  },
  {
    id: 2,
    assignee_id: 2,
    title: 'title task 2',
    description: 'description task 2',
    priority: PriorityEnum.Medium,
    status: StatusEnum.Todo,
    created_time: created_time,
    updated_time: updated_time,
  },
];

const users: User[] = [
  {
    id: 1,
    name: 'test',
    limitTaskInDay: 2,
  },
  {
    id: 2,
    name: 'test',
    limitTaskInDay: 10,
  },
];

describe('TaskService', () => {
  let service: TaskService;
  let repository: Repository<Task>;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        TaskService,
        UserService,
        {
          provide: getRepositoryToken(Task),
          useValue: {
            find: jest.fn().mockResolvedValue(tasks),
            findOne: jest.fn().mockResolvedValue(tasks[0]),
            createQueryBuilder: jest.fn().mockReturnThis(),
            select: jest.fn().mockReturnThis(),
            where: jest.fn().mockReturnThis(),
            getRawOne: jest.fn().mockResolvedValue({ totalTaskInDay: 1 }),
            save: jest.fn().mockResolvedValue(tasks[0]),
            delete: jest.fn().mockResolvedValue(1),
          },
        },
        {
          provide: getRepositoryToken(User),
          useValue: {
            find: jest.fn().mockResolvedValue(users),
            findOne: jest.fn().mockResolvedValue(users[0]),
            save: jest.fn().mockResolvedValue(users[0]),
            delete: jest.fn().mockResolvedValue(1),
          },
        },
      ],
    }).compile();

    service = module.get<TaskService>(TaskService);
    repository = module.get<Repository<Task>>(getRepositoryToken(Task));
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });

  describe('findAll()', () => {
    it('should return an array of tasks', async () => {
      const data = await service.findAll({});
      expect(data).toEqual(tasks);
    });
  });

  describe('findOne()', () => {
    it('should get a single task', () => {
      const repoSpy = jest.spyOn(repository, 'findOne');
      const filter = { id: 1 };
      expect(service.findOne(filter)).resolves.toEqual(tasks[0]);
      expect(repoSpy).toBeCalledWith({ where: { ...filter } });
    });
  });

  describe('create()', () => {
    it('should return error user not found', async () => {
      try {
        await service.create({ ...createTaskDto, assignee_id: 5 });
      } catch (e) {
        expect(e).toBeInstanceOf(BadRequestException);
        expect(e.message).toBe('User not found');
      }
    });

    it('should return error (Total task in day must be less than limit task per user)', async () => {
      await repository.save({ ...new Task(), ...createTaskDto, id: 1 });
      await repository.save({ ...new Task(), ...createTaskDto, id: 2 });
      try {
        await service.create({ ...createTaskDto });
      } catch (e) {
        expect(e).toBeInstanceOf(BadRequestException);
        expect(e.message).toBe(
          'Total task in day must be less than limit task per user',
        );
      }
    });

    it('should successfully insert a task', () => {
      expect(service.create(createTaskDto)).resolves.toEqual(tasks[0]);
    });
  });

  describe('update()', () => {
    it('should return error user not found', async () => {
      try {
        await service.update(1, { ...createTaskDto, assignee_id: 5 });
      } catch (e) {
        expect(e).toBeInstanceOf(BadRequestException);
        expect(e.message).toBe('User not found');
      }
    });

    it('should return error task does not exist', async () => {
      try {
        await service.update(2, { ...createTaskDto });
      } catch (e) {
        expect(e).toBeInstanceOf(BadRequestException);
        expect(e.message).toBe('Task does not exist');
      }
    });

    it('should successfully insert a task', () => {
      expect(service.update(1, createTaskDto)).resolves.toEqual(tasks[0]);
    });
  });

  // describe('delete()', () => {
  //   it('should return 1', async () => {
  //     const removeSpy = jest.spyOn(repository, 'delete');
  //     const retVal = service.delete(1);
  //     expect(removeSpy).toBeCalledWith({ id: 1 });
  //     expect(retVal).resolves.toEqual(1);
  //   });
  // });
});
