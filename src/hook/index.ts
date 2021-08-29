import "reflect-metadata";
import { Container } from "inversify";
import { AuthHook } from "./auth.hook";
import { PasswordHook } from "./password.hook";
import { TaskLimitHook } from "./task-limit.hook";

export const HOOKS = {
    AuthHook: Symbol.for("AuthHook"),
    PasswordHook: Symbol.for("PasswordHook"),
    TaskLimitHook: Symbol.for("TaskLimitHook"),
};

export class HookContainer {

    public static Load(container: Container) {
        container.bind<AuthHook>(HOOKS.AuthHook).to(AuthHook);
        container.bind<PasswordHook>(HOOKS.PasswordHook).to(PasswordHook);
        container.bind<TaskLimitHook>(HOOKS.TaskLimitHook).to(TaskLimitHook);
    }
}
