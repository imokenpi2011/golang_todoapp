drop table users;
drop table sessons;
drop table todos;

create table users (
  id            serial primary key,
  uuid          varchar(64) NOT  NULL UNIQUE,
  name          varchar(255),
  email         varchar(255),
  password      varchar(255),
  created_at    timestamp not null
);

create table sessions (
  id            serial primary key,
  uuid          varchar(64) NOT  NULL UNIQUE,
  email         varchar(255),
  user_id       integer references users(id),
  created_at    timestamp not null
);

create table todos (
  id            serial primary key,
  content       text,
  user_id       integer references users(id),
  created_at    timestamp not null
);