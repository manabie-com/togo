import { index, modelOptions, prop, Ref } from '@typegoose/typegoose';
import { ENTITY_OPTIONS } from "../config";
import { User } from "./user";
import { BaseEntity } from "./base-entity";
import { Column, Entity, PrimaryGeneratedColumn } from "typeorm";

@index({content: 'text'})
@modelOptions(ENTITY_OPTIONS('tasks'))
@Entity('tasks')
export class Task extends BaseEntity {

    @prop({required: [true, 'USER_NAME_REQUIRED'], trim: true, unique: true, immutable: true})
    @PrimaryGeneratedColumn('uuid')
    public id!: string;

    @prop({required: [true, 'TASK_CONTENT_REQUIRED'], trim: true})
    @Column({type: 'text', nullable: false})
    public content!: string;

    @prop({required: [true, 'USER_ID_REQUIRED'], ref: () => User})
    public user_id!: string;

    @prop({ref: 'User', foreignField: '_id', localField: 'userId', justOne: true})
    public user?: Ref<User>;
}
