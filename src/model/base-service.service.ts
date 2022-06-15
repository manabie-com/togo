// import { Injectable, Inject } from '@nestjs/common';
// import { LIMIT_TASK } from 'src/constance/variable';
// import { LimitTask } from './limitTask.entity';

// @Injectable()
// export class LimitTaskService {
//   constructor(
//     @Inject(LIMIT_TASK)
//     private limitTaskModel: typeof LimitTask
//   ) {}

//   async findAll(): Promise<LimitTask[]> {
//     return this.limitTaskModel.findAll<LimitTask>();
//   }
// }