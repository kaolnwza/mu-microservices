create table user_wallet (
    uuid uuid not null default uuid_generate_v4(),
    user_uuid uuid not null unique,
    fund int not null default 0
)