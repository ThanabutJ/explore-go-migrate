CREATE TABLE IF NOT EXISTS users(
    user_id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR (50) NOT NULL,
    password VARCHAR (50) NOT NULL,
    email VARCHAR (300) NOT NULL,
    PRIMARY KEY(user_id)
);
