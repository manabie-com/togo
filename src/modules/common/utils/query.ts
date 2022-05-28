import { Connection, ObjectLiteral } from 'typeorm';
import { safeKey } from '.';

export function orderByBuilder(
  sort: string,
  alias?: string,
  mapField?: ObjectLiteral,
): Record<string, 'ASC' | 'DESC'>[] {
  const sorts = sort.split(',');
  const ASC = 'ASC';
  const DESC = 'DESC';

  return sorts
    .filter((val) => !!val)
    .map((sort) => {
      const orderBy = {};
      const arrOrder = sort.split('|').filter((val) => !!val);

      if (arrOrder.length != 2) {
        return '';
      }

      let field = arrOrder[0];

      if (mapField && mapField[safeKey(field)] !== undefined) {
        field = mapField[safeKey(field)];
      } else if (field.indexOf('.') == -1) {
        field = `${alias}.${arrOrder[0]}`;
      }

      let order = arrOrder[1].toUpperCase();

      if (![ASC, DESC].includes(order)) {
        order = ASC;
      }

      orderBy[safeKey(field)] = order;

      return orderBy;
    });
}

export async function cleanTable(connection: Connection, tableName: string): Promise<void> {
  const queryRunner = connection.createQueryRunner();

  await queryRunner.query(`DELETE FROM [${tableName}]`);
}

const tableForeignKeys: Record<string, any>[] = [];

export async function dropForeignKeys(connection: Connection): Promise<void> {
  const queryRunner = connection.createQueryRunner();

  const getTableNameQueryString = `
  SELECT TABLE_NAME
  FROM INFORMATION_SCHEMA.TABLES
  WHERE TABLE_TYPE = 'BASE TABLE' AND TABLE_CATALOG='${process.env.DB_NAME}' AND TABLE_NAME != 'migrations'
  `;

  const tableNames = (await connection.query(getTableNameQueryString)).map((x) => x.TABLE_NAME);

  const { length } = tableNames;

  for (let i = 0; i < length; i++) {
    const tableName = tableNames[safeKey(i)];
    const table = await queryRunner.getTable(tableName);

    const foreignKeys = Object.assign([], table.foreignKeys);

    if (foreignKeys && foreignKeys.length > 0) {
      await queryRunner.dropForeignKeys(tableName, foreignKeys);
      tableForeignKeys.push({ tableName, foreignKeys: foreignKeys });
    }
  }
}

export async function createForeignKeys(connection: Connection): Promise<void> {
  const queryRunner = connection.createQueryRunner();

  const { length } = tableForeignKeys;

  for (let i = length - 1; i >= 0; i--) {
    const tableForeignKey = tableForeignKeys[safeKey(i)];
    const tableName = tableForeignKey.tableName;

    const foreignKeys = tableForeignKey.foreignKeys;

    if (foreignKeys.length > 0) {
      await queryRunner.createForeignKeys(tableName, foreignKeys);
    }
  }
}
