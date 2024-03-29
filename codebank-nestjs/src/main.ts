import { ConfigService } from '@nestjs/config';
import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  const config: ConfigService = app.get(ConfigService);
  const port: number = config.get<number>('PORT');


  await app.listen(port, ()=> {
    console.log('[WEB]', config.get<string>('BASE_URL'), port);
  });
}
bootstrap();
