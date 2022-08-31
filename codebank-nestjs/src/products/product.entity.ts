import { Column, Entity, PrimaryColumn } from "typeorm";

@Entity()
export class Product {
    @PrimaryColumn()
    id: string

    @Column()
    name: string
    
    @Column()
    description: string
    
    @Column()
    image_url: string
    
    @Column()
    slug: string
    
    @Column()
    price: number
    
    @Column()
    created_at: string
}
