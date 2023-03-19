create table room_message(
    "uuid" uuid primary key default uuid_generate_v4(),
    "room_uuid" uuid not null references "room"(uuid),
    "user_uuid" uuid not null,
    "message" varchar(255) not null,
    "created_at" timestamptz not null default now()
)