const jwt = require('jsonwebtoken');
const { signJWT, verifyToken } = require('../../../src/services/services.user');

describe("[UNIT TEST]: JWT TEST.", () => {
  // =============== CASE 01 ================
  it('Sign token and verify it', async () => {
    const token = await signJWT({
      userid: 'UserID',
      etcField: 'ETC'
    }, 1000);

    const verify = await verifyToken(token); // On verify success return a object.
    expect(!!verify).toBe(true);
  });

  // =============== CASE 02 ================
  it('Verify expired token.', async () => {
    const token = await signJWT({
      userid: 'UserID',
      etcField: 'ETC'
    }, 1);

    // Sleep
    await new Promise((solve) => {
      setTimeout(() => {
        solve(true);
      }, 1500);
    })

    const verify = await verifyToken(token); 
    expect(!!verify).toBe(false);
  });

  // =============== CASE 03 ================
  it('Verify self jwt sign token.', async () => {
    const token = await jwt.sign({
      userid: 'blabla'
    }, 'SEFL_SECRET');
    const verify = await verifyToken(token); 
    expect(!!verify).toBe(false);
  });
});