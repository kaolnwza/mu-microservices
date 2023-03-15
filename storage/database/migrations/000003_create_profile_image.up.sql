create table profile_image (
    uuid uuid primary key default uuid_generate_v4(),
    upload_uuid uuid not null references "upload" (uuid),
    user_uuid uuid unique not null
)