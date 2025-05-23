model Order {
    id: int;
    order_date: string; 
    notification_attempts: int;
    status: OrderStatus;
    total_amount: float64;
    created_at: string; 
    customer_id: int;
    items: OrderItem[];
}

model OrderItem {
    id: int;
    quantity: int;
    price: float64;
    created_at: string; 
    order_id: int;
    product_id: int;
}

model OrderInput {
    customer_id: int;
    items: OrderItemInput[];
}

model OrderItemInput {
    product_id: int;
    quantity: int;
}

model OrderCreationResponse {
    message: string;
    orderId: string;
}

model UpdateOrderInput {
    status: OrderStatus;
}

model OrderUpdateResponse {
    message: string;
}
