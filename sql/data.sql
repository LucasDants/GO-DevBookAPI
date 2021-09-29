insert into users (name, nick, email, password)
values
("Lucas", "lucas", "lucas@gmail.com", "$2a$10$oLElq5t7ZiFCAzU6.JIFUuMg5z1reoudjf9GVN.Ntyo6ZDeJn2Fna"),
("Gabriel", "gabriel", "gabriel@gmail.com", "$2a$10$oLElq5t7ZiFCAzU6.JIFUuMg5z1reoudjf9GVN.Ntyo6ZDeJn2Fna");

insert into followers(userID, followerID)
values
(1, 2),
(2, 1);

insert into publications(title, content, author_id)
values
("Primeira", "Publicacao", 1),
("Segunda", 'Publicacao', 2);