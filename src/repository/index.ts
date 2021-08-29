import "reflect-metadata";
import { Container } from "inversify";
import { TaskRepository } from "./task.repository";
import { UserRepository } from "./user.repository";

export const REPOS = {
    UserRepository: Symbol.for("UserRepository"),
    TaskRepository: Symbol.for("TaskRepository"),
};

export class RepositoryContainer {
    public static Load(container: Container) {
        container.bind<UserRepository>(REPOS.UserRepository).to(UserRepository);
        container.bind<TaskRepository>(REPOS.TaskRepository).to(TaskRepository);
    }
}

export * from "./base.repository";
export * from "./user.repository";
export * from "./task.repository";
