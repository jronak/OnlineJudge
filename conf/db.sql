
CREATE DATABASE IF NOT EXISTS  OnlineJudge;

USE OnlineJudge;

CREATE TABLE IF NOT EXISTS `user` (
	`uid` int AUTO_INCREMENT PRIMARY KEY,
	`username` varchar(15) NOT NULL UNIQUE,
	`password` varchar(100) NOT NULL,
	`name` varchar(20) NOT NULL,
	`college` varchar(30) DEFAULT "",
	`email` varchar(30) NOT NULL UNIQUE,
	`score` smallint DEFAULT 0,
	`rank`  smallint DEFAULT 0 ,
	`is_editor` smallint DEFAULT 0
)ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS `problem`(
	`pid` int AUTO_INCREMENT PRIMARY KEY,
	`uid` int NOT NULL,
	`statement` text NOT NULL,
	`description` text NOT NULL,
	`constraints` text NOT NULL,
	`sample_input` text NOT NULL,
	`sample_output` text NOT NULL,
	`solution_description` mediumtext DEFAULT "",
	`solution_code` mediumtext DEFAULT "",
	`type` varchar(30) NOT NULL,
	`difficulty` varchar(30) NOT NULL,
	`created_at` datetime NOT NULL,
	`points` smallint NOT NULL,
	`solve_count` int NOT NULL,
	 FOREIGN KEY `fk_user`(`uid`)
	 REFERENCES `user`(`uid`)
	 ON UPDATE CASCADE
	 ON DELETE CASCADE  
)ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS `testcases`(
	`id` int AUTO_INCREMENT PRIMARY KEY,
	`pid` int NOT NULL ,
	`tid` int NOT NULL ,
	`input` longtext NOT NULL,
	`output` longtext NOT NULL,
	`timeout` tinyint NOT NULL,
	 FOREIGN KEY `fk_problem`(`pid`)
	 REFERENCES `problem`(`pid`)
	 ON UPDATE CASCADE
	 ON DELETE CASCADE
)ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS `problemlogs`(
	`id` int AUTO_INCREMENT PRIMARY KEY,
	`pid` int NOT NULL ,
	`uid` int NOT NULL ,
	`solved` int NOT NULL,
	`points` smallint NOT NULL,
	`time` datetime NOT NULL,
	 FOREIGN KEY `fk_user`(`uid`)
	 REFERENCES `user`(`uid`)
	 ON UPDATE CASCADE
	 ON DELETE CASCADE,
	 FOREIGN KEY `fk_problem`(`pid`)
	 REFERENCES `problem`(`pid`)
	 ON UPDATE CASCADE
	 ON DELETE CASCADE    
)ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS `contest`(
	`id` int AUTO_INCREMENT PRIMARY KEY,
	`name` varchar(30) NOT NULL,
	`description` mediumtext NOT NULL,
	`startTime` datetime NOT NULL,
	`endTime` datetime NOT NULL
)ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS `contestlogs`(
	`id` int AUTO_INCREMENT PRIMARY KEY,
	`uid` int NOT NULL,
	`cid` int NOT NULL,
	`points` int NOT NULL,
	`time` datetime NOT NULL,
	FOREIGN KEY `fk_user`(`uid`)
	REFERENCES `user`(`uid`)
	ON UPDATE CASCADE
	ON DELETE CASCADE,
	FOREIGN KEY `fk_contest`(`cid`)
	REFERENCES `contest`(`id`)
	ON UPDATE CASCADE
	ON DELETE CASCADE
)ENGINE=InnoDB;


