create table if not exists costumers (
    id serial primary key,
    name varchar(100) not null,
    email varchar(100) not null unique,
    tax_id varchar(50) not null,
    created_at timestamp default current_timestamp
);

create table if not exists orders (
    id serial primary key,
    order_date timestamp default current_timestamp,
    status varchar(50) not null,
    total_amount numeric(10, 2) not null,
    created_at timestamp default current_timestamp,
    
    customer_id integer references costumers(id) on delete cascade
);

create table if not exists categories_products (
    id serial primary key,
    name varchar(100) not null,
    description text,
    created_at timestamp default current_timestamp
);

create table if not exists products (
    id serial primary key,
    name varchar(100) not null,
    description text,
    price numeric(10, 2) not null,
    image_url varchar(255),
    preparation_time integer not null,
    created_at timestamp default current_timestamp,

    category_id integer references categories_products(id) on delete cascade
);

create table if not exists order_items (
    id serial primary key, 
    quantity integer not null,
    price numeric(10, 2) not null,
    created_at timestamp default current_timestamp,

    order_id integer references orders(id) on delete cascade,
    product_id integer references products(id) on delete cascade
);