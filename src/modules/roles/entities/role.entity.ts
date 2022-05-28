import { ApiProperty } from '@nestjs/swagger';
import { Index, Entity, Column, JoinTable, ManyToMany } from 'typeorm';
import { Expose } from 'class-transformer';
import { IsString, MaxLength } from 'class-validator';
import { CustomBaseEntity } from '@modules/common/entities/base.entity';
import { Permission } from '@modules/permissions/permission.entity';

@Entity('roles')
export class Role extends CustomBaseEntity {
  static entityName: string = 'Role';

  @Expose()
  @ApiProperty()
  @Column({ length: 100, unique: true, type: 'nvarchar', nullable: true })
  @Index('UIDX_Role_name')
  @MaxLength(100)
  @IsString()
  name: string;

  @Expose()
  @ApiProperty()
  @Column({ length: 300, nullable: true })
  @MaxLength(300)
  description: string;

  @ManyToMany(() => Permission)
  @JoinTable({
    name: 'role_permissions',
    joinColumn: {
      name: 'role_id',
      referencedColumnName: 'id',
    },
    inverseJoinColumn: {
      name: 'permission_id',
      referencedColumnName: 'id',
    },
  })
  permissions: Permission[];

  constructor(partial: Partial<Role>) {
    super();
    this.entityName = Role.entityName;
    Object.assign(this, partial);
  }
}
