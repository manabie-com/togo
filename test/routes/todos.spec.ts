import server from "../../server";
import supertest from "supertest";

const request = supertest(server);

describe("Todos APIs", () => {
  describe("POST /api/todos", () => {
    it("Should POST successfully", async () => {
      const response = await request
        .post("/api/todos")
        .send({ task: "Foo", userId: 10172512 })
        .expect(200);
    });
  });
});
