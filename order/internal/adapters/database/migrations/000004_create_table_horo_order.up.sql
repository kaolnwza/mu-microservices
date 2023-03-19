create type horo_order_status as enum ('fail', 'unpaid', 'paid', 'cancel_by_user', 'cancel_by_seer', 'success', 'inprogress', 'refund', 'confirmed');

create table horo_order (
    uuid uuid not null primary key default uuid_generate_v4(), 
    user_uuid uuid not null,
    horo_service_uuid uuid not null references "horo_service"(uuid),
    voucher_uuid uuid,
    payment_uuid uuid not null,
    price int not null,
    status horo_order_status not null default 'unpaid',
    start_time timestamp with time zone not null,
    end_time timestamp with time zone not null
)