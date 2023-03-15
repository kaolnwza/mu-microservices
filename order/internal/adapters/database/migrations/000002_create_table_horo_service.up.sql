CREATE TYPE "horo_location" AS ENUM ('hybrid', 'online', 'onsite');

CREATE TYPE "horo_type" AS ENUM ('gypsy_card', 'astrology', 'physiognomy', 'taro_card', 'custom');
-- ยิปซี, โหรา, โหงวเฮ้ง, ทาโร่

create table horo_service(
    uuid uuid not null primary key default uuid_generate_v4(), 
    seer_uuid uuid not null,
    horo_location horo_location not null,
    horo_type horo_type not null default 'custom',
    title varchar(255),
    description varchar(255),
    price int not null,
    available boolean not null default true,
    meeting_status boolean not null default false,
    meeting_location varchar(255),
    chat_status boolean not null default false,
    voice_call_status boolean not null default false,
    video_call_status boolean not null default false
)