const { expect, beforeAll, afterAll } = require("@jest/globals");
const supertest = require("supertest");
const mongoose = require("mongoose");

require("dotenv").config();

const url = "http://localhost:24/api/register";

// Connect to the DB
beforeAll(async () => {
  await mongoose.connect(process.env.MONGO_URI);
});

// Test if new user is created
it("should create a new user", async () => {
    const response = await supertest(url)
      .post("/")
      .send({ name: "sample", email: "sample@m.com", password: "1234" });
    expect(response.statusCode).toEqual(200);
    expect(response.text).toEqual("{\"isAdded\":true}");
  });

// Test if new premium user is created
it("should create a new premium user", async () => {
    const response = await supertest(url)
      .post("/")
      .send({ name: "sample", email: "sample1@m.com", password: "1234", isPremium: true });
    expect(response.statusCode).toEqual(200);
    expect(response.text).toEqual("{\"isAdded\":true}");
  });

// Disconnect the DB connection after test run
afterAll(async () => {
  await mongoose.disconnect();
});