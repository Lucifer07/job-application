create type roles as enum ('admin','user');
create table users(
	id bigserial primary key,
	username varchar not null,
	email varchar not null,
	password varchar not null,
	role roles default 'user',
	created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp null,
	UNIQUE(email)
);
create table jobs(
	id bigserial primary key,
	name varchar not null,
	quota bigint not null,
	expired_date timestamp not null,
	created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp null
);
create type statuses as enum ('applied','rejected','accepted');
create table aplication_record(
	id bigserial primary key,
	user_id bigint not null,
	job_id bigint not null,
	status statuses default 'applied',
	created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp null,
	foreign key (user_id)references users(id),
	foreign key(job_id)references jobs(id)
);
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;

select * from jobs;
select id,name,quota,expired_date from jobs where deleted_at is null and id =2;
select * from aplication_record;
