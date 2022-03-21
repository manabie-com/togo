const { expect, beforeAll, afterAll } = require("@jest/globals");
const supertest = require("supertest");
const mongoose = require("mongoose");

require("dotenv").config();

const url = "http://localhost:24/api/login";

// Connect to the DB
beforeAll(async () => {
  await mongoose.connect(process.env.MONGO_URI);
});

// Test if user already exist
it("should check if user does not exist", async () => {
    const response = await supertest(url)
      .post("/")
      .send({ email: "hello@m.com", password: "1234" });
    expect(response.statusCode).toEqual(200);
    expect(response.text).toEqual("{\"auth\":\"User does not exist\"}");
  });

// Test if user can login successfully
it("should check if user can login successfully", async () => {
    const response = await supertest(url)
      .post("/")
      .send({ email: "sample@m.com", password: "1234" });
    expect(response.statusCode).toEqual(200);
  });

// Disconnect the DB connection after all tests are run
afterAll(async () => {
  await mongoose.disconnect();
});