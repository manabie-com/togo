const { userService, authService } = require("../../../../apis/services");

describe("authService.login", () => {
  it("should throw if user's password is not match", () => {
    const user = null;
    userService.getUserByEmail = jest.fn().mockResolvedValue(user);

    expect(async () => {
      await authService.login("a@manabie.com", "password1");
    }).rejects.toThrow();
    expect(userService.getUserByEmail).toHaveBeenCalled();
  });

  it("should throw if user's password is not match", async () => {
    const user = {
      isPasswordMatch: jest.fn().mockResolvedValue(true),
    };

    userService.getUserByEmail = jest.fn().mockResolvedValue(user);
    const result = await authService.login("a@manabie.com", "password1");
    expect(result).toMatchObject(user);
    expect(userService.getUserByEmail).toHaveBeenCalled();
    expect(user.isPasswordMatch).toHaveBeenCalled();
  });
});

describe("authService.getTokenFromHeaders", () => {
  it("should throw if token is not contain Bearer", () => {
    const headers = {
      authorization: "a",
    };
    expect(() => {
      authService.getTokenFromHeaders(headers);
    }).toThrowError(/authenticate/);
  });

  it("should throw if token do not have 3 parts", () => {
    const headers = {
      authorization: "Bearer a",
    };
    expect(() => {
      authService.getTokenFromHeaders(headers);
    }).toThrowError(/Invalid/);
  });

  it("should return a valid token from headers", () => {
    const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjYyODUwNGFlMjA5YmFhZmViNjY5M2RhZiIsIm1heFRhc2siOjEwLCJpYXQiOjE2NTI5NTQ0Mzh9.No2n32pB8whz0M1yF8RHB14P0mrXZlqYB4QgV6uHL1s"
    const headers = {
      authorization: `Bearer ${token}`,
    };
    expect(authService.getTokenFromHeaders(headers)).toBe(token);
  });
});
