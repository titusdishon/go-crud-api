CREATE TABLE `users`
(
    id bigint auto_increment,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
    summary varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    PRIMARY KEY(`id`)
);

INSERT INTO
 users(`name`, `email`, `summary`) 
 VALUES
 ('Test user one', 'testemail1@gmail.com', 'professional gopher-Golang', ''), 
 ('Test user two', 'testemail2@gmail.com', 'professional frontend developer', ''), 
 ('test user three', 'testuser3@gmail.com', 'professional script writter- Typescript', '');