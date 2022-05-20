import {MigrationInterface, QueryRunner, Table, TableForeignKey} from "typeorm";

export class createTaskTable1652977055891 implements MigrationInterface {

    public async up(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.createTable(
            new Table({
              name: "tasks",
              columns: [
                {
                  name: "id",
                  type: "int",
                  isPrimary: true,
                  isGenerated: true,
                  generationStrategy: "increment"
                },
                {
                    name: "content",
                    type: "varchar",
                },
                {
                    name: "userId",
                    type: "int",
                },
                {
                  name: "createdAt",
                  type: "timestamp",
                  default: "now()",
                },
                {
                  name: "updatedAt",
                  type: "timestamp",
                  default: "now()",
                },
              ],
            }),
            true
          );

		const foreignKey = new TableForeignKey({
			columnNames: ["userId"],
			referencedColumnNames: ["id"],
			referencedTableName: "users",
			onDelete: "CASCADE"
		});
		await queryRunner.createForeignKey("tasks", foreignKey);
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
      await queryRunner.dropTable('tasks');
    }

}
