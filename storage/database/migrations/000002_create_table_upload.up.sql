create table upload (
    uuid uuid primary key default uuid_generate_v4(),
    user_uuid uuid,
    bucket varchar(50) not null,
    path varchar(255) not null,
    created_at timestamp with time zone not null default now(),
    deleted_at timestamp with time zone
)