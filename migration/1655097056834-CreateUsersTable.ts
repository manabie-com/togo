import { MigrationInterface, QueryRunner } from "typeorm";

export class CreateUsersTable1655097056834 implements MigrationInterface {
  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(
      "CREATE TABLE IF NOT EXISTS `users` (\
        `id` int (11) NOT NULL AUTO_INCREMENT, \
        `user_id` int (11), \
        `api_key` varchar(255) NOT NULL, \
        `daily_maximum_tasks` tinyint NOT NULL DEFAULT 3, \
        `is_active` tinyint(1) NOT NULL DEFAULT 1,\
        PRIMARY KEY (`id`) \
      )"
    );
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query("DROP TABLE IF EXISTS `users`;");
  }
}
