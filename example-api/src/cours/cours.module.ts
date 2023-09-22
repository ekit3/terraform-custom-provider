import { Module } from '@nestjs/common';
import { CoursService } from './cours.service';
import { CoursController } from './cours.controller';
import { Cour, CourSchema } from './schemas/cour.schema';
import { MongooseModule } from '@nestjs/mongoose';

@Module({
  imports: [MongooseModule.forFeature([{ name: Cour.name, schema: CourSchema }])],
  providers: [CoursService],
  controllers: [CoursController]
})
export class CoursModule {}
