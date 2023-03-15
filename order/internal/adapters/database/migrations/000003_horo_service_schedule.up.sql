create table horo_service_schedule (
    uuid uuid not null primary key default uuid_generate_v4(), 
    horo_service_uuid uuid not null references "horo_service"(uuid),
    start_time timestamp with time zone not null,
    end_time timestamp with time zone not null
)