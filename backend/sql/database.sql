-- CREATE TABLE USERS(
-- 	id int primary key not null AUTO_INCREMENT,
-- 	email varchar(50) unique not null,
-- 	password varchar(255) not null,
-- 	username varchar(50),
-- 	firstname varchar(50),
-- 	lastname varchar(50),
-- 	address varchar(100),
-- 	state varchar(50),
-- 	zipcode varchar(5)
-- );

-- CREATE TABLE producttype (
--     producttype_id INT NOT NULL,
--     name VARCHAR(255),
--     PRIMARY KEY (producttype_id)
-- );

-- CREATE TABLE product (
--     product_id INT NOT NULL AUTO_INCREMENT,
--     name VARCHAR(255),
--     price FLOAT,
--     detail VARCHAR(255),
--     product_img LONGBLOB,
--     producttype_id INT,
--     PRIMARY KEY (product_id)
--     FOREIGN KEY (producttype_id) REFERENCES producttype(producttype_id)
-- );
