const { login } = require("..");

describe("Test login", () => {
  it("should get login successful with true username & password", () => {
    const username = "firstUser";
    const password = "example";
    const actual = login(username, password);
    expect(actual).toBeTruthy();
  });

  it("should get login failed with false username & password", () => {
    const username = "firstUser1";
    const password = "example3";
    const actual = login(username, password);
    expect(actual).toBeFalsy();
  });

  it("should get login failed with empty username & password", () => {
    const username = "";
    const password = "";
    const actual = login(username, password);
    expect(actual).toBeFalsy();
  });

  it("should get login failed with space key username & password", () => {
    const username = "       ";
    const password = "       ";
    const actual = login(username, password);
    expect(actual).toBeFalsy();
  });
});
