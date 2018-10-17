drop database `Miiko`;
create database if not exists `Miiko` default character set utf8mb4 collate=utf8mb4_unicode_ci;
use `Miiko`;

-- Drop Table
drop table if exists `pins`;
drop table if exists `minimum-reactions`;
drop table if exists `channel-welcome`;
drop table if exists `role-absynthe`;
drop table if exists `role-shadow`;
drop table if exists `role-obsidian`;
drop table if exists `role-moderator`;
drop table if exists `role-administrator`;
drop table if exists `roles-sar`;

-- Legacy tables
drop table if exists `welcome`;

-- Create Table

-- Channels

-- Welcome Channel
create table if not exists `channel-welcome` (
	`server` varchar(32) primary key,
	`channel` varchar(32) not null
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_unicode_ci;

-- Presentation Channel
create table if not exists `channel-presentation` (
	`server` varchar(32) primary key,
	`channel` varchar(32) not null
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_unicode_ci;

-- Messages

-- Pins
create table if not exists `pins` (
	`server` varchar(32) not null,
	`message` varchar(32) primary key,
	`member` varchar(32) not null
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_unicode_ci;

-- Minimum Reactions
create table if not exists `minimum-reactions` (
	`channel` varchar(32) primary key,
	`minimum` int not null
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_unicode_ci;

-- Roles

-- Administrator
create table if not exists `role-administrator` (
	`server` varchar(32) primary key,
	`role` varchar(32) not null
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_unicode_ci;

-- Moderator
create table if not exists `role-moderator` (
	`server` varchar(32) primary key,
	`role` varchar(32) not null
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_unicode_ci;

-- Light
create table if not exists `role-light` (
	`server` varchar(32) primary key,
	`role` varchar(32) not null
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_unicode_ci;

-- Absynthe
create table if not exists `role-absynthe` (
	`server` varchar(32) primary key,
	`role` varchar(32) not null
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_unicode_ci;

-- Obsidian
create table if not exists `role-obsidian` (
	`server` varchar(32) primary key,
	`role` varchar(32) not null
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_unicode_ci;

-- Shadow
create table if not exists `role-shadow` (
	`server` varchar(32) primary key,
	`role` varchar(32) not null
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_unicode_ci;

-- Guardless
create table if not exists `role-eel` (
	`server` varchar(32) primary key,
	`role` varchar(32) not null
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_unicode_ci;

-- Non-Playing Characters
create table if not exists `role-npc` (
	`server` varchar(32) primary key,
	`role` varchar(32) not null
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_unicode_ci;

-- Self-Assignable Roles
create table if not exists `roles-sar` (
	`server` varchar(32) not null,
	`role` varchar(32) not null,
	constraint `pk_roles_sar` primary key (`server`, `role`)
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_unicode_ci;

-- Status
create table if not exists `status` (
	`id` int auto_increment primary key,
	`status` varchar(32) not null
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_unicode_ci;

-- Views

-- drop view `love`;
-- drop view `pins-count`;

-- Pins Count
create or replace view `pins-count` as
	select `server`, `member`, count(`message`) as `count`
		from `pins`
		group by `server`, `member` 
		order by `server` asc,`count` desc, `member` asc
;
