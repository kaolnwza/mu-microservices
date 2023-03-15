create table room (
    "uuid" uuid primary key default uuid_generate_v4(),
    "horo_order_uuid" uuid unique not null,
    "start_time" timestamp with time zone not null,
    "end_time" timestamp with time zone not null,
    "created_at" timestamp with time zone not null default now(),
    "status" boolean not null default true
)