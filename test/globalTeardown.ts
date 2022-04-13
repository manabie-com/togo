import * as fs from 'fs';

module.exports = async () => {
  await (global as any).__APP__.close();
  await fs.unlinkSync('test.db');
  setTimeout(() => process.exit(), 1000);
  return Promise.resolve();
};
