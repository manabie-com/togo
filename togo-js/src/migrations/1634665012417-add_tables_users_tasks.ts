import {MigrationInterface, QueryRunner} from "typeorm";

export class addTablesUsersTasks1634665012417 implements MigrationInterface {
    name = 'addTablesUsersTasks1634665012417'

    public async up(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`CREATE TABLE "user" ("id" SERIAL NOT NULL, "userId" character varying NOT NULL, "password" character varying NOT NULL, "maxTodo" integer NOT NULL DEFAULT '5', "createdAt" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updatedAt" TIMESTAMP WITH TIME ZONE DEFAULT now(), CONSTRAINT "PK_cace4a159ff9f2512dd42373760" PRIMARY KEY ("id"))`);
        await queryRunner.query(`CREATE UNIQUE INDEX "IDX_d72ea127f30e21753c9e229891" ON "user" ("userId") `);
        await queryRunner.query(`CREATE TABLE "task" ("id" SERIAL NOT NULL, "content" character varying NOT NULL, "createdAt" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), "updatedAt" TIMESTAMP WITH TIME ZONE DEFAULT now(), "userId" integer, CONSTRAINT "PK_fb213f79ee45060ba925ecd576e" PRIMARY KEY ("id"))`);
        await queryRunner.query(`CREATE INDEX "IDX_4a54e88f8c42954be40d039f6a" ON "task" ("createdAt") `);
        await queryRunner.query(`ALTER TABLE "task" ADD CONSTRAINT "FK_f316d3fe53497d4d8a2957db8b9" FOREIGN KEY ("userId") REFERENCES "user"("id") ON DELETE NO ACTION ON UPDATE NO ACTION`);
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`ALTER TABLE "task" DROP CONSTRAINT "FK_f316d3fe53497d4d8a2957db8b9"`);
        await queryRunner.query(`DROP INDEX "public"."IDX_4a54e88f8c42954be40d039f6a"`);
        await queryRunner.query(`DROP TABLE "task"`);
        await queryRunner.query(`DROP INDEX "public"."IDX_d72ea127f30e21753c9e229891"`);
        await queryRunner.query(`DROP TABLE "user"`);
    }

}
