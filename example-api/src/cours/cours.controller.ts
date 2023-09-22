import { Body, Controller, Delete, Get, Param, Post, Put } from '@nestjs/common';
import { CoursService } from './cours.service';
import { Cour } from './schemas/cour.schema';
import { CourDto } from './dto/cour.dto';

@Controller('cours')
export class CoursController {
    constructor(private readonly courService: CoursService) {}

    @Post()
    async create(@Body() createCatDto: CourDto): Promise<Cour>  {
      return await this.courService.create(createCatDto);
    }
  
    @Get()
    async findAll(): Promise<Cour[]> {
      return this.courService.findAll();
    }
  
    @Get(':id')
    async findOne(@Param('id') id: string): Promise<Cour> {
      return this.courService.findOne(id);
    }

    @Put(':id')
    async updateOne(@Body() updateCatDto: CourDto,@Param('id') id: string) {
      return this.courService.update(updateCatDto,id)
    }
  
    @Delete(':id')
    async delete(@Param('id') id: string) {
      return this.courService.delete(id);
    }
}
