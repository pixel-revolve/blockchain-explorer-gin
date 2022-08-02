create database if not exists `gin`;

use `gin`;

drop table if exists `sys_users`;

create table if not exists `sys_users`(
    `id` bigint primary key auto_increment,
    `uuid` binary(16),
    `username` varchar(30),
    `password` varchar(30),
    `nick_name` varchar(30),
    `phone` varchar(30),
    `email` varchar(40),
    `created_at` datetime,
    `updated_at` datetime,
    `deleted_at` datetime

)engine=Innodb;

