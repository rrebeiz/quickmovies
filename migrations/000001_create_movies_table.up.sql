create table if not exists movies (
    id bigserial primary key,
    title text not null,
    year integer not null,
    runtime integer not null,
    genres text[] not null,
    version integer not null default 1,
    created_at timestamp(0) without time zone not null default now(),
    updated_at timestamp(0) without time zone not null default now()
)