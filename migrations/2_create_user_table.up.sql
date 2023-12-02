CREATE TABLE users
(
    id                     varchar(36) primary key not null,
    name                   varchar(255) not null ,
    created_at             timestamp with time zone,
    updated_at             timestamp with time zone,
    deleted_at             timestamp with time zone
);