CREATE TABLE "post_like" (
  "uuid" uuid primary key default uuid_generate_v4(),
  "user_uuid" uuid not null,
  "post_uuid" uuid not null references "post"(uuid),
  unique("user_uuid", "post_uuid")
);
