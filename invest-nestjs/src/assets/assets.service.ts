import { Injectable } from '@nestjs/common';
import { PrismaService } from 'src/prisma/prisma/prisma.service';

@Injectable()
export class AssetsService {

    constructor(private prismaService: PrismaService){}

    async create(data: {id: string, symbol: string, price: number}){
        return await this.prismaService.asset.create({
            data
        })
    }
}
