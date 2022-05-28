/* eslint-disable security/detect-non-literal-fs-filename */
/* eslint-disable security/detect-non-literal-require */
import * as fs from 'fs';
import * as Path from 'path';
import * as request from 'supertest';
import { useContainer } from 'class-validator';
import { DatabaseService } from './database/database.service';
import { HttpService, Injectable, ValidationPipe } from '@nestjs/common';
import { Repository, ObjectType } from 'typeorm';
import { getRepositoryToken } from '@nestjs/typeorm';
import { Connection } from 'typeorm';
import { UserService } from '@modules/users/user.service';
import { User } from '@modules/users/entities/user.entity';
import { RolesService } from '@modules/roles/roles.service';
import { Role } from '@modules/roles/entities/role.entity';
import { TodoTaskService } from '@modules/todo-tasks/todo-task.service';
import { TodoTask } from '@modules/todo-tasks/entities/todo-task.entity';

/**
 * This class is used to support database
 * tests with unit tests in NestJS.
 *
 * This class is inspired by https://github.com/jgordor
 * https://github.com/nestjs/nest/issues/409#issuecomment-364639051
 */
@Injectable()
export class TestUtils {
  /**
   * database service
   */
  databaseService: DatabaseService;

  /**
   * cache for loaded data
   */
  protected cache: { [key: string]: any[] } = {};

  /**
   * Creates an instance of TestUtils
   */
  constructor(databaseService: DatabaseService) {
    if (process.env.NODE_ENV !== 'test') {
      throw new Error('ERROR-TEST-UTILS-ONLY-FOR-TESTS');
    }

    this.databaseService = databaseService;
  }

  /**
   * Shutdown the http server
   * and close database connections
   */
  async shutdownServer(server) {
    await server.httpServer.close();
    await this.closeDbConnection();
  }

  /**
   * Closes the database connections
   */
  async closeDbConnection() {
    const connection = this.databaseService.connection;

    if (connection.isConnected) {
      await this.databaseService.connection.close();
    }
  }

  /**
   * Return the repository
   * @param entity - The entity type
   */
  async getRepository<T>(entity: ObjectType<T>): Promise<Repository<T>> {
    const result = this.databaseService.connection.entityMetadatas.find((x) => x.name == entity.name);

    if (!result) {
      throw new Error(
        `Specified entity (${entity.name}) is not initialized. Please add the entity into module/test/database/database-ormconfig.constant.ts`,
      );
    }

    return this.databaseService.getRepository(entity);
  }

  /**
   * Returns the order id
   * @param entityName - The entity name of which you want to have the order from
   */
  getOrder(entityName) {
    const order: string[] = JSON.parse(fs.readFileSync(Path.join(__dirname, '../test/fixtures/_order.json'), 'utf8'));

    return order.indexOf(entityName);
  }

  /**
   * Returns the entites of the database
   */
  async getEntities() {
    const entities = [];

    this.databaseService.connection.entityMetadatas.forEach((x) =>
      entities.push({ name: x.name, tableName: x.tableName, order: this.getOrder(x.name) }),
    );

    return entities;
  }

  /**
   * Cleans the database and reloads the entries
   */
  async reloadFixtures() {
    const entities = await this.getEntities();

    await this.cleanAll(entities);
    await this.loadAll(entities);
  }

  /**
   * Cleans all the entities
   */
  async cleanAll(entities) {
    try {
      for (const entity of entities.sort((a, b) => b.order - a.order)) {
        const repository = await this.databaseService.getRepository(entity.name);

        await repository.query(`DELETE FROM ${entity.tableName};`);
      }
    } catch (error) {
      throw new Error(`ERROR: Cleaning test db: ${error}`);
    }
  }

  /**
   * Insert the data from the src/test/fixtures folder
   */
  async loadAll(entities: any[]) {
    try {
      for (const entity of entities.sort((a, b) => a.order - b.order)) {
        const repository = await this.databaseService.getRepository(entity.name);

        if (entity.name in this.cache) {
          const items = this.cache[entity.name];

          await repository.save(items);
        } else {
          const fixtureFile = Path.join(__dirname, `../test/fixtures/${entity.name}.ts`);

          if (fs.existsSync(fixtureFile)) {
            const items = require(fixtureFile);

            await repository.save(items);
            this.cache[entity.name] = items;
          }
        }
      }
    } catch (error) {
      throw new Error(`ERROR [TestUtils.loadAll()]: Loading fixtures on test db: ${error}`);
    }
  }

  getConnectionServiceGroup() {
    return [
      {
        provide: Connection,
        useValue: this.databaseService.connection,
      },
    ];
  }

  getCommonServiceGroup() {
    return [...this.getConnectionServiceGroup()];
  }

  async startApp(app: any, selectModule?: any) {
    app.useGlobalPipes(
      new ValidationPipe({
        transform: true,
        whitelist: true,
        forbidNonWhitelisted: true,
        skipMissingProperties: false,
        forbidUnknownValues: false,
      }),
    );

    if (selectModule) {
      useContainer(app.select(selectModule), { fallbackOnErrors: true });
    }

    await app.init();

    return request(app.getHttpServer());
  }

  getUserServiceGroup() {
    return [
      UserService,
      {
        provide: getRepositoryToken(User),
        useValue: this.getRepository(User),
      },
      ...this.getConnectionServiceGroup(),
      ...this.getRoleServiceGroup(),
    ];
  }

  getRoleServiceGroup() {
    return [
      RolesService,
      {
        provide: getRepositoryToken(Role),
        useValue: this.getRepository(Role),
      },
    ];
  }

  getTodoTaskServiceGroup() {
    return [
      TodoTaskService,
      {
        provide: getRepositoryToken(TodoTask),
        useValue: this.getRepository(TodoTask),
      },
    ];
  }
}
