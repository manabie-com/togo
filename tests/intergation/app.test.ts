import request from "supertest";

import app from "@/app";

// eslint-disable-next-line @typescript-eslint/no-var-requires
const mockingoose = require("mockingoose");

const server = request(app);
jest.setTimeout(60000);

beforeEach(() => {
   jest.resetModules(); // Most important - it clears the cache
});

describe("The end point should not found", () => {
   it("End point '/' should return Not Found", async () => {
      const res = await server.get("/");
      expect(res.status).toBe(404);
      expect(res.body.message).toMatch(/Not found: There is no resource behind the URI/g);
   });
   it("End point '/api' should return Not Found", async () => {
      const res = await server.get("/api");
      expect(res.status).toBe(404);
      expect(res.body.message).toMatch(/Not found: There is no resource behind the URI/g);
   });
   it("End point '/api' should return Not Found", async () => {
      const res = await server.get("/api");
      expect(res.status).toBe(404);
      expect(res.body.message).toMatch(/Not found: There is no resource behind the URI/g);
   });
});

afterAll(async () => {
   mockingoose.resetAll();
   jest.clearAllMocks();
   // avoid jest open handle error
   await new Promise((resolve) => {
      setTimeout(() => resolve(null), 500);
   });
});
