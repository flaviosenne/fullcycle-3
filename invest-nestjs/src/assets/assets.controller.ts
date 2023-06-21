import { Body, Controller, Post } from '@nestjs/common';
import { AssetsService } from './assets.service';

@Controller('assets')
export class AssetsController {
  constructor(private readonly assetsService: AssetsService) {}

  @Post()
  async create(@Body() body: {id: string, symbol: string, price: number}){
    return await this.assetsService.create(body)
  }
}
