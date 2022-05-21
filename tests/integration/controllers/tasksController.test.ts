import request from "supertest";

import app from "@/app";
import UsersModel from "@/models/users.model";
import TasksModel from "@/models/tasks.model";
import { isAfter, isEqual } from "date-fns";

// eslint-disable-next-line @typescript-eslint/no-var-requires
const mockingoose = require("mockingoose");

const server = request(app);

const stubUsers = [
   {
      _id: "628695e4abfb7fd6ca76ff99",
      limit: 1,
      username: "testing",
   },
];

const stubSTask: { name: string; user: string; createdAt: Date }[] = [];

const finderUsersMockOne = (query: { getQuery: () => any }) => {
   const { id } = query.getQuery();
   return stubUsers.filter((data) => data._id === id)[0];
};

const finderTasksMockOne = (query: { getQuery: () => any }) => {
   const { createdAt } = query.getQuery();
   console.log("===========finderTasksMockOne=============\n");
   console.log("query.getQuery(): ", query);
   return stubSTask.filter(
      (data) => isEqual(data.createdAt, createdAt.$gte) || isAfter(data.createdAt, createdAt.$gte)
   );
};

mockingoose(UsersModel).toReturn(finderUsersMockOne, "findOne");

mockingoose(TasksModel).toReturn(finderTasksMockOne, "find");
jest.spyOn(TasksModel, "create").mockImplementation((value) => {
   const newData = {
      ...(value as unknown as { name: string; user: string }),
      createdAt: new Date(),
   };
   stubSTask.push(newData);
   return Promise.resolve();
});
jest.setTimeout(30000);

describe("Create new task", () => {
   it("With no body and no params userID should return status 400", async () => {
      const res = await server.post("/api/task/");
      expect(res.status).toBe(400);
      expect(res.body.message).toMatch(/Bad Request: Please check parameter/g);
   });

   it("With no body should return status 400", async () => {
      const res = await server.post("/api/task/628695e4abfb7fd6ca76ff99");
      expect(res.status).toBe(400);
      expect(res.body.message).toMatch(/Bad Request: Please check parameter/g);
   });

   it("With no params userID should return status 400", async () => {
      const res = await server.post("/api/task/").send({ name: "testing" });
      expect(res.status).toBe(400);
      expect(res.body.message).toMatch(/Bad Request: Please check parameter/g);
   });

   it("With name in body empty should return status 400", async () => {
      const res = await server.post("/api/task/628695e4abfb7fd6ca76ff99").send({ name: "" });
      expect(res.status).toBe(400);
      expect(res.body.message).toMatch(/Bad Request: Please check parameter/g);
   });

   it("With userID is wrong should return status 404", async () => {
      const res = await server.post("/api/task/wrong_userID").send({ name: "testing task" });
      expect(res.status).toBe(404);
      expect(res.body.message).toMatch(/Not found: There is no resource behind the URI/g);
   });

   it("Normal case in the body should return status 201", async () => {
      const res = await server
         .post("/api/task/628695e4abfb7fd6ca76ff99")
         .send({ name: "testing normal case" });

      expect(res.status).toBe(201);
      expect(res.body.message).toMatch(/OK: New resource has been created/g);
   });

   it("Normal case but out of limit users should return status 201", async () => {
      const res = await server
         .post("/api/task/628695e4abfb7fd6ca76ff99")
         .send({ name: "testing normal case but out of limit" });

      expect(res.status).toBe(400);
      expect(res.body.message).toMatch(
         /Bad request: This user is out of the limit in order to create a new task./g
      );
   });
});

afterAll(async () => {
   // avoid jest open handle error
   await new Promise((resolve) => {
      setTimeout(() => resolve(null), 500);
   });
});
