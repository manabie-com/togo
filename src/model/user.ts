import { modelOptions, prop, Ref } from '@typegoose/typegoose';
import { ENTITY_OPTIONS } from "../config";
import { Task } from "./task";
import { BaseEntity } from "./base-entity";
import { Column, Entity, PrimaryColumn } from "typeorm";

@modelOptions(ENTITY_OPTIONS('users'))
@Entity('users')
export class User extends BaseEntity {

    @prop({required: [true, 'USER_NAME_REQUIRED'], trim: true, unique: true, immutable: true})
    @PrimaryColumn({type: 'text'})
    public id!: string;

    @prop({required: [true, 'PASSWORD_REQUIRED'], select: false})
    @Column({type: 'text'})
    public password!: string;

    @prop({required: [true, 'MAX_TO_DO_REQUIRED'], default: 5, min: 0})
    @Column({type: 'int', default: 5})
    public max_todo!: number;

    @prop({ref: 'Task', foreignField: '_id', localField: 'userId'})
    public tasks?: Ref<Task>[];
}
