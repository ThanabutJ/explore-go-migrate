ALTER TABLE `users` 
DROP COLUMN `last_name`,
DROP COLUMN `first_name`,
DROP INDEX `email_UNIQUE` ,
DROP INDEX `username_UNIQUE` ;
;
