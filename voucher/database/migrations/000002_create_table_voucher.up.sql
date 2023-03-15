create type "discount_type" as enum ('percent', 'numberal');

create table voucher(
    uuid uuid primary key default uuid_generate_v4(),
    voucher_name varchar(255),
    voucher_code varchar(20) not null,
    discount_type discount_type not null default 'numberal',
    discount int not null,
    max_discount int,
    voucher_quantity int not null default 0,
    voucher_status boolean not null default false,
    created_at timestamp with time zone not null default now(),
    expired_at timestamp with time zone not null 
);