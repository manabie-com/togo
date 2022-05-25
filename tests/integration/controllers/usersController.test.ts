import request from "supertest";

import app from "@/app";
import UsersModel from "@/models/users.model";

// eslint-disable-next-line @typescript-eslint/no-var-requires
const mockingoose = require("mockingoose");

const server = request(app);

const stubUsers = [
   {
      _id: "628695e4abfb7fd6ca76ff99",
      limit: 10,
      username: "testing",
   },
];
const finderUsersMockOne = (query: { getQuery: () => any }) => {
   const { username } = query.getQuery();

   return stubUsers.filter((data) => data.username === username)[0];
};

mockingoose(UsersModel).toReturn(finderUsersMockOne, "findOne");

describe("Create new user", () => {
   it("With no body should return status 400", async () => {
      const res = await server.post("/api/user/");
      expect(res.status).toBe(400);
      expect(res.body.message).toMatch(/Bad Request: Please check parameter/g);
   });

   it("With username in body empty should return status 400", async () => {
      const res = await server.post("/api/user/").send({ username: "" });
      expect(res.status).toBe(400);
      expect(res.body.message).toMatch(/Bad Request: Please check parameter/g);
   });

   it("With username in body is already have in DB should return status 400", async () => {
      const res = await server.post("/api/user/").send({ username: "testing" });
      expect(res.status).toBe(400);
      expect(res.body.message).toMatch(/Username is invalid or already taken/g);
   });

   it("Normal case with no limit in the body should return status 201", async () => {
      const res = await server.post("/api/user/").send({ username: "testing_1" });
      expect(res.status).toBe(201);
      expect(res.body.message).toMatch(/OK: New resource has been created/g);
   });

   it("Normal case in the body should return status 201", async () => {
      const res = await server.post("/api/user/").send({ limit: 100, username: "testing_1" });
      expect(res.status).toBe(201);
      expect(res.body.message).toMatch(/OK: New resource has been created/g);
   });
});

describe("Get profile", () => {
   it("GET method with no params should return status 400", async () => {
      const res = await server.get("/api/user/");
      expect(res.status).toBe(400);
      expect(res.body.message).toMatch(/Bad Request: Please check parameter/g);
   });

   it("Not have username in DB should return status 404", async () => {
      const res = await server.get("/api/user/not_have_username");
      expect(res.status).toBe(404);
      expect(res.body.message).toMatch(/Not found: There is no resource behind the URI/g);
   });

   it("Normal case should return status 200", async () => {
      const res = await server.get("/api/user/testing");
      expect(res.status).toBe(200);
      expect(res.body.message).toMatch(/OK: Success/g);
      expect(res.body.data.limit).toBe(10);
      expect(res.body.data.username).toBe("testing");
   });
});

afterAll(async () => {
   // avoid jest open handle error
   await new Promise((resolve) => {
      setTimeout(() => resolve(null), 500);
   });
});
