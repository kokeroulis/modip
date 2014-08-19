CREATE SEQUENCE seq_bookIds;

CREATE TABLE book
(
    id                 int         primary key default nextval('seq_bookIds'),
    title              text        unique,
    teacher            int         not null references teacher(id) on delete cascade
);

