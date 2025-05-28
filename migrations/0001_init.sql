create table
    if not exists "customers" (
        id serial primary key,
        name varchar(100) not null,
        email varchar(100) not null,
        tax_id varchar(50) not null,
        created_at timestamp default current_timestamp,
        updated_at timestamp default current_timestamp,
        constraint "unique_email" unique ("email"),
        constraint "unique_tax_id" unique ("tax_id")
    );

create table
    if not exists "payments" (
        id serial primary key,
        amount numeric(10, 2) not null,
        payment_date timestamp default current_timestamp,
        status varchar(50) not null,
        provider varchar(50) not null,
        created_at timestamp default current_timestamp,
        updated_at timestamp default current_timestamp
    );

create table
    if not exists "orders" (
        id serial primary key,
        order_date timestamp default current_timestamp,
        notification_attempts int not null,
        status varchar(50) not null,
        total_amount numeric(10, 2) not null,
        created_at timestamp default current_timestamp,
        updated_at timestamp default current_timestamp,
        customer_id integer references customers (id) on delete cascade,
        payment_id integer references payments (id) on delete cascade
    );

create table
    if not exists "product_categories" (
        id serial primary key,
        name varchar(100) not null,
        description text,
        created_at timestamp default current_timestamp,
        updated_at timestamp default current_timestamp
    );

create table
    if not exists "products" (
        id serial primary key,
        name varchar(100) not null,
        description text,
        price numeric(10, 2) not null,
        image_url varchar(255),
        preparation_time integer not null,
        created_at timestamp default current_timestamp,
        updated_at timestamp default current_timestamp,
        category_id integer references product_categories (id) on delete cascade
    );

create table
    if not exists "order_items" (
        id serial primary key,
        quantity integer not null,
        price numeric(10, 2) not null,
        created_at timestamp default current_timestamp,
        updated_at timestamp default current_timestamp,
        order_id integer references orders (id) on delete cascade,
        product_id integer references products (id) on delete cascade
    );


insert into customers (name, email, tax_id)
values ('João Silva', 'joao.silva@email.com', '90103576002');

insert into product_categories (name, description) values
('Lanche', 'Sanduíches e hambúrgueres'),
('Acompanhamento', 'Batatas, saladas e outros acompanhamentos'),
('Bebida', 'Refrigerantes, sucos e outras bebidas'),
('Sobremesa', 'Doces e sobremesas');

insert into products (name, description, price, image_url, preparation_time, category_id) values
('Hambúrguer Clássico', 'Pão, carne, queijo, alface e tomate', 18.90, 'https://images.unsplash.com/photo-1550547660-d9450f859349', 15, 1),
('Cheeseburger Bacon', 'Hambúrguer com queijo e bacon crocante', 22.50, 'https://images.unsplash.com/photo-1553979459-d2229ba7433b', 18, 1);

insert into products (name, description, price, image_url, preparation_time, category_id) values
('Batata Frita', 'Porção de batatas fritas crocantes', 9.90, 'https://images.unsplash.com/photo-1598679253544-2c97992403ea', 8, 2),
('Onion Rings', 'Anéis de cebola empanados', 11.50, 'https://images.unsplash.com/photo-1639024471283-03518883512d', 10, 2);

insert into products (name, description, price, image_url, preparation_time, category_id) values
('Refrigerante Lata', '350ml - Coca-Cola, Guaraná ou Fanta', 6.00, 'https://images.unsplash.com/photo-1527960471264-932f39eb5846', 2, 3),
('Suco Natural', 'Suco de laranja ou limão', 7.50, 'https://images.unsplash.com/photo-1600271886742-f049cd451bba', 3, 3);

insert into products (name, description, price, image_url, preparation_time, category_id) values
('Sorvete', 'sorvete', 14.00, 'https://images.unsplash.com/photo-1497034825429-c343d7c6a68f', 7, 4),
('Petit Gateau', 'Bolo de chocolate com recheio cremoso e sorvete', 16.00, 'https://t4.ftcdn.net/jpg/02/21/31/01/240_F_221310131_cUVS5tnUMG1qv3GWzzj8w2bgDUtLSmRv.jpg', 10, 4);