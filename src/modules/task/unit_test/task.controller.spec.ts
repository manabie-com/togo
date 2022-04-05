import { Test, TestingModule } from '@nestjs/testing';
import { CreateTaskDTO } from './../context';
import { TaskController } from './../task.controller';
import { TaskService } from './../task.service';
import { PriorityEnum, StatusEnum } from './../../../common';
import { Task } from '../task.entity';
import { BadRequestException } from '@nestjs/common';

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

const created_time = new Date();
const updated_time = new Date();

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

describe('TaskController', () => {
  let taskController: TaskController;
  let taskService: TaskService;

  beforeEach(async () => {
    const app: TestingModule = await Test.createTestingModule({
      controllers: [TaskController],
      providers: [
        TaskService,
        {
          provide: TaskService,
          useValue: {
            findAll: jest.fn().mockResolvedValue(tasks),
            findOne: jest
              .fn()
              .mockImplementation((id: string) => Promise.resolve(tasks[0])),
            create: jest.fn().mockImplementation((context: CreateTaskDTO) =>
              Promise.resolve({
                id: 1,
                created_time: created_time,
                updated_time: updated_time,
                ...context,
              }),
            ),
            update: jest
              .fn()
              .mockImplementation((id: number, context: CreateTaskDTO) =>
                Promise.resolve({
                  id: 1,
                  created_time: created_time,
                  updated_time: updated_time,
                  ...context,
                }),
              ),
            delete: jest
              .fn()
              .mockImplementation((id: number) => Promise.resolve(1)),
          },
        },
      ],
    }).compile();

    taskController = app.get<TaskController>(TaskController);
    taskService = app.get<TaskService>(TaskService);
  });

  it('should be defined', () => {
    expect(taskController).toBeDefined();
  });

  describe('findAll()', () => {
    it('should find all task ', () => {
      taskController.findAll({});
      expect(taskService.findAll).toHaveBeenCalled();
    });
  });

  describe('findOne()', () => {
    it('should find a task', () => {
      taskController.findOne(1);
      expect(taskService.findOne).toHaveBeenCalled();
    });
  });

  describe('create()', () => {
    it('should create a task', () => {
      taskController.create(createTaskDto);
      expect(taskController.create(createTaskDto)).resolves.toEqual({
        status: 'success',
        data: {
          id: 1,
          created_time: created_time,
          updated_time: updated_time,
          ...createTaskDto,
        },
      });
      expect(taskService.create).toHaveBeenCalledWith(createTaskDto);
    });

    it('should return error User not found', async () => {
      taskController.create(createTaskDto);
      try {
        await taskController.create({...createTaskDto, assignee_id: 2});
      } catch (e) {
        expect(e).toBeInstanceOf(BadRequestException);
      }
    });
  });

  describe('update()', () => {
    it('should update a task', () => {
      const task_id = 1;
      taskController.update(task_id, updateTaskDto);
      expect(taskController.update(task_id, updateTaskDto)).resolves.toEqual({
        status: 'success',
        data: {
          id: task_id,
          created_time: created_time,
          updated_time: updated_time,
          ...updateTaskDto,
        },
      });
      expect(taskService.update).toHaveBeenCalledWith(1, updateTaskDto);
    });

    it('should return error User not found', async () => {
      taskController.update(1, createTaskDto);
      try {
        await taskController.update(1, {...createTaskDto, assignee_id: 2});
      } catch (e) {
        expect(e).toBeInstanceOf(BadRequestException);
      }
    });
  });

  describe('delete()', () => {
    it('should delete the task', () => {
      const task_id = 1;
      taskController.delete(task_id);
      expect(taskController.delete(task_id)).resolves.toEqual({
        status: 'success',
        data: 1,
      });
      expect(taskService.delete).toHaveBeenCalledWith(task_id);
    });
  });
});
