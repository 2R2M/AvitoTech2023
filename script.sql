create table "user"
(
    id varchar(255) not null
        constraint user_id
            primary key
);

alter table "user"
    owner to postgres;

create table usersegment
(
    user_id    varchar(255) not null,
    slug       varchar(255) not null,
    expired_at timestamp,
    constraint usersegment_pk
        unique (user_id, slug, expired_at)
);

alter table usersegment
    owner to postgres;

create table segment
(
    slug    varchar(255) not null
        constraint segment_pk
            primary key,
    percent double precision
);

alter table segment
    owner to postgres;

create table operation
(
    id             varchar(255) not null
        constraint operation_pk
            primary key,
    user_id        varchar(255) not null,
    segment_id     varchar(255) not null,
    operation_type varchar(255) not null,
    created_at     timestamp default CURRENT_TIMESTAMP,
    expires_at     timestamp,
    status         varchar(255) not null
);

alter table operation
    owner to postgres;

create index operation_id_index
    on operation (status);

create index operation_operation_type_index
    on operation (operation_type);

