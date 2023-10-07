create table books
( id serial primary key,
  author varchar not null,
     varchar not null,
  rating numeric not null default 0,
  publication_date int not null default 0,
);

insert into books (author, book_name, rating, publication_date) values ('Лев Толстой', 'Война и мир', 4.9, 1869);
insert into books (author, book_name, rating, publication_date) values ('Эрих Мария Ремарк', 'Три товарища', 4.8, 1936);
insert into books (author, book_name, rating, publication_date) values ('Эрнест Хемингуэй', 'Прощай, оружие!', 4.4, 1929);
