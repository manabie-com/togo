const validAuthorized = require("..");
const MOCK_TOKEN = require("../../constants");

describe("Test valid authorized", () => {
  it("should validAuthorized successful with true token", () => {
    const actual = validAuthorized(MOCK_TOKEN);
    expect(actual).toBeTruthy();
  });

  it("should validAuthorized failed with false token", () => {
    const mockToken = "xxxxxxx.xxx.xxx";
    const actual = validAuthorized(mockToken);
    expect(actual).toBeFalsy();
  });

  it("should validAuthorized failed with empty token", () => {
    const mockToken = "";
    const actual = validAuthorized(mockToken);
    expect(actual).toBeFalsy();
  });

  it("should validAuthorized failed with space key token", () => {
    const mockToken = "            ";
    const actual = validAuthorized(mockToken);
    expect(actual).toBeFalsy();
  });
});
