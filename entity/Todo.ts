import { Entity, PrimaryGeneratedColumn, Column } from "typeorm";
import { ITodo } from "../core/models/ITodo";

@Entity({ name: "todos" })
export class Todo implements ITodo {
  @PrimaryGeneratedColumn() id: number;
  @Column() task: string;
  @Column() userId: number;
  @Column() creationDate?: string;
}
