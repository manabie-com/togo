const { expect, beforeAll, afterAll } = require("@jest/globals");
const supertest = require("supertest");
const mongoose = require("mongoose");

require("dotenv").config();


// Test for not premium user

// Token should be guaranteed new from not premium user login to maintain a valid test

const url = "http://localhost:24/api";
let token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjYyMGE5MzM3OGMwZjFmOGY4Nzg4ODkyOSIsImVtYWlsIjoic2FtcGxlQG0uY29tIiwiaWF0IjoxNjQ0ODYwMzA1fQ.UCKNOdA2SveyS600xMemklfyv1O_qbKg5soNBbi5b3g";

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