CREATE TABLE "feed_image" (
  "uuid" uuid primary key default uuid_generate_v4(),
  "upload_uuid" uuid not null,
  "feed_uuid" uuid not null,
  "image_order" int not null default 1
);