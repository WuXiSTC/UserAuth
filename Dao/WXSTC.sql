DROP DATABASE IF EXISTS WXSTC;
CREATE DATABASE WXSTC DEFAULT CHARACTER SET utf8;
USE WXSTC;
DROP USER IF EXISTS WXSTC@'%';
CREATE USER WXSTC@'%' IDENTIFIED WITH mysql_native_password BY 'WXSTC' PASSWORD EXPIRE NEVER;
GRANT SELECT, UPDATE, INSERT, DELETE, CREATE, DROP on WXSTC.* to WXSTC@'%';
FLUSH PRIVILEGES;
USE WXSTC;
DROP TABLE IF EXISTS Users;
CREATE TABLE Users
(
    ID   varchar(255) primary key,
    PASS varchar(255) not null
);