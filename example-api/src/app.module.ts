import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { MongooseModule } from '@nestjs/mongoose';
import { CoursModule } from './cours/cours.module';

@Module({
  imports: [MongooseModule.forRoot(process.env.DB_URL), CoursModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
