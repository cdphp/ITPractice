create database it_practice default charset=utf8;
  use it_practice;

/**
 * temp 示例
 * id 主键
 *
 * is_delete 是否删除
 * created_at 创建时间
 * updated_at 更新时间
 */
drop table if exists temp;
create table temp (
    id int not null primary key auto_increment,

    is_delete tinyint not null default 0,
    created_at int not null default 0,
    updated_at int not null default 0
)charset=utf8;

/**
 * users 用户表
 * id 主键
 * username 用户名
 * email 邮箱
 * password 密码
 * type 类型(1:普通用户,2:高级用户,3:管理员)
 * is_delete 是否删除
 * created_at 创建时间
 * updated_at 更新时间
 */
drop table if exists users;
create table users (
    id int not null primary key auto_increment,
    username varchar(30) not null,
    email varchar(50) not null,
    password char(32) not null,
    type tinyint not null default 1,
    is_delete tinyint not null default 0,
    created_at int not null default 0,
    updated_at int not null default 0
)charset=utf8;


/**
 * users_info 示例
 * id 主键
 *
 * is_delete 是否删除
 * created_at 创建时间
 * updated_at 更新时间
 */
drop table if exists users_info;
create table users_info (
    id int not null primary key auto_increment,
    user_id int not null,
    avatar varchar(100) not null default '',
    bg varchar(100) not null default '',
    about varchar(255) not null default '',
    labels varchar(255) not null default '',
    is_delete tinyint not null default 0,
    created_at int not null default 0,
    updated_at int not null default 0
)charset=utf8;

/**
 * tokens
 * id 主键
 * token 令牌
 * user_id 用户id
 * expire 过期时间
 * auth 用户权限
 * logout_at 登出时间
 * is_delete 是否删除
 * created_at 创建时间
 * updated_at 更新时间
 */
drop table if exists tokens;
 create table tokens (
   id int not null primary key auto_increment,
   token varchar(64) not null,
   user_id int not null,
   expire int not null,
   auth varchar(30) not null default '',
   logout_at int not null default 0,
   is_delete tinyint not null default 0,
   created_at int not null default 0,
   updated_at int not null default 0
 )charset=utf8;

 /**
  * articles 文章
  * id 主键
  * title 标题
  * digest 摘要
  * content
  * user_id 作者id
  * labels 标签
  * clicks 点击量
  * is_delete 是否删除
  * created_at 创建时间
  * updated_at 更新时间
  */
 drop table if exists articles;
 create table articles (
     id int not null primary key auto_increment,
     title varchar(30) not null,
     digest varchar(255) not null default '',
     content text not null,
     user_id int not null,
     labels varchar(255) not null default '',
     clicks int not null default 0,
     is_delete tinyint not null default 0,
     created_at int not null default 0,
     updated_at int not null default 0
 )charset=utf8;

 /**
  * comments 评论
  * id 主键
  * user_id 评论者id
  * type 类型(1:文章)
  * content 内容
  * target_id 目标id
  * root_id 回复id
  * is_delete 是否删除
  * created_at 创建时间
  * updated_at 更新时间
  */
 drop table if exists comments;
 create table comments (
     id int not null primary key auto_increment,
     user_id int not null,
     type tinyint not null default 1,
     content varchar(255) not null,
     target_id int not null,
     root_id int not null default 0,
     is_delete tinyint not null default 0,
     created_at int not null default 0,
     updated_at int not null default 0
 )charset=utf8;
