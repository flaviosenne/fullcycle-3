import { DataSource } from 'typeorm';
import { Product } from './product.entity';

export const productProviders = [
    {
        provide: 'PRODUCT_REPOSITORY',
        useFactory: (datasource: DataSource) => datasource.getRepository(Product),
        inject: ['DATABASE_CONNECTION']
    }
]