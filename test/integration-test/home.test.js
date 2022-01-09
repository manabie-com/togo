
const supertest = require('supertest');
const { expect } = require("chai");
const app = require('../../src/app');
const port = 9002;


describe("[INTEGRATION TEST]: Home Info", () => {
  describe("[GET] - /api/public/home", () => {
    it("Successfull receive message includes 'REST API VERSION'", (done) => {
      supertest(app.listen(port))
        .get("/api/public/home")
        .end((err, res) => {
          if (err) return done(err);
          expect(res.status).equals(200);
          expect(res.body?.data?.message).includes('REST API VERSION');
          done();
        });
    });
  });
});
