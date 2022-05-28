import { Entity, Column, Index } from 'typeorm';
import { Expose } from 'class-transformer';
import { MaxLength } from 'class-validator';
import { CustomBaseEntity } from '@modules/common/entities/base.entity';

@Entity('black_listed_token')
export class BlacklistedToken extends CustomBaseEntity {
  static entityName: string = 'BlacklistedToken';

  @Expose()
  @Index()
  @Column({ name: 'user_id' })
  userId: string;

  @Expose()
  @MaxLength(3000)
  @Column({ length: 3000, name: 'token' })
  token: string;

  constructor(partial: Partial<BlacklistedToken>) {
    super();
    this.entityName = BlacklistedToken.entityName;
    Object.assign(this, partial);
  }
}
