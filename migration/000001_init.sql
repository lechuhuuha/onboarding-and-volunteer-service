CREATE DATABASE IF NOT EXISTS `volunteer_manage`;

USE `volunteer_manage`;

CREATE TABLE `countries` (
  `country_id` int NOT NULL AUTO_INCREMENT,
  `country_name` varchar(45) NOT NULL,
  PRIMARY KEY (`country_id`)
);

CREATE TABLE `departments` (
  `department_id` int NOT NULL AUTO_INCREMENT,
  `department_name` varchar(45) NOT NULL,
  `location` varchar(100) NOT NULL,
  `create_date` datetime NOT NULL,
  `update_date` datetime DEFAULT NULL,
  `status` tinyint NOT NULL,
  PRIMARY KEY (`department_id`)
);

CREATE TABLE `requests` (
  `request_id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `request_type` varchar(45) NOT NULL,
  `request_status` tinyint NOT NULL,
  `reject_notes` varchar(500) DEFAULT NULL,
  `create_date` datetime NOT NULL,
  `update_date` datetime DEFAULT NULL,
  `censor_id` int DEFAULT NULL,
  PRIMARY KEY (`request_id`),
  KEY `fk_request_user_idx` (`user_id`),
  KEY `fk_request_censor_idx` (`censor_id`),
  CONSTRAINT `fk_request_censor` FOREIGN KEY (`censor_id`) REFERENCES `users` (`user_id`),
  CONSTRAINT `fk_request_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
);

CREATE TABLE `roles` (
  `role_id` int NOT NULL AUTO_INCREMENT,
  `role_name` varchar(45) NOT NULL,
  PRIMARY KEY (`role_id`)
);

CREATE TABLE `user_identities` (
  `ui_id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `number` varchar(20) NOT NULL,
  `type` varchar(45) NOT NULL,
  `status` tinyint NOT NULL,
  `date_expiry` datetime NOT NULL,
  `place_issue` varchar(200) NOT NULL,
  PRIMARY KEY (`ui_id`),
  KEY `fk_ui_users_idx` (`user_id`),
  CONSTRAINT `fk_ui_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
);

CREATE TABLE `users` (
  `user_id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `surname` varchar(45) NOT NULL,
  `email` varchar(45) NOT NULL,
  `mobile` varchar(45) NOT NULL,
  `role_id` int NOT NULL,
  `status` tinyint NOT NULL COMMENT '0: inactive\n1: active',
  `create_date` datetime NOT NULL,
  `update_date` datetime DEFAULT NULL,
  `department_id` int DEFAULT NULL,
  `sex` varchar(45) NOT NULL,
  `dob` datetime NOT NULL,
  `country_id` int NOT NULL,
  `resident_country` int NOT NULL,
  `avt_image` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  KEY `fk_users_roles_idx` (`role_id`),
  KEY `fk_users_depts_idx` (`department_id`),
  KEY `fk_users_countries_idx` (`country_id`),
  KEY `fk_users_resident_countries_idx` (`resident_country`),
  CONSTRAINT `fk_users_countries` FOREIGN KEY (`country_id`) REFERENCES `countries` (`country_id`),
  CONSTRAINT `fk_users_depts` FOREIGN KEY (`department_id`) REFERENCES `departments` (`department_id`),
  CONSTRAINT `fk_users_resident_countries` FOREIGN KEY (`resident_country`) REFERENCES `countries` (`country_id`),
  CONSTRAINT `fk_users_roles` FOREIGN KEY (`role_id`) REFERENCES `roles` (`role_id`)
);

