/* eslint-disable @typescript-eslint/no-misused-promises */

import {
  createStubInstance,
  expect,
  sinon,
  StubbedInstanceWithSinonAccessor,
} from '@loopback/testlab';
import {TodoController} from '../../../controllers';
import {Todo} from '../../../models';
import {TodoRepository} from '../../../repositories';
import {givenTodo} from '../../helpers';

describe('ProductController', () => {
  let todoRepo: StubbedInstanceWithSinonAccessor<TodoRepository>;

  let controller: TodoController;
  let aTodo: Todo;
  let aTodoWithId: Todo;
  let aChangedTodo: Todo;
  let aListOfTodos: Todo[];

  beforeEach(resetRepositories);

  describe('Todo', () => {
    it('creates a todo', async () => {
      const create = todoRepo.stubs.create;
      create.resolves(aTodo);
      const result = await controller.create(aTodo);
      aTodo.id = result.id;
      expect(result).to.eql(aTodo);
      sinon.assert.calledWith(create, aTodo);
    });
  });

  describe('find todo by id', () => {
    it('returns a todo if it exists', async () => {
      const findById = todoRepo.stubs.findById;
      findById.resolves(aTodoWithId);
      const result = await controller.findById(aTodoWithId.id as string);
      expect(result).to.eql(aTodoWithId);
      sinon.assert.calledWith(findById, aTodoWithId.id);
    });
  });

  describe('find todos', () => {
    it('returns multiple todo if they exist', async () => {
      const find = todoRepo.stubs.find;
      find.resolves(aListOfTodos);
      expect(await controller.find()).to.eql(aListOfTodos);
      sinon.assert.called(find);
    });

    it('returns empty list if no todos exist', async () => {
      const find = todoRepo.stubs.find;
      const expected: Todo[] = [];
      find.resolves(expected);
      expect(await controller.find()).to.eql(expected);
      sinon.assert.called(find);
    });
  });

  describe('replace product', () => {
    it('successfully replaces existing items', async () => {
      const replaceById = todoRepo.stubs.replaceById;
      replaceById.resolves();
      await controller.replaceById(aTodoWithId.id as string, aChangedTodo);
      sinon.assert.calledWith(replaceById, aTodoWithId.id, aChangedTodo);
    });
  });

  describe('update todo', () => {
    it('successfully updates existing items', async () => {
      const updateById = todoRepo.stubs.updateById;
      updateById.resolves();
      await controller.updateById(aTodoWithId.id as string, aChangedTodo);
      sinon.assert.calledWith(updateById, aTodoWithId.id, aChangedTodo);
    });
  });

  describe('delete todo', () => {
    it('successfully deletes existing items', async () => {
      const deleteById = todoRepo.stubs.deleteById;
      deleteById.resolves();
      await controller.deleteById(aTodoWithId.id as string);
      sinon.assert.calledWith(deleteById, aTodoWithId.id);
    });
  });

  function resetRepositories() {
    todoRepo = createStubInstance(TodoRepository);
    aTodo = givenTodo();
    aTodoWithId = givenTodo({
      id: '6217163b43cdc079a9d40554',
    });
    aListOfTodos = [
      aTodoWithId,
      givenTodo({
        id: '6217ad9e4c302422528a8609',
        title: 'New todo 02',
      }),
    ] as Todo[];
    aChangedTodo = givenTodo({
      id: aTodoWithId.id,
      title: 'New todo',
    });

    controller = new TodoController(todoRepo);
  }
});
