import {MigrationInterface, QueryRunner, Table} from "typeorm";

export class createUserTable1652976701370 implements MigrationInterface {

    public async up(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.createTable(
            new Table({
              name: "users",
              columns: [
                {
                  name: "id",
                  type: "int",
                  isPrimary: true,
                  isGenerated: true,
                  generationStrategy: "increment"
                },
                {
                    name: "email",
                    type: "varchar",
                },
                {
                    name: "password",
                    type: "varchar",
                },
				        {
                    name: "maxTasks",
                    type: "int",
                    default: 5
                },
                {
                  name: "role",
                  type: "enum",
                  enumName: "UserRole",
                  enum: ['admin', 'member'],
                  default: "'member'",
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
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
		await queryRunner.dropTable('users');
	}

}
