import "mocha";
import { expect } from "chai";
import db from "../../providers/typeorm";
import { ITodoRepository } from "../../core/repository/ITodoRepository";
import { getConnection } from "typeorm";
import { Todo } from "../../entity/Todo";
import { TodoRepositoryFactory } from "../../factories/TodoRepositoryFactory";

describe("Todo Repository is able to save data", () => {
  let todoRepository: ITodoRepository;

  before(async () => {
    await db.initialize();
    todoRepository = await TodoRepositoryFactory.createInstance();
  });

  afterEach(async () => {
    const repository = getConnection().getRepository(Todo);
    repository.clear();
  });

  it("Should return a todo data after saving", async () => {
    const params = {
      task: "Never gonna give you up, Never gonna let you down",
      userId: 10172512,
    };

    const data = await todoRepository.saveTodo(params);

    expect(data.task).to.equal(params.task);
    expect(data.userId).to.equal(params.userId);
  });
});
