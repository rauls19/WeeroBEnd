CREATE TABLE users (
    id          bigint NOT NULL AUTO_INCREMENT,
    hashid         varchar(300) UNIQUE,
    name        varchar(50),
    surname     varchar(150),
    birthday    date,
	email       varchar(50),
    password        varchar(255),
	Age         integer,
	Interested  bit,
	Location    varchar(255),
	Description varchar(255),
	Mobilephone int UNIQUE,
    PRIMARY KEY(id)
);