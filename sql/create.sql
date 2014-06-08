CREATE DATABASE bank;
USE bank;

CREATE TABLE `Transaction` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `amount` int(11) DEFAULT NULL,
    `time` date DEFAULT NULL,
    `note` VARCHAR(255) DEFAULT NULL,
    `deletionTime` DATETIME DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;

GRANT SELECT, INSERT, UPDATE, DELETE ON bank.* TO 'bank'@'localhost' IDENTIFIED BY 'bank';
