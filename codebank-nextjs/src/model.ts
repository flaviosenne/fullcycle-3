export interface Product {
    id: string
    name: string
    description: string
    image_url: string
    slug: string
    price: number
    created_at: string
}

export const products: Product[] = [
    {
      id: 'uuid',
      name: 'teste',
      description: 'muito texto',
      price: 50.90,
      image_url: 'https://source.unsplash.com/random?product,'+Math.random(),
      slug: 'produto-teste',
      created_at: '2021-06-06T00:00:00'
    },
    {
      id: 'uuid',
      name: 'teste',
      description: 'muito texto',
      price: 50.90,
      image_url: 'https://source.unsplash.com/random?product,'+ Math.random(),
      slug: 'produto-teste',
      created_at: '2021-06-06T00:00:00'
    }
  ] 
  