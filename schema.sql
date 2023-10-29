SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

CREATE TABLE `album` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `title` varchar(100) NOT NULL,
                         `artist` varchar(100) NOT NULL,
                         `price` decimal(10,2) NOT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;