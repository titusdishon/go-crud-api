CREATE TABLE `users`
(
    id bigint auto_increment,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    summary varchar(255) NOT NULL,
    PRIMARY KEY(`id`)
);

INSERT INTO
 users(name, email, summary) 
 VALUES
 ('Test user', 'testemail@gmail.com', 'professional gopher-Golang'), 
 ('test user two', 'testuser2@gmail.com', 'professional script writter- Typescript');