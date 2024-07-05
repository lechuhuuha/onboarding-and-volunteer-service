CREATE DATABASE IF NOT EXISTS `volunteer_manage`;

USE `volunteer_manage`;

CREATE TABLE `countries` (
  `id` int NOT NULL AUTO_INCREMENT,
  `country_name` varchar(45) NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `departments` (
  `id` int NOT NULL AUTO_INCREMENT,
  `department_name` varchar(45) NOT NULL,
  `location` varchar(100) NOT NULL,
  `create_at` datetime DEFAULT NULL,
  `update_at` datetime DEFAULT NULL,
  `status` tinyint NOT NULL,
  PRIMARY KEY (`id`)
);


CREATE TABLE `requests` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `request_type` varchar(45) NOT NULL,
  `request_status` tinyint NOT NULL DEFAULT '0',
  `reject_notes` varchar(500) DEFAULT NULL,
  `create_at` datetime DEFAULT NULL,
  `update_at` datetime DEFAULT NULL,
  `censor_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_request_user_idx` (`user_id`),
  KEY `fk_request_censor_idx` (`censor_id`),
  CONSTRAINT `fk_request_censor` FOREIGN KEY (`censor_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_request_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);


CREATE TABLE `roles` (
  `id` int NOT NULL AUTO_INCREMENT,
  `role_name` varchar(45) NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `user_identities` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `number` varchar(20) NOT NULL,
  `type` varchar(45) NOT NULL,
  `status` tinyint NOT NULL,
  `date_expiry` datetime NOT NULL,
  `place_issue` varchar(200) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_ui_users_idx` (`user_id`),
  CONSTRAINT `fk_ui_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);


CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `surname` varchar(45) NOT NULL,
  `email` varchar(45) NOT NULL,
  `mobile` varchar(45) NOT NULL,
  `role_id` int NOT NULL,
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '0: inactive\n1: active',
  `create_at` datetime DEFAULT NULL,
  `update_at` datetime DEFAULT NULL,
  `department_id` int DEFAULT NULL,
  `sex` varchar(45) NOT NULL,
  `dob` datetime NOT NULL,
  `country_id` int NOT NULL,
  `resident_country` int NOT NULL,
  `avt_image` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_users_roles_idx` (`role_id`),
  KEY `fk_users_depts_idx` (`department_id`),
  KEY `fk_users_countries_idx` (`country_id`),
  KEY `fk_users_resident_countries_idx` (`resident_country`),
  CONSTRAINT `fk_users_countries` FOREIGN KEY (`country_id`) REFERENCES `countries` (`id`),
  CONSTRAINT `fk_users_depts` FOREIGN KEY (`department_id`) REFERENCES `departments` (`id`),
  CONSTRAINT `fk_users_resident_countries` FOREIGN KEY (`resident_country`) REFERENCES `countries` (`id`),
  CONSTRAINT `fk_users_roles` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)
);


