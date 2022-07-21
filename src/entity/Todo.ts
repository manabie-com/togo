/* 
 * TYPEORM entity 
 * does the heavy lifting when accessing database
 */
import { Entity, PrimaryGeneratedColumn, Column, CreateDateColumn } from 'typeorm';

@Entity()
export class Todo {
	@PrimaryGeneratedColumn()
	id: number;

	@Column()
	userId: number;

	@Column()
	task: string;

	@Column()
	done: boolean;

	@CreateDateColumn()
	createdAt: Date;
}
