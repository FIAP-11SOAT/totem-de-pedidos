create table
    if not exists "customers" (
        id serial primary key,
        name varchar(100) not null,
        email varchar(100) not null,
        tax_id varchar(50) not null,
        created_at timestamp default current_timestamp,
        updated_at timestamp default current_timestamp,
        constraint "unique_email" unique ("email")
        constraint "unique_tax_id" unique ("tax_id"),
    );

create table
    if not exists "payments" (
        id serial primary key,
        amount numeric(10, 2) not null,
        payment_date timestamp default current_timestamp,
        status varchar(50) not null,
        provider varchar(50) not null,
        created_at timestamp default current_timestamp,
        updated_at timestamp default current_timestamp,
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
        customer_id integer references customers (id) on delete cascade
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
