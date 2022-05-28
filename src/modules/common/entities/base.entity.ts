import { ApiProperty } from '@nestjs/swagger';
import {
  BaseEntity,
  Column,
  CreateDateColumn,
  DeleteDateColumn,
  Index,
  PrimaryGeneratedColumn,
  UpdateDateColumn,
} from 'typeorm';
import { Expose, Exclude } from 'class-transformer';

export abstract class CustomBaseEntity extends BaseEntity {
  @Expose()
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @Exclude()
  @Index()
  @Column({ name: 'created_by', nullable: true })
  createdById: number;

  @Exclude()
  @Index()
  @Column({ name: 'updated_by', nullable: true })
  updatedById: number;

  @ApiProperty()
  @Expose()
  @Index()
  @CreateDateColumn({ name: 'created_at' })
  createdAt: Date;

  @ApiProperty()
  @Expose()
  @Index()
  @UpdateDateColumn({ name: 'updated_at' })
  updatedAt: Date;

  @Exclude()
  @Index()
  @DeleteDateColumn({ name: 'deleted_at', nullable: true })
  deletedAt?: Date;

  @Exclude()
  // eslint-disable-next-line @typescript-eslint/naming-convention
  private _entityName: string = '';

  public get entityName(): string {
    return this._entityName;
  }

  public set entityName(value: string) {
    this._entityName = value;
  }
}
