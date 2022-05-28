import { ApiProperty } from '@nestjs/swagger';
import { Index, Entity, Column } from 'typeorm';
import { Expose } from 'class-transformer';
import { IsEnum, IsString, MaxLength } from 'class-validator';
import { CustomBaseEntity } from '@modules/common/entities/base.entity';
import { PermissionAction } from './permission.action.enum';
import { PermissionStatus } from './permission.status.enum';
import { PermissionPossession } from './permission.possession.enum';

@Entity('permissions')
export class Permission extends CustomBaseEntity {
  static entityName: string = 'Permission';

  @Expose()
  @ApiProperty()
  @Column({ length: 100, unique: true })
  @Index('UIDX_Permission_Name')
  @MaxLength(100)
  @IsString()
  name: string;

  @Expose()
  @ApiProperty()
  @Column({ length: 100 })
  @Index('UIDX_Permission_Resource')
  @MaxLength(100)
  @IsString()
  resource: string;

  @Expose()
  @ApiProperty({
    enum: PermissionPossession,
    enumName: 'PermissionPossession',
  })
  @Index('IDX_Permission_possession')
  @Column({ type: 'simple-enum', default: PermissionPossession.Own, enum: PermissionPossession })
  @IsEnum(PermissionPossession, { each: true })
  possession: PermissionPossession;

  @Expose()
  @ApiProperty({
    enum: PermissionAction,
    enumName: 'PermissionAction',
  })
  @Index('IDX_Permission_action')
  @MaxLength(100)
  @IsString()
  @Column({ length: 100, default: PermissionAction.List })
  action: string;

  @Expose()
  @ApiProperty({
    enum: PermissionStatus,
    enumName: 'PermissionStatus',
  })
  @Index('IDX_Permission_Status')
  @Column({ type: 'simple-enum', default: PermissionStatus.Active, enum: PermissionStatus })
  status: PermissionStatus;

  constructor(partial: Partial<Permission>) {
    super();
    this.entityName = Permission.entityName;
    Object.assign(this, partial);
  }
}
