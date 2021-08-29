import "reflect-metadata";
import { Container } from "inversify";
import { UserService } from "./user.service";
import { TaskService } from "./task.service";

export const SERVICES = {
    UserService: Symbol.for("UserService"),
    TaskService: Symbol.for("TaskService"),
};

export class ServiceContainer {

    public static Load(container: Container) {
        container.bind<UserService>(SERVICES.UserService).to(UserService);
        container.bind<TaskService>(SERVICES.TaskService).to(TaskService);
    }
}

export * from "./user.service";
export * from "./task.service";
