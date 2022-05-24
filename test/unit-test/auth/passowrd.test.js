const { hashPassword, comparePassword } = require('../../../src/services/services.user');

describe("[UNIT TEST]: PASSWORD TEST.", () => {
  // =============== CASE 01 ================
  it('Hash password then compare it', async () => {
    const hash = await hashPassword('hello_world');
    const verify = await comparePassword('hello_world', hash);
    expect(verify).toBe(true);
  });

  // =============== CASE 02 ================
  it('Hash password and compare with an other password', async () => {
    const hash = await hashPassword('hello_world');
    const verify = await comparePassword('hello_world_002', hash);
    expect(verify).toBe(false);
  });

  // =============== CASE 03 ================
  it('Compare two raw password not hasd.', async () => {
    const verify = await comparePassword('hello_world_002', 'hello_world_001');
    expect(verify).toBe(false);
  });
});