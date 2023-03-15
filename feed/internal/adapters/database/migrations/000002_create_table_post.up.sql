


CREATE TABLE "post" (
  "uuid" uuid primary key default uuid_generate_v4(),
  "user_uuid" uuid not null,
  "title" varchar(50),
  "text" varchar(255),
  "created_at" timestamp with time zone not null default now(),
  "deleted_at" timestamp with time zone 
);

-- CREATE TABLE "feed_like" (
--   "uuid" primary key default uuid_generate_v4(),
--   "user_uuid" uuid not null,
--   "feed_uuid" uuid not null references "feed"(uuid),
--   unique("user_uuid", "feed_uuid")
-- );

-- CREATE TABLE "feed_comment" (
--   "uuid" primary key default uuid_generate_v4(),
--   "user_uuid" uuid not null,
--   "feed_uuid" uuid not null references "feed"(uuid),
--   "comment" varchar(255) not null
-- );

-- CREATE TABLE "feed_comment_like" (
--   "uuid" primary key default uuid_generate_v4(),
--   "user_uuid" uuid not null,
--   "comment_uuid" uuid not null references "feed_comment"(uuid),
--   unique("user_uuid", "comment_uuid")
-- );

