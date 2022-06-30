import { Controller, Get } from '@nestjs/common';

@Controller('my-first')
export class MyFirstController {
    @Get('hello-word')
    index(){
        return {'key': 'value'}
    }
}
