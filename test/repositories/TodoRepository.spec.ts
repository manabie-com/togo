import "mocha";
import { expect } from "chai";
import { ITodoRepository } from "../../core/repository/ITodoRepository";
import { createConnection, getConnection } from "typeorm";
import { Todo } from "../../entity/Todo";
import { TodoRepository } from "../../repositories/TodoRepository";
import moment from "moment";

const testConnection = "testConnection";

describe("Todo Repository is able to save and fetch data", () => {
  let todoRepository: ITodoRepository;
  const creationDate = moment().format("YYYY-MM-DD HH:mm:ss");

  beforeEach(async () => {
    const connection = await createConnection({
      type: "sqlite",
      database: ":memory:",
      dropSchema: true,
      entities: [Todo],
      synchronize: true,
      logging: false,
      name: testConnection,
    });

    todoRepository = connection.getCustomRepository(TodoRepository);
  });

  afterEach(async () => {
    await getConnection(testConnection).close();
  });

  it("Should return a todo data after saving", async () => {
    const params = {
      task: "Never gonna give you up, Never gonna let you down",
      userId: 10172512,
      creationDate,
    };

    const data = await todoRepository.saveTodo(params);

    expect(data.task).to.equal(params.task);
    expect(data.userId).to.equal(params.userId);
  });

  it("Should list current day tasks based on User ID", async () => {
    const userId = 10172512;
    const params = [
      {
        task: "Never gonna let you down",
        userId,
        creationDate,
      },
      {
        task: "Never gonna run around and desert you",
        userId,
        creationDate,
      },
    ];

    for (let item of params) {
      await todoRepository.saveTodo(item);
    }

    // const data = await todoRepository.getCurrentTasksByUserId(userId);

    // expect(data).to.equal(params.length);
  });
});
