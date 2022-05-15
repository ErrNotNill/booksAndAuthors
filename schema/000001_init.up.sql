CREATE TABLE books
(
    id          int       not null unique,
    title         varchar(255) not null,
);

CREATE TABLE authors
(
    id          int      not null unique,
    firstname       varchar(255) not null,
    lastname varchar(255) not null,
    patronymic varchar(255) not null,
);