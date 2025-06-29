set
  names utf8mb4;

set
  foreign_key_checks = 0;

-- 禁用外键约束检查
drop table if exists comments;

create table comments (
  id bigint not null auto_increment comment '评论 ID, 自增主键',
  user_id bigint not null comment '发布评论的用户 ID',
  video_id bigint not null comment '评论的视频 ID',
  comment_text varchar(255) not null comment '评论的内容',
  created_at timestamp not null default current_timestamp comment '评论的创建时间',
  deleted_at timestamp null default null comment '评论的删除时间',
  primary key (id),
  key idx_video_id (video_id) using btree comment '评论的视频 ID 的索引'
) engine = InnoDB
-- 自增主键的起始值为 1000
auto_increment = 1000 default charset = utf8 comment = '评论表';

drop table if exists follows;

create table follows (
  id bigint not null auto_increment comment '自增主键',
  user_id bigint not null comment '用户 ID',
  follower_id bigint not null comment '粉丝 ID',
  created_at timestamp not null default current_timestamp comment '关注关系的创建时间',
  deleted_at timestamp null default null comment '关注关系的删除时间',
  primary key (id),
  key idx_user_id_follower_id (user_id, follower_id) using btree,
  key idx_follower_id (follower_id) using btree,
  key idx_user_id (user_id) using btree
) engine = InnoDB auto_increment = 1000 default charset = utf8 comment = '关注表';

drop table if exists likes;

create table likes (
  id bigint not null auto_increment comment '自增主键',
  user_id bigint not null comment '点赞的用户 ID',
  video_id bigint not null comment '点赞的视频 ID',
  created_at timestamp not null default current_timestamp comment '点赞的创建时间',
  deleted_at timestamp not null default null comment '点赞的删除时间',
  primary key (id),
  key idx_user_id_video_id (user_id, video_id) using btree,
  key idx_user_id (user_id) using btree,
  key idx_video_id (video_id) using btree
) engine = InnoDB auto_increment = 1000 default charset = utf8 comment = '点赞表';

drop table if exists users;

create table users (
  id bigint not null auto_increment comment '用户 ID, 自增主键',
  username varchar(255) not null comment '用户名',
  password varchar(255) not null comment '用户密码',
  avatar varchar(255) not null comment '用户头像',
  background_image varchar(255) not null comment '用户背景',
  signature varchar(255) not null comment '用户签名',
  primary key (id),
  key idx_username_password (username, password) using btree
) engine = InnoDB auto_increment = 1000 default charset = utf8 comment = '用户表';

drop table if exists videos;

create table videos (
  id bigint not null auto_increment comment '视频 ID, 自增主键',
  author_id bigint not null comment '视频作者 ID',
  play_url varchar(255) not null comment '播放 url',
  cover_url varchar(255) not null comment '封面 url',
  publish_time varchar(255) default null comment '发布时间戳',
  title varchar(255) default null comment '视频名称',
  primary key (id),
  key idx_publish_time (publish_time) using btree,
  key idx_author_id (author_id) using btree
) engine = InnoDB auto_increment = 1000 default charset = utf8 comment = '视频表';

drop table if exists messages;

create table messages (
  id bigint not null auto_increment comment '消息 ID, 自增主键',
  rx_user_id bigint not null comment '消息接收者的用户 ID',
  tx_user_id bigint not null comment '消息发送者的用户 ID',
  content text not null comment '消息内容',
  created_at datetime (6) not null comment '消息创建时间',
  primary key (id),
  key idx_created_at (created_at) using btree,
  key idx_rx_user_id_tx_user_id (rx_user_id, tx_user_id) using btree
) engine = InnoDB auto_increment = 1000 default charset = utf8 comment = '消息表';

set
  foreign_key_checks = 1;
