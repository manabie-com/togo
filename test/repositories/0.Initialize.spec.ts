import "mocha";
import { expect } from "chai";

import { createTypeormConn } from "../../utils/createTypeormConn";
import { getTypeormConn } from "../../utils/getTypeormConn";

let connection;

describe("Connection utility can establish and disconnect connections", () => {
  before(async () => {
    connection = await createTypeormConn();
  });

  after(async () => {
    connection = await getTypeormConn();
  });

  it("Should test open connection", async () => {
    expect(connection.isConnected).to.be.true;
  });

  it("Should test closed connection", async () => {
    await connection.close();
    expect(connection.isConnected).to.be.false;
  });
});
