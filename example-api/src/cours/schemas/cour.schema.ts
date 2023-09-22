import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { HydratedDocument } from 'mongoose';

export type CourDocument = HydratedDocument<Cour>;

@Schema()
export class Cour {
  @Prop()
  name: string;

  @Prop()
  time: number;

  @Prop()
  summary: string;
}

export const CourSchema = SchemaFactory.createForClass(Cour);