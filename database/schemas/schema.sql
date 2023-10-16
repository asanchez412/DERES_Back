CREATE DATABASE IF NOT EXISTS `deres`;

USE `deres`;  
CREATE TABLE IF NOT EXISTS `users` (
  `user_name` varchar(63) NOT NULL,
  `password` varchar(63) NOT NULL,
  PRIMARY KEY (`user_name`)
);

INSERT INTO `user` (`user_name`, `password`) VALUES
('admin', 'admin'),
('user', 'user');