const db = require('../../src/infrastructure/models');

class DBQueries {
  async set(data) {
    try {
      const tables = data ? Object.keys(data) : null;

      await this.reset(tables);

      for (let id = 0; id < tables.length; id++) {
        const table = tables[id];
        const bulkCreateData = Object.keys(data[table]).map(
          id => data[table][id],
        );

        if (bulkCreateData.length) {
          await db[table].bulkCreate(bulkCreateData);
        }
      }
    } catch (error) {
      throw new Error(error);
    }
  }

  async reset(tables) {
    for (const table of tables) {
      await db[table].destroy({
        where: {},
        force: true,
        truncate : true,
      });
    }

    return true;
  }
}

module.exports = new DBQueries();
