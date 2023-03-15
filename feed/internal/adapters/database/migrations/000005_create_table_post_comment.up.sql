create table post_comment (
    "uuid" uuid primary key default uuid_generate_v4(),
    "post_uuid" uuid not null references "post" (uuid),
    "user_uuid" uuid not null,
    "comment" varchar(255) not null,
    "created_at" timestamp with time zone not null default now(),
    "deleted_at"  timestamp with time zone 
)