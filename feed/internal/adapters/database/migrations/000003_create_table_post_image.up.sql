CREATE TABLE "post_image" (
    "uuid" uuid primary key default uuid_generate_v4(),
    "post_uuid" uuid not null references "post" (uuid),
    "upload_uuid" uuid not null,
    "order" int not null,
    "created_at" timestamp with time zone not null default now(),
    "deleted_at" timestamp with time zone 
);