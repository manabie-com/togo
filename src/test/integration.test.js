const { expect, beforeAll, afterAll } = require("@jest/globals");
const supertest = require("supertest");
const mongoose = require("mongoose");
const TaskMongo = require("../model/taskMongo");
const UserMongo = require("../model/userMongo");

require("dotenv").config();

const url = "http://localhost:5000/task";

// Delete the existing test account and tasks first
beforeAll(async () => {
  await mongoose.connect(process.env.MONGO_URL_LOCAL);
  await TaskMongo.deleteMany({ user_name: "test1" });
  await UserMongo.deleteOne({ user_name: "test1" });
  await UserMongo.insertMany({ user_name: "test1", task_daily_limit: 3 });
});

it("should return 'invalid route'", async () => {
  const response = await supertest(url).get("/");
  expect(response.statusCode).toEqual(500);
  expect(response.text).toEqual("Invalid route");
});

it("should return 'user does not exist'", async () => {
  const response = await supertest(url)
    .post("/")
    .send({ user_name: "test2", task_name: "task 1" });
  expect(response.statusCode).toEqual(404);
  expect(response.text).toEqual("User does not exist");
});

it("should create a task - 1", async () => {
  const response = await supertest(url)
    .post("/")
    .send({ user_name: "test1", task_name: "task 1" });
  expect(response.statusCode).toEqual(201);
  expect(response.text).toEqual("Task created");
});

it("should create a task - 2", async () => {
  const response = await supertest(url)
    .post("/")
    .send({ user_name: "test1", task_name: "task 2" });
  expect(response.statusCode).toEqual(201);
  expect(response.text).toEqual("Task created");
});

it("should create a task - 3", async () => {
  const response = await supertest(url)
    .post("/")
    .send({ user_name: "test1", task_name: "task 3" });
  expect(response.statusCode).toEqual(201);
  expect(response.text).toEqual("Task created");
});

it("should return 'Daily limit reached", async () => {
  const response = await supertest(url)
    .post("/")
    .send({ user_name: "test1", task_name: "task 4" });
  expect(response.statusCode).toEqual(403);
  expect(response.text).toEqual("Daily limit reached.");
});

// Disconnect the DB connection after all tests are run
afterAll(async () => {
  await mongoose.disconnect();
});
