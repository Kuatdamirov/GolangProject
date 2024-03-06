create table articles(
    id serial primary key,
    title varchar(100),
    anons VARCHAR(255),
    fullText text
)