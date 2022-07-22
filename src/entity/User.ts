/* 
 * TYPEORM entity 
 * does the heavy lifting when accessing database
 */
import { Entity, PrimaryGeneratedColumn, Column, Unique } from 'typeorm';

@Entity()
@Unique(['username'])
export class User {
	@PrimaryGeneratedColumn()
	id: number

	@Column()
	username: string

	@Column()
	password: string

	@Column()
	type: string
}
