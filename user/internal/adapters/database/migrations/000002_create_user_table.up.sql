CREATE TYPE "user_role" AS ENUM ('user', 'seer');

create table "user" (
    uuid uuid primary key default uuid_generate_v4(),
    display_name varchar(100) not null,
    birthday date ,
    description varchar(255),
    tel_number varchar(10),
    role user_role default 'user'
);

