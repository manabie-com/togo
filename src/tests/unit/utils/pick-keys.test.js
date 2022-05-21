const pickKeys = require("../../../utils/pick-keys");

describe("pickKeys", () => {
  it("should be return an object composed of the picked object properties", () => {
    const object = { name: "Ha", age: 10, job: "developer", address: "Hanoi" };
    const result = pickKeys(object, ["name", "job", "role"]);
    expect(result).toEqual({ name: "Ha", job: "developer" });
  });
});
