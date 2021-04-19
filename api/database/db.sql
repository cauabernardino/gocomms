CREATE DATABASE IF NOT EXISTS gocomms;
USE gocomms;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id int auto_increment primary key,
    name varchar(100) not null,
    username varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(100) not null,
    created_on timestamp default current_timestamp()
) ENGINE=INNODB;