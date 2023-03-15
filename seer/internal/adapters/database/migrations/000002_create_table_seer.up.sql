create table seer (
  uuid uuid not null primary key default uuid_generate_v4(), 
  user_uuid uuid unique not null,
  onsite_available boolean not null default false,
  chat_available boolean not null default false,
  call_available boolean not null default false,
  video_call_available boolean not null default false,
  major varchar(50) not null,
  major_description varchar(255),
  description_profile varchar(255),
  map_coordinate varchar(255)
)