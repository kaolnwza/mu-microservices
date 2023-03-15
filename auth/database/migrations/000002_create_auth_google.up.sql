create table "auth_google" (
    uuid uuid primary key default uuid_generate_v4(),
    user_uuid uuid unique not null,
    email varchar(100) unique not null ,
    created_at timestamp with time zone not null default now()
)