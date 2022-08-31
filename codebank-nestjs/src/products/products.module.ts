import { productProviders } from './products.provider';
import { DatabaseModule } from './../database/database.module';
import { Module } from '@nestjs/common';
import { ProductsService } from './products.service';
import { ProductsController } from './products.controller';

@Module({
  imports: [DatabaseModule],
  controllers: [ProductsController],
  providers: [
    ...productProviders,
    ProductsService]
})
export class ProductsModule {}
