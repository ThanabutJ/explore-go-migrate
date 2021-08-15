ALTER TABLE `users` 
ADD COLUMN `first_name` VARCHAR(45) NULL AFTER `email`,
ADD COLUMN `last_name` VARCHAR(45) NULL AFTER `first_name`,
ADD UNIQUE INDEX `username_UNIQUE` (`username` ASC) VISIBLE,
ADD UNIQUE INDEX `email_UNIQUE` (`email` ASC) VISIBLE;
;

