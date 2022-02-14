import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

export type TaskDocument = Task & Document;

@Schema({ timestamps: true })
export class Task {
  @Prop({ required: true })
  owner: string;

  @Prop({ required: true })
  description: string;
}

export const TaskSchema = SchemaFactory.createForClass(Task);
