import { Test } from '@nestjs/testing';
import { TypeOrmModule } from '@nestjs/typeorm';
import { JwtModule } from '@nestjs/jwt';
import { ConfigModule, ConfigService } from '@nestjs/config';
import * as bcrypt from 'bcrypt';
import { UserController } from './user.controller';
import { UserService } from './user.service';
import { UserRepository } from './user.repository';
import { UserEntity } from './entities/user.entity';
import { TaskEntity } from './entities/task.entity';
import { Constants } from '../../utils/constants';
import * as httpMocks from 'node-mocks-http';
describe('UserController', () => {
    let userController: UserController;
    let userService: UserService;
    let userRepository: UserRepository;

    beforeAll(async () => {
        const moduleRef = await Test.createTestingModule({
            imports: [
                TypeOrmModule.forRoot({
                    type: 'mysql',
                    host: '127.0.0.1',
                    port: 3306,
                    username: 'root',
                    password: '1',
                    database: 'interview',
                    synchronize: true,
                    entities: [UserEntity, TaskEntity],
                }),
                TypeOrmModule.forFeature([UserEntity, TaskEntity]),
                JwtModule.registerAsync({
                    imports: [ConfigModule],
                    useFactory: (config: ConfigService) => {
                        return {
                            secret: config.get<string>("JWT_SECRET")
                        }
                    },
                    inject: [ConfigService]
                }),
                ConfigModule.forRoot({
                    isGlobal: true,
                    envFilePath: ".env"
                }),
            ],
            controllers: [UserController],
            providers: [UserService, UserRepository],
        }).compile();

        userService = moduleRef.get<UserService>(UserService);
        userController = moduleRef.get<UserController>(UserController);
        userRepository = moduleRef.get<UserRepository>(UserRepository);
    });


    describe('Login', () => {
        it('should login successfull', async () => {
            const request = {
                email: "hoaidien93@gmail.com",
                password: "hoaidien93@gmail.com"
            }
            const result = new UserEntity({
                created_at: new Date(),
                id: 1,
                password: bcrypt.hashSync(request.password, Constants.SALT_OR_ROUNDS),
                fullName: "Testing",
                role: Constants.USER_ROLE,
                email: request.email,
                updated_at: new Date()
            })

            jest.spyOn(userRepository, 'findUserByEmail').mockImplementation(() => Promise.resolve(result));
            expect((await userController.postLogin(request)).code).toBe(Constants.SUCCESS_CODE)
        });
        it('should login failed', async () => {
            const request = {
                email: "hoaidien93@gmail.com",
                password: "hoaidien93@gmail.com"
            }
            const result = null;
            jest.spyOn(userRepository, 'findUserByEmail').mockImplementation(() => Promise.resolve(result));
            expect((await userController.postLogin(request)).code).toBe(Constants.FAIL_CODE)
        });
    })

    describe('Register', () => {
        it('should register successfull', async () => {
            const request = {
                email: "hoaidien93@gmail.com",
                password: "hoaidien93@gmail.com",
                fullName: "Hoài Diễn",
            }
            const result = new UserEntity({
                created_at: new Date(),
                id: 1,
                password: bcrypt.hashSync(request.password, Constants.SALT_OR_ROUNDS),
                fullName: "Testing",
                role: Constants.USER_ROLE,
                email: request.email,
                updated_at: new Date()
            })
            jest.spyOn(userRepository, 'register').mockImplementation(() => Promise.resolve(result));
            expect((await userController.postRegister(request)).code).toBe(Constants.SUCCESS_CODE)
        });

        it('should register failed', async () => {
            const request = {
                email: "hoaidien93@gmail.com",
                password: "hoaidien93@gmail.com",
                fullName: "Hoài Diễn",
                provinceId: 1,
            }
            const result = null
            jest.spyOn(userRepository, 'register').mockImplementation(() => Promise.resolve(result));
            expect((await userController.postRegister(request)).code).toBe(Constants.FAIL_CODE)
        });
    })

    describe('Find User', () => {
        it('find user successfull', async () => {
            const result = {
                listUsers: [new UserEntity({
                    id: 1,
                    email: "hoaidien93@gmail.com",
                    fullName: "Hoài Diễn",
                })],
                pageIndex: 1,
                totalPage: 1
            }
            jest.spyOn(userRepository, 'find').mockImplementation(() => Promise.resolve(result));
            expect((await userController.findUser({
                pageIndex: 1,
                perPage: 10,
            })).data).toBe(result)
        });
    })

    describe('Add Task', () => {
        it('should add task successfull', async () => {
            const req = httpMocks.createRequest({
                method: "POST",
                url: "/add-task",
                user: {
                    id: 1,
                    email: "hoaidien93@gmail.com"
                }
            })
            const taskDTO = {
                name: "Sample task",
                note: "Nothing"
            }
            jest.spyOn(userRepository, 'addNewTask').mockImplementation(() => Promise.resolve(true));
            expect((await userController.addTask(req, taskDTO)).code).toBe(Constants.SUCCESS_CODE)
        });

        it('should add task failed', async () => {
            const req = httpMocks.createRequest({
                method: "POST",
                url: "/add-task",
                user: {
                    id: 1,
                    email: "hoaidien93@gmail.com"
                }
            })
            const taskDTO = {
                name: "Sample task",
                note: "Nothing"
            }
            jest.spyOn(userRepository, 'addNewTask').mockImplementation(() => Promise.resolve(false));
            expect((await userController.addTask(req, taskDTO)).code).toBe(Constants.FAIL_CODE)
        });
    })

    describe('Config max Task', () => {
        it('should config task successfull', async () => {
            const configDTO = {
                id: 1,
                maxTask: 2
            }
            jest.spyOn(userRepository, 'configureMaxTask').mockImplementation(() => Promise.resolve(true));
            expect((await userController.configureMaxTask(configDTO)).code).toBe(Constants.SUCCESS_CODE)
        });

        it('should config task failed', async () => {
            const configDTO = {
                id: 1,
                maxTask: 2
            }
            jest.spyOn(userRepository, 'configureMaxTask').mockImplementation(() => Promise.resolve(false));
            expect((await userController.configureMaxTask(configDTO)).code).toBe(Constants.FAIL_CODE)
        });
    })
});