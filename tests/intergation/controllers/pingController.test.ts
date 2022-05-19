import request from "supertest";

import app from "@/app";

import { version } from "package.json";

const server = request(app);

describe("/api/ping", () => {
   it("Should return status 200", async () => {
      const res = await server.get("/api/ping");
      expect(res.status).toBe(200);
      expect(res.body.message).toMatch(/OK: Success/g);
      expect(res.body.url).toMatch("/");
      expect(res.body.version).toMatch(version);
   });
});

afterAll(async () => {
   // avoid jest open handle error
   await new Promise((resolve) => {
      setTimeout(() => resolve(null), 500);
   });
});
