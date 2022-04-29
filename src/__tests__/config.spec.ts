describe('config', () => {
  afterEach(() => {
    // reset test env
    process.env.NODE_ENV = 'test';
  });
  it('should not throw error on empty NODE_ENV', () => {
    process.env.NODE_ENV = '';
    const config = require('../config');
    expect(config).toBeDefined();
  });
});
