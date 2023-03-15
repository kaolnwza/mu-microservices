create table room_user (
    "uuid" uuid primary key default uuid_generate_v4(),
    "room_uuid" uuid not null references "room"(uuid),
    "user_uuid" uuid not null
)