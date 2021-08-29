import "reflect-metadata";
import { Container } from "inversify";
import { RepositoryContainer } from "../repository";
import { ServiceContainer } from "../service";
import { HookContainer } from "../hook";
import { CoreContainer } from "./index";

export class AppContainer {
    public static Load(): Container {
        const container = new Container({skipBaseClassChecks: true});
        CoreContainer.Load(container);
        RepositoryContainer.Load(container);
        ServiceContainer.Load(container);
        HookContainer.Load(container);
        return container;
    }
}
