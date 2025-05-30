model Product {
    id: string;
    name: string;
    description?: string;
    price: float32;
    categoryId: string;
    imageUrl?: string;
    available: boolean;
}

model ProductInput {
    name: string;
    description?: string;
    price: float32;
    categoryId: string;
    imageUrl?: string;
    available: boolean;
}