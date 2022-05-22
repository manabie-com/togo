const os = require("../../../libs/os");

describe("getOsEnv", () => {
  it("should throw if environment file is not configured", () => {
    expect(() => {
      os.getOsEnv("a");
    }).toThrowError(/a is not set/);
  });

  it("should not throw if environment file is configured", () => {
    process.env["a"] = jest.fn().mockReturnValue("a");

    expect(() => {
      os.getOsEnv("a");
    }).not.toThrow();
  });
});

describe("normalizePort", () => {
  it("should return passed port if this port is not a number", () => {
    const port = "a";
    const result = os.normalizePort(port);

    expect(result).toBe(port);
  });

  it("should return false when a negative number is passed", () => {
    const result = os.normalizePort(-1);

    expect(result).toBe(false);
  });

  it("should return a valid port when this port is a string number", () => {
    const port = "3000";
    const result = os.normalizePort(port);

    expect(result).toBe(parseInt(port, 10));
  });

  it("should return a valid port when this port is a number", () => {
    const port = 3000;
    const result = os.normalizePort(port);

    expect(result).toBe(parseInt(port, 10));
  });
});
