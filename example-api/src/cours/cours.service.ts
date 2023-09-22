import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Cour } from './schemas/cour.schema';
import { Model } from 'mongoose';
import { CourDto } from './dto/cour.dto';

@Injectable()
export class CoursService {

    constructor(@InjectModel(Cour.name) private readonly CourModel: Model<Cour>) {}

      async create(courDto: CourDto): Promise<Cour> {
        const createdCour = await this.CourModel.create(courDto);
        console.log(createdCour);
        return createdCour;
      }
    
      async findAll(): Promise<Cour[]> {
        return this.CourModel.find().exec();
      }
    
      async findOne(id: string): Promise<Cour> {
        return this.CourModel.findOne({ _id: id }).exec();
      }

      async update(courDto: CourDto, id:string): Promise<any>{
        let updateCour = await this.CourModel.updateOne(courDto).exec()
        console.log(this.findOne(id))
        return this.findOne(id);
      }

      async delete(id: string) {
        const deletedCour = await this.CourModel
          .findByIdAndRemove({ _id: id })
          .exec();
        return deletedCour;
      }

}
