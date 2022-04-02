import { Test } from '@nestjs/testing';
import { getRepositoryToken } from '@nestjs/typeorm';
import { UpdateToDoDto } from 'src/dto';
import { Task } from 'src/entities/task.entity';
import { ToDoList } from 'src/entities/toDoList.entity';
import { ToDoService } from '../services/todo.service';

describe('The todoService', () => {
  let todoService: ToDoService;
  let findOne: jest.Mock;
  let findAndCount: jest.Mock;
  let save: jest.Mock;
  let create: jest.Mock;
  let find: jest.Mock;
  beforeEach(async () => {
    findOne = jest.fn();
    findAndCount = jest.fn();
    save = jest.fn();
    create = jest.fn();
    find = jest.fn();
    const module = await Test.createTestingModule({
      providers: [
        ToDoService,
        {
          provide: getRepositoryToken(ToDoList),
          useValue: {
            findOne,
            findAndCount,
            save,
            create,
            find,
          },
        },
      ],
    }).compile();
    todoService = await module.get(ToDoService);
  });
  describe('when getting a todo by id', () => {
    describe('and the todo is matched', () => {
      let todo: ToDoList;
      beforeEach(() => {
        todo = new ToDoList({});
        findOne.mockReturnValue(Promise.resolve(todo));
      });
      it('should return the todo', async () => {
        const todoId = 1;
        const fetchedUser = await todoService.findOne(todoId);
        expect(fetchedUser).toEqual(todo);
      });
    });
  });
  describe('when getting todo list by task id', () => {
    const todoList = [new ToDoList({})];
    beforeEach(() => {
      find.mockReturnValue(Promise.resolve(todoList));
    });
    it('should return the todo list', async () => {
      const taskId = 1;
      const fetchedUser = await todoService.find(taskId);
      expect(fetchedUser).toEqual(todoList);
    });
  });
  describe('update todo', () => {
    describe('task id not matched', () => {
      const task = new Task({
        id: 1,
        title: 'task 1',
        desc: 'desc ...',
      });
      const todo = new ToDoList({ id: 1, title: 'todo 1', desc: '', task });
      beforeEach(() => {
        findOne.mockReturnValue(Promise.resolve(task));
      });
      it('throw error', async () => {
        const updateToDo = new UpdateToDoDto();
        const updateTaskId = 2;
        await expect(
          todoService.update(todo.id, updateTaskId, updateToDo),
        ).rejects.toThrow();
      });
    });
    describe('task id matched', () => {
      const task = new Task({
        id: 1,
        title: 'task 1',
        desc: 'desc ...',
      });
      const prevTodo = new ToDoList({
        id: 1,
        title: 'todo 1',
        desc: '',
        isDone: false,
        task,
      });
      const afterTodo = new ToDoList({
        id: 1,
        title: 'title',
        desc: 'desc',
        isDone: true,
        task,
      });
      beforeEach(() => {
        findOne.mockReturnValue(Promise.resolve(prevTodo));
        save.mockReturnValue(Promise.resolve(prevTodo));
      });
      it('should return todo', async () => {
        const updateToDo = new UpdateToDoDto();
        updateToDo.desc = 'desc';
        updateToDo.title = 'title';
        updateToDo.isDone = true;
        const updateTaskId = 1;
        const updatedTodo = await todoService.update(
          task.id,
          updateTaskId,
          updateToDo,
        );
        expect(updatedTodo).toEqual(afterTodo);
      });
    });
  });
});
