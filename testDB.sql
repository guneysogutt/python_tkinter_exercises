CREATE TABLE IF NOT EXISTS `users` (
	`id` int(10) NOT NULL auto_increment,
	`username` varchar(255) NOT NULL,
	`password` varchar(255) NOT NULL,
	PRIMARY KEY( `id` )
);


CREATE TABLE IF NOT EXISTS `items` (
	`id` int(10) NOT NULL auto_increment,
	`name` varchar(255) NOT NULL,
	`quantity` numeric(9,2),
	PRIMARY KEY( `id` )
);

INSERT INTO users (username,password)
VALUES ('guneys', 'a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3');