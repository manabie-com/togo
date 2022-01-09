
const koa = require("koa");
const chai = require('chai');
const chaiHttp = require("chai-http");
const mongoose = require("mongoose");
const { expect } = require("chai");
const { MongoMemoryServer } = require("mongodb-memory-server");
const app = require('../../src/app');

const serverAddress = 'http://localhost:9200';
chai.use(chaiHttp);

describe("[INTEGRATION TEST]: Home Info", () => {
    const mockDB = new MongoMemoryServer();
    beforeAll(async () => {
        await mockDB.start();
        const mongoUri = mockDB.getUri();
        await mongoose.connect(mongoUri, {
            useNewUrlParser: true,
            useUnifiedTopology: true,
        });
        app.listen(9200);
    });

    afterEach(async () => {
        const collections = mongoose.connection.collections;
        for (const key in collections) {
            const collection = collections[key];
            await collection.deleteMany({});
        }
    });

    afterAll(async () => {
        await mongoose.connection.dropDatabase();
        await mongoose.connection.close();
        await mockDB.stop();
    });

    // Test API Home page check server already accept connection
    describe("[HOME INFO] - GET /api/public/home", () => {
        it("Api successfull return 200 status and message includes REST API VERSION...`", (done) => {
            chai
                .request(serverAddress)
                .get("/api/public/home")
                .end((err, res) => {
                    expect(res.status).be.equal(200);
                    expect(res.body.success).equals(true);
                    expect(res.body.data.message).includes('REST API VERSION');
                    done();
                });
        });
    })
});
