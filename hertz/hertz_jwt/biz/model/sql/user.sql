drop table if exists users;

create table users (
  id bigint unsigned not null auto_increment comment 'Primary key',
  username varchar(128) not null default '' comment 'Username',
  email varchar(128) not null default '' comment 'User email',
  password varchar(128) not null default '' comment 'User password',
  created_at timestamp not null default current_timestamp comment 'User information create time',
  updated_at timestamp not null default current_timestamp on update current_timestamp comment 'User information update time',
  deleted_at timestamp null default null comment 'User information delete time',
  primary key (id),
  unique key idx_username (username) comment 'Username index',
  unique key idx_email (email) comment 'User email index'
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_general_ci comment = 'User information table';
