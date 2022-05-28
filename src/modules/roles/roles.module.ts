import { Module } from '@nestjs/common';
import { RolesService } from './roles.service';

@Module({
  // controllers: [RolesController],
  providers: [RolesService],
})
export class RolesModule {}
