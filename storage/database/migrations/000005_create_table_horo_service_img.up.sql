create table horo_service_img (
    uuid uuid primary key default uuid_generate_v4(),
    horo_service_uuid uuid not null,
    upload_uuid uuid not null,
    image_order int not null default 1
)