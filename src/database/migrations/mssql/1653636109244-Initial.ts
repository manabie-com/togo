import { MigrationInterface, QueryRunner } from 'typeorm';

export class Initial1653636109244 implements MigrationInterface {
  name = 'Initial1653636109244';

  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(
      `CREATE TABLE "black_listed_token" ("id" uniqueidentifier NOT NULL CONSTRAINT "DF_d86ed9ade5ece3723348d562d42" DEFAULT NEWSEQUENTIALID(), "created_by" int, "updated_by" int, "created_at" datetime2 NOT NULL CONSTRAINT "DF_019d277bbc91060a816d94ffa2e" DEFAULT getdate(), "updated_at" datetime2 NOT NULL CONSTRAINT "DF_856aa87b893802be6b1fb2e599b" DEFAULT getdate(), "deleted_at" datetime2, "user_id" nvarchar(255) NOT NULL, "token" nvarchar(3000) NOT NULL, CONSTRAINT "PK_d86ed9ade5ece3723348d562d42" PRIMARY KEY ("id"))`,
    );
    await queryRunner.query(`CREATE INDEX "IDX_f4e97792cfe213d7e545ef6469" ON "black_listed_token" ("created_by") `);
    await queryRunner.query(`CREATE INDEX "IDX_2cc9f45b3a6e9fc786ba98a41b" ON "black_listed_token" ("updated_by") `);
    await queryRunner.query(`CREATE INDEX "IDX_019d277bbc91060a816d94ffa2" ON "black_listed_token" ("created_at") `);
    await queryRunner.query(`CREATE INDEX "IDX_856aa87b893802be6b1fb2e599" ON "black_listed_token" ("updated_at") `);
    await queryRunner.query(`CREATE INDEX "IDX_1c846b375f568d13277214b05c" ON "black_listed_token" ("deleted_at") `);
    await queryRunner.query(`CREATE INDEX "IDX_8ef9962491416367a8787a1f91" ON "black_listed_token" ("user_id") `);
    await queryRunner.query(
      `CREATE TABLE "permissions" ("id" uniqueidentifier NOT NULL CONSTRAINT "DF_920331560282b8bd21bb02290df" DEFAULT NEWSEQUENTIALID(), "created_by" int, "updated_by" int, "created_at" datetime2 NOT NULL CONSTRAINT "DF_337088ff813c697c964f49f58fd" DEFAULT getdate(), "updated_at" datetime2 NOT NULL CONSTRAINT "DF_b01aa3668b2d129f4bf24f002cf" DEFAULT getdate(), "deleted_at" datetime2, "name" nvarchar(100) NOT NULL, "resource" nvarchar(100) NOT NULL, "possession" nvarchar(255) CONSTRAINT CHK_c6a7315d66653b4947e4fc4212_ENUM CHECK(possession IN ('own','any')) NOT NULL CONSTRAINT "DF_7009b13149cc92b91ed410066a3" DEFAULT 'own', "action" nvarchar(100) NOT NULL CONSTRAINT "DF_1c1e0637ecf1f6401beb9a68abe" DEFAULT 'list', "status" nvarchar(255) CONSTRAINT CHK_d5b4c7e9353e87508d337c9531_ENUM CHECK(status IN ('active','inactive')) NOT NULL CONSTRAINT "DF_bbf6febffd0f64508b38a2cd514" DEFAULT 'active', CONSTRAINT "UQ_48ce552495d14eae9b187bb6716" UNIQUE ("name"), CONSTRAINT "PK_920331560282b8bd21bb02290df" PRIMARY KEY ("id"))`,
    );
    await queryRunner.query(`CREATE INDEX "IDX_c398f7100db3e0d9b6a6cd6bea" ON "permissions" ("created_by") `);
    await queryRunner.query(`CREATE INDEX "IDX_58fae278276b7c2c6dde2bc19a" ON "permissions" ("updated_by") `);
    await queryRunner.query(`CREATE INDEX "IDX_337088ff813c697c964f49f58f" ON "permissions" ("created_at") `);
    await queryRunner.query(`CREATE INDEX "IDX_b01aa3668b2d129f4bf24f002c" ON "permissions" ("updated_at") `);
    await queryRunner.query(`CREATE INDEX "IDX_1ea42cae477fc1dc619a5cd280" ON "permissions" ("deleted_at") `);
    await queryRunner.query(`CREATE INDEX "UIDX_Permission_Name" ON "permissions" ("name") `);
    await queryRunner.query(`CREATE INDEX "UIDX_Permission_Resource" ON "permissions" ("resource") `);
    await queryRunner.query(`CREATE INDEX "IDX_Permission_possession" ON "permissions" ("possession") `);
    await queryRunner.query(`CREATE INDEX "IDX_Permission_action" ON "permissions" ("action") `);
    await queryRunner.query(`CREATE INDEX "IDX_Permission_Status" ON "permissions" ("status") `);
    await queryRunner.query(
      `CREATE TABLE "roles" ("id" uniqueidentifier NOT NULL CONSTRAINT "DF_c1433d71a4838793a49dcad46ab" DEFAULT NEWSEQUENTIALID(), "created_by" int, "updated_by" int, "created_at" datetime2 NOT NULL CONSTRAINT "DF_e5a52fc6f7a8dae64f645b09146" DEFAULT getdate(), "updated_at" datetime2 NOT NULL CONSTRAINT "DF_8651ace7d160b9cf59cd0e0e2ef" DEFAULT getdate(), "deleted_at" datetime2, "name" nvarchar(100), "description" nvarchar(300), CONSTRAINT "UQ_648e3f5447f725579d7d4ffdfb7" UNIQUE ("name"), CONSTRAINT "PK_c1433d71a4838793a49dcad46ab" PRIMARY KEY ("id"))`,
    );
    await queryRunner.query(`CREATE INDEX "IDX_4a39f3095781cdd9d6061afaae" ON "roles" ("created_by") `);
    await queryRunner.query(`CREATE INDEX "IDX_747b580d73db0ad78963d78b07" ON "roles" ("updated_by") `);
    await queryRunner.query(`CREATE INDEX "IDX_e5a52fc6f7a8dae64f645b0914" ON "roles" ("created_at") `);
    await queryRunner.query(`CREATE INDEX "IDX_8651ace7d160b9cf59cd0e0e2e" ON "roles" ("updated_at") `);
    await queryRunner.query(`CREATE INDEX "IDX_7fd0c79dc4e6083ddea850ac38" ON "roles" ("deleted_at") `);
    await queryRunner.query(`CREATE INDEX "UIDX_Role_name" ON "roles" ("name") `);
    await queryRunner.query(
      `CREATE TABLE "users" ("id" uniqueidentifier NOT NULL CONSTRAINT "DF_a3ffb1c0c8416b9fc6f907b7433" DEFAULT NEWSEQUENTIALID(), "created_by" int, "updated_by" int, "created_at" datetime2 NOT NULL CONSTRAINT "DF_c9b5b525a96ddc2c5647d7f7fa5" DEFAULT getdate(), "updated_at" datetime2 NOT NULL CONSTRAINT "DF_6d596d799f9cb9dac6f7bf7c23c" DEFAULT getdate(), "deleted_at" datetime2, "username" nvarchar(100) NOT NULL, "email" nvarchar(100), "mobile" nvarchar(100), "display_name" nvarchar(512), "password" nvarchar(100), "role_id" uniqueidentifier, CONSTRAINT "PK_a3ffb1c0c8416b9fc6f907b7433" PRIMARY KEY ("id"))`,
    );
    await queryRunner.query(`CREATE INDEX "IDX_f32b1cb14a9920477bcfd63df2" ON "users" ("created_by") `);
    await queryRunner.query(`CREATE INDEX "IDX_b75c92ef36f432fe68ec300a7d" ON "users" ("updated_by") `);
    await queryRunner.query(`CREATE INDEX "IDX_c9b5b525a96ddc2c5647d7f7fa" ON "users" ("created_at") `);
    await queryRunner.query(`CREATE INDEX "IDX_6d596d799f9cb9dac6f7bf7c23" ON "users" ("updated_at") `);
    await queryRunner.query(`CREATE INDEX "IDX_073999dfec9d14522f0cf58cd6" ON "users" ("deleted_at") `);
    await queryRunner.query(`CREATE UNIQUE INDEX "UIDX_User_Username" ON "users" ("username") `);
    await queryRunner.query(`CREATE UNIQUE INDEX "UIDX_User_Email" ON "users" ("email") WHERE email IS NOT NULL`);
    await queryRunner.query(`CREATE INDEX "IDX_User_Mobile" ON "users" ("mobile") `);
    await queryRunner.query(`CREATE INDEX "IDX_User_Display_Name" ON "users" ("display_name") `);
    await queryRunner.query(`CREATE INDEX "IDX_User_Role_Id" ON "users" ("role_id") `);
    await queryRunner.query(
      `CREATE TABLE "user_task_configs" ("id" uniqueidentifier NOT NULL CONSTRAINT "DF_a025c2aa82902f4729560897481" DEFAULT NEWSEQUENTIALID(), "created_by" int, "updated_by" int, "created_at" datetime2 NOT NULL CONSTRAINT "DF_3f0c3c49ab74952c14b5fce89ad" DEFAULT getdate(), "updated_at" datetime2 NOT NULL CONSTRAINT "DF_1ae801798cb6126c984431c1b62" DEFAULT getdate(), "deleted_at" datetime2, "number_of_task_per_day" int NOT NULL, "user_id" uniqueidentifier NOT NULL, "date" date NOT NULL, CONSTRAINT "PK_a025c2aa82902f4729560897481" PRIMARY KEY ("id"))`,
    );
    await queryRunner.query(`CREATE INDEX "IDX_de23bdafe97360aa65fef45988" ON "user_task_configs" ("created_by") `);
    await queryRunner.query(`CREATE INDEX "IDX_bfd508c7ee289cda9fc1004efc" ON "user_task_configs" ("updated_by") `);
    await queryRunner.query(`CREATE INDEX "IDX_3f0c3c49ab74952c14b5fce89a" ON "user_task_configs" ("created_at") `);
    await queryRunner.query(`CREATE INDEX "IDX_1ae801798cb6126c984431c1b6" ON "user_task_configs" ("updated_at") `);
    await queryRunner.query(`CREATE INDEX "IDX_795dae0c1f1e7ee852c72e8435" ON "user_task_configs" ("deleted_at") `);
    await queryRunner.query(
      `CREATE INDEX "IDX_User_Task_Number_Of_Task_Per_Day" ON "user_task_configs" ("number_of_task_per_day") `,
    );
    await queryRunner.query(`CREATE INDEX "IDX_User_Task_User_Id" ON "user_task_configs" ("user_id") `);
    await queryRunner.query(
      `CREATE TABLE "tasks" ("id" uniqueidentifier NOT NULL CONSTRAINT "DF_8d12ff38fcc62aaba2cab748772" DEFAULT NEWSEQUENTIALID(), "created_by" int, "updated_by" int, "created_at" datetime2 NOT NULL CONSTRAINT "DF_cb3724030e9674f2c17b7573aa5" DEFAULT getdate(), "updated_at" datetime2 NOT NULL CONSTRAINT "DF_02edb0ba1ef4287a15bc4c271ee" DEFAULT getdate(), "deleted_at" datetime2, "summary" nvarchar(200) NOT NULL, "description" nvarchar(500), "assignee_id" uniqueidentifier, "status" nvarchar(255) NOT NULL CONSTRAINT "DF_6086c8dafbae729a930c04d8651" DEFAULT 'todo', CONSTRAINT "PK_8d12ff38fcc62aaba2cab748772" PRIMARY KEY ("id"))`,
    );
    await queryRunner.query(`CREATE INDEX "IDX_9fc727aef9e222ebd09dc8dac0" ON "tasks" ("created_by") `);
    await queryRunner.query(`CREATE INDEX "IDX_5d927ef9f86fac1f1671d093a0" ON "tasks" ("updated_by") `);
    await queryRunner.query(`CREATE INDEX "IDX_cb3724030e9674f2c17b7573aa" ON "tasks" ("created_at") `);
    await queryRunner.query(`CREATE INDEX "IDX_02edb0ba1ef4287a15bc4c271e" ON "tasks" ("updated_at") `);
    await queryRunner.query(`CREATE INDEX "IDX_68d6a27df9f5cc119cac1df190" ON "tasks" ("deleted_at") `);
    await queryRunner.query(`CREATE INDEX "IDX_TodoTask_Assignee_Id" ON "tasks" ("assignee_id") `);
    await queryRunner.query(
      `CREATE TABLE "role_permissions" ("role_id" uniqueidentifier NOT NULL, "permission_id" uniqueidentifier NOT NULL, CONSTRAINT "PK_25d24010f53bb80b78e412c9656" PRIMARY KEY ("role_id", "permission_id"))`,
    );
    await queryRunner.query(`CREATE INDEX "IDX_178199805b901ccd220ab7740e" ON "role_permissions" ("role_id") `);
    await queryRunner.query(`CREATE INDEX "IDX_17022daf3f885f7d35423e9971" ON "role_permissions" ("permission_id") `);
    await queryRunner.query(
      `ALTER TABLE "users" ADD CONSTRAINT "FK_a2cecd1a3531c0b041e29ba46e1" FOREIGN KEY ("role_id") REFERENCES "roles"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`,
    );
    await queryRunner.query(
      `ALTER TABLE "user_task_configs" ADD CONSTRAINT "FK_e46f698148c0fcb72db3dba9800" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`,
    );
    await queryRunner.query(
      `ALTER TABLE "tasks" ADD CONSTRAINT "FK_855d484825b715c545349212c7f" FOREIGN KEY ("assignee_id") REFERENCES "users"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`,
    );
    await queryRunner.query(
      `ALTER TABLE "role_permissions" ADD CONSTRAINT "FK_178199805b901ccd220ab7740ec" FOREIGN KEY ("role_id") REFERENCES "roles"("id") ON DELETE CASCADE ON UPDATE CASCADE`,
    );
    await queryRunner.query(
      `ALTER TABLE "role_permissions" ADD CONSTRAINT "FK_17022daf3f885f7d35423e9971e" FOREIGN KEY ("permission_id") REFERENCES "permissions"("id") ON DELETE CASCADE ON UPDATE CASCADE`,
    );
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(`ALTER TABLE "role_permissions" DROP CONSTRAINT "FK_17022daf3f885f7d35423e9971e"`);
    await queryRunner.query(`ALTER TABLE "role_permissions" DROP CONSTRAINT "FK_178199805b901ccd220ab7740ec"`);
    await queryRunner.query(`ALTER TABLE "tasks" DROP CONSTRAINT "FK_855d484825b715c545349212c7f"`);
    await queryRunner.query(`ALTER TABLE "user_task_configs" DROP CONSTRAINT "FK_e46f698148c0fcb72db3dba9800"`);
    await queryRunner.query(`ALTER TABLE "users" DROP CONSTRAINT "FK_a2cecd1a3531c0b041e29ba46e1"`);
    await queryRunner.query(`DROP INDEX "IDX_17022daf3f885f7d35423e9971" ON "role_permissions"`);
    await queryRunner.query(`DROP INDEX "IDX_178199805b901ccd220ab7740e" ON "role_permissions"`);
    await queryRunner.query(`DROP TABLE "role_permissions"`);
    await queryRunner.query(`DROP INDEX "IDX_TodoTask_Assignee_Id" ON "tasks"`);
    await queryRunner.query(`DROP INDEX "IDX_68d6a27df9f5cc119cac1df190" ON "tasks"`);
    await queryRunner.query(`DROP INDEX "IDX_02edb0ba1ef4287a15bc4c271e" ON "tasks"`);
    await queryRunner.query(`DROP INDEX "IDX_cb3724030e9674f2c17b7573aa" ON "tasks"`);
    await queryRunner.query(`DROP INDEX "IDX_5d927ef9f86fac1f1671d093a0" ON "tasks"`);
    await queryRunner.query(`DROP INDEX "IDX_9fc727aef9e222ebd09dc8dac0" ON "tasks"`);
    await queryRunner.query(`DROP TABLE "tasks"`);
    await queryRunner.query(`DROP INDEX "IDX_User_Task_User_Id" ON "user_task_configs"`);
    await queryRunner.query(`DROP INDEX "IDX_User_Task_Number_Of_Task_Per_Day" ON "user_task_configs"`);
    await queryRunner.query(`DROP INDEX "IDX_795dae0c1f1e7ee852c72e8435" ON "user_task_configs"`);
    await queryRunner.query(`DROP INDEX "IDX_1ae801798cb6126c984431c1b6" ON "user_task_configs"`);
    await queryRunner.query(`DROP INDEX "IDX_3f0c3c49ab74952c14b5fce89a" ON "user_task_configs"`);
    await queryRunner.query(`DROP INDEX "IDX_bfd508c7ee289cda9fc1004efc" ON "user_task_configs"`);
    await queryRunner.query(`DROP INDEX "IDX_de23bdafe97360aa65fef45988" ON "user_task_configs"`);
    await queryRunner.query(`DROP TABLE "user_task_configs"`);
    await queryRunner.query(`DROP INDEX "IDX_User_Role_Id" ON "users"`);
    await queryRunner.query(`DROP INDEX "IDX_User_Display_Name" ON "users"`);
    await queryRunner.query(`DROP INDEX "IDX_User_Mobile" ON "users"`);
    await queryRunner.query(`DROP INDEX "UIDX_User_Email" ON "users"`);
    await queryRunner.query(`DROP INDEX "UIDX_User_Username" ON "users"`);
    await queryRunner.query(`DROP INDEX "IDX_073999dfec9d14522f0cf58cd6" ON "users"`);
    await queryRunner.query(`DROP INDEX "IDX_6d596d799f9cb9dac6f7bf7c23" ON "users"`);
    await queryRunner.query(`DROP INDEX "IDX_c9b5b525a96ddc2c5647d7f7fa" ON "users"`);
    await queryRunner.query(`DROP INDEX "IDX_b75c92ef36f432fe68ec300a7d" ON "users"`);
    await queryRunner.query(`DROP INDEX "IDX_f32b1cb14a9920477bcfd63df2" ON "users"`);
    await queryRunner.query(`DROP TABLE "users"`);
    await queryRunner.query(`DROP INDEX "UIDX_Role_name" ON "roles"`);
    await queryRunner.query(`DROP INDEX "IDX_7fd0c79dc4e6083ddea850ac38" ON "roles"`);
    await queryRunner.query(`DROP INDEX "IDX_8651ace7d160b9cf59cd0e0e2e" ON "roles"`);
    await queryRunner.query(`DROP INDEX "IDX_e5a52fc6f7a8dae64f645b0914" ON "roles"`);
    await queryRunner.query(`DROP INDEX "IDX_747b580d73db0ad78963d78b07" ON "roles"`);
    await queryRunner.query(`DROP INDEX "IDX_4a39f3095781cdd9d6061afaae" ON "roles"`);
    await queryRunner.query(`DROP TABLE "roles"`);
    await queryRunner.query(`DROP INDEX "IDX_Permission_Status" ON "permissions"`);
    await queryRunner.query(`DROP INDEX "IDX_Permission_action" ON "permissions"`);
    await queryRunner.query(`DROP INDEX "IDX_Permission_possession" ON "permissions"`);
    await queryRunner.query(`DROP INDEX "UIDX_Permission_Resource" ON "permissions"`);
    await queryRunner.query(`DROP INDEX "UIDX_Permission_Name" ON "permissions"`);
    await queryRunner.query(`DROP INDEX "IDX_1ea42cae477fc1dc619a5cd280" ON "permissions"`);
    await queryRunner.query(`DROP INDEX "IDX_b01aa3668b2d129f4bf24f002c" ON "permissions"`);
    await queryRunner.query(`DROP INDEX "IDX_337088ff813c697c964f49f58f" ON "permissions"`);
    await queryRunner.query(`DROP INDEX "IDX_58fae278276b7c2c6dde2bc19a" ON "permissions"`);
    await queryRunner.query(`DROP INDEX "IDX_c398f7100db3e0d9b6a6cd6bea" ON "permissions"`);
    await queryRunner.query(`DROP TABLE "permissions"`);
    await queryRunner.query(`DROP INDEX "IDX_8ef9962491416367a8787a1f91" ON "black_listed_token"`);
    await queryRunner.query(`DROP INDEX "IDX_1c846b375f568d13277214b05c" ON "black_listed_token"`);
    await queryRunner.query(`DROP INDEX "IDX_856aa87b893802be6b1fb2e599" ON "black_listed_token"`);
    await queryRunner.query(`DROP INDEX "IDX_019d277bbc91060a816d94ffa2" ON "black_listed_token"`);
    await queryRunner.query(`DROP INDEX "IDX_2cc9f45b3a6e9fc786ba98a41b" ON "black_listed_token"`);
    await queryRunner.query(`DROP INDEX "IDX_f4e97792cfe213d7e545ef6469" ON "black_listed_token"`);
    await queryRunner.query(`DROP TABLE "black_listed_token"`);
  }
}
