CREATE TABLE user (
    id integer not null primary key AUTOINCREMENT,
    username varchar(100) not null,
    email varchar(100) not null,
    firstName varchar(50) not null,
    lastName varchar(50) not null,
    joinedTs datetime not null,
    status integer not null default 0
);