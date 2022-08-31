import { DataSource } from "typeorm";

export const databaseProviders = [
    {
        provide: 'DATABASE_CONNECTION',
        useFactory: async () => {
            const dataSource = new DataSource({
                type: 'postgres',
                host: 'db',
                port: 5432,
                username: 'postgres',
                password: 'root',
                database: 'store',
                entities: [
                    __dirname + '/../**/*.entity{.ts,.js}',
                ],
                migrations: [__dirname+'/../migrations/**/*{.ts,.js}'],
                synchronize: true,
              });
            
            return dataSource.initialize()
        }
    }
]