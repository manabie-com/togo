import { Entity, PrimaryGeneratedColumn, Column } from "typeorm";
import { IUser } from "../core/models/IUser";

@Entity({ name: "users" })
export class User implements IUser {
  @PrimaryGeneratedColumn() id: number;
  @Column() userId: number;
  @Column() apiKey: string;
  @Column() dailyMaximumTasks: number;
  @Column() isActive: number;
}
