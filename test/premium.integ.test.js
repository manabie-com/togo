const { expect, beforeAll, afterAll } = require("@jest/globals");
const supertest = require("supertest");
const mongoose = require("mongoose");

require("dotenv").config();


// Test for premium user

// Token should be guaranteed new from premium user login to maintain a valid test

const url = "http://localhost:24/api";
let token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjYyMGE5MzM3OGMwZjFmOGY4Nzg4ODkyYiIsImVtYWlsIjoic2FtcGxlMUBtLmNvbSIsImlhdCI6MTY0NDg2MDI2Nn0.QmKDd02_6LoUlzvuvxBAvmI8N-Ke4vSBdFKcyFvZg4o";

// Connect to the DB
beforeAll(async () => {
  await mongoose.connect(process.env.MONGO_URI);
});

// Test if route will fail for unauthenticated GET request.
it("should return '{\"auth\":\"failed\"}'", async () => {
  const response = await supertest(url).get("/");
  expect(response.statusCode).toEqual(200);
  expect(response.text).toEqual("{\"auth\":\"failed\"}");
});

// Test if route will fail for unauthenticated POST request.
it("should return '{\"auth\":\"failed\"}", async () => {
  const response = await supertest(url)
    .post("/")
    .send({ description: "This is a test" });
  expect(response.statusCode).toEqual(200);
  expect(response.text).toEqual("{\"auth\":\"failed\"}");
});

// First post request
it("should create a post 1", async () => {
  const response = await supertest(url)
    .post("/")
    .set({ Authorization: 'bearer ' + token, 'Content-Type': 'application/json' })
    .send({ description: "Test 1" });
  expect(response.statusCode).toEqual(200);
  expect(response.text).toEqual("{\"isAdded\":true}");
});

// Second post request
it("should create a post 2", async () => {
    const response = await supertest(url)
      .post("/")
      .set({ Authorization: 'bearer ' + token, 'Content-Type': 'application/json' })
      .send({ description: "Test 2" });
    expect(response.statusCode).toEqual(200);
    expect(response.text).toEqual("{\"isAdded\":true}");
  });
  
// Third post request
  it("should create a post 3", async () => {
    const response = await supertest(url)
      .post("/")
      .set({ Authorization: 'bearer ' + token, 'Content-Type': 'application/json' })
      .send({ description: "Test 3" });
    expect(response.statusCode).toEqual(200);
    expect(response.text).toEqual("{\"isAdded\":true}");
  });

// Fourth post request
it("should create a post 4", async () => {
  const response = await supertest(url)
    .post("/")
    .set({ Authorization: 'bearer ' + token, 'Content-Type': 'application/json' })
    .send({ description: "Test 4" });
  expect(response.statusCode).toEqual(200);
  expect(response.text).toEqual("{\"isAdded\":true}");
});

// Fifth post request
it("should create a post 5", async () => {
  const response = await supertest(url)
    .post("/")
    .set({ Authorization: 'bearer ' + token, 'Content-Type': 'application/json' })
    .send({ description: "Test 5" });
  expect(response.statusCode).toEqual(200);
  expect(response.text).toEqual("{\"isAdded\":true}");
});
  
// Test if max limit will be triggered
it("should return 'Daily limit reached", async () => {
  const response = await supertest(url)
    .post("/")
    .set({ Authorization: 'bearer ' + token, 'Content-Type': 'application/json' })
    .send({ description: "Max limit test" });
  expect(response.statusCode).toEqual(429);
  expect(response.text).toEqual("Maximum daily limit reached");
});

// Test for get all task request
it("should get all task of the user", async () => {
  const response = await supertest(url)
    .get("/")
    .set({ Authorization: 'bearer ' + token, 'Content-Type': 'application/json' })
  expect(response.statusCode).toEqual(200);
});

// Disconnect the DB connection after test run
afterAll(async () => {
  await mongoose.disconnect();
});