drop database `Miiko`;
create database if not exists `Miiko` default character set utf8mb4;
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
) engine=InnoDB default charset=utf8mb4;

-- Presentation Channel
create table if not exists `channel-presentation` (
	`server` varchar(32) primary key,
	`channel` varchar(32) not null
) engine=InnoDB default charset=utf8mb4;

-- Messages

-- Pins
create table if not exists `pins` (
	`server` varchar(32) not null,
	`message` varchar(32) primary key,
	`member` varchar(32) not null
) engine=InnoDB default charset=utf8mb4;

-- Minimum Reactions
create table if not exists `minimum-reactions` (
	`channel` varchar(32) primary key,
	`minimum` int not null
) engine=InnoDB default charset=utf8mb4;

-- Roles

-- Administrator
create table if not exists `role-administrator` (
	`server` varchar(32) primary key,
	`role` varchar(32) not null
) engine=InnoDB default charset=utf8mb4;

-- Moderator
create table if not exists `role-moderator` (
	`server` varchar(32) primary key,
	`role` varchar(32) not null
) engine=InnoDB default charset=utf8mb4;

-- Light
create table if not exists `role-light` (
	`server` varchar(32) primary key,
	`role` varchar(32) not null
) engine=InnoDB default charset=utf8mb4;

-- Absynthe
create table if not exists `role-absynthe` (
	`server` varchar(32) primary key,
	`role` varchar(32) not null
) engine=InnoDB default charset=utf8mb4;

-- Obsidian
create table if not exists `role-obsidian` (
	`server` varchar(32) primary key,
	`role` varchar(32) not null
) engine=InnoDB default charset=utf8mb4;

-- Shadow
create table if not exists `role-shadow` (
	`server` varchar(32) primary key,
	`role` varchar(32) not null
) engine=InnoDB default charset=utf8mb4;

-- Guardless
create table if not exists `role-eel` (
	`server` varchar(32) primary key,
	`role` varchar(32) not null
) engine=InnoDB default charset=utf8mb4;

-- Non-Playing Characters
create table if not exists `role-npc` (
	`server` varchar(32) primary key,
	`role` varchar(32) not null
) engine=InnoDB default charset=utf8mb4;

-- Self-Assignable Roles
create table if not exists `roles-sar` (
	`server` varchar(32) not null,
	`role` varchar(32) not null,
	constraint `pk_roles_saf` primary key (`server`, `role`)
) engine=InnoDB default charset=utf8mb4;

-- Views

-- drop view `love`;
-- drop view `pins-count`;

-- Pins Count
-- create view `pins-count` as
-- select `server`, `member`, count(`message`) as `count`
-- from `pins`
-- group by `server`, `member`
-- order by `server`, `count` desc
-- ;
