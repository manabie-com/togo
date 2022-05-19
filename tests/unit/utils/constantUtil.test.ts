afterEach(() => {
   jest.resetModules();
   jest.resetAllMocks();
   jest.useFakeTimers("modern");
});

const oldEnv = process.env;

describe("PORT", () => {
   it("Pass none port should return port 3000", async () => {
      const { PORT } = await import("@/utils/constant.util");
      expect(PORT).toEqual("3000");
   });
   it("Pass port 5000 should return port 5000", async () => {
      process.env.PORT = "5000";
      const { PORT } = await import("@/utils/constant.util");

      expect(PORT).toEqual("5000");
   });
});

afterAll(() => {
   process.env = oldEnv;
});
