drop table if exists users;

create table users (
  id bigint unsigned not null auto_increment comment 'Primary key',
  name varchar(128) not null default '' comment 'Username',
  gender int(8) not null default 0 comment 'User gender',
  age int(64) not null default 0 comment 'User age',
  introduce text null comment 'User introduce',
  created_at timestamp not null default current_timestamp comment 'User information create time',
  updated_at timestamp not null default current_timestamp on update current_timestamp comment 'User information update time',
  deleted_at timestamp null default null comment 'User information delete time',
  primary key (id),
  key idx_name (name, deleted_at) comment 'Username index'
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_general_ci comment = 'User information table';
