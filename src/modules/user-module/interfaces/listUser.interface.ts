import { UserEntity } from "../entities/user.entity";

export interface ListUser{
    pageIndex: number,
    totalPage: number,
    listUsers: UserEntity[]
}