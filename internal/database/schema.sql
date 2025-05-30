-- Create tables
CREATE TABLE `images` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`name` varchar(100),
	`format` varchar(100),
	`url` varchar(256) DEFAULT NULL,
	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB AUTO_INCREMENT = 6 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE casinos (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`name` varchar(100),
	`location` varchar(256),
	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	UNIQUE KEY `id` (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 2 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE `parameters` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`lotery_name` varchar(100),
	`date_start` varchar(100),
	`date_end` varchar(100),
	`status` enum ('scan', 'draw') DEFAULT ('scan'),
	`client_data` tinyint (1) DEFAULT '0',
	`home_page` bigint unsigned DEFAULT NULL,
	`client_page` bigint unsigned DEFAULT NULL,
	`result_page` bigint unsigned DEFAULT NULL,
	`general_rules` text,
	`secret` varchar(256),
	`secret_length` int,
	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	UNIQUE KEY `id` (`id`),
	KEY `parameters_home_page_images_id_fk` (`home_page`),
	KEY `parameters_client_page_images_id_fk` (`client_page`),
	KEY `parameters_result_page_images_id_fk` (`result_page`)
) ENGINE = InnoDB AUTO_INCREMENT = 2 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE `kiosks` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`name` varchar(100),
	`macadress_wifi` varchar(100),
	`macadress_ethernet` varchar(100),
	`location` varchar(256),
	`id_casino` bigint unsigned DEFAULT NULL,
	`id_parameter` bigint unsigned DEFAULT NULL,
	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	UNIQUE KEY `id` (`id`),
	KEY `kiosks_id_parameters_parameters_id_fk` (`id_parameter`),
	KEY `kiosks_id_casino_casinos_id_fk` (`id_casino`)
) ENGINE = InnoDB AUTO_INCREMENT = 3 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE `publicity_images` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`id_parameter` bigint unsigned DEFAULT NULL,
	`kiosk_id` bigint unsigned DEFAULT NULL,
	`image_id` bigint unsigned DEFAULT NULL,
	PRIMARY KEY (`id`),
	UNIQUE KEY `id` (`id`),
	KEY `publicity_images_id_parameter_id_parameters_fk` (`id_parameter`),
	KEY `publicity_images_kiosk_id_kiosks_id_fk` (`kiosk_id`),
	KEY `publicity_images_image_id_images_id_fk` (`image_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 3 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE `specific_rules` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`id_casino` bigint unsigned DEFAULT NULL,
	`specific_rule` text NOT NULL,
	PRIMARY KEY (`id`),
	UNIQUE KEY `id` (`id`),
	KEY `specific_rules_id_casino_casinos_id_fk` (`id_casino`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE `rewards` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`name` varchar(100),
	`big_win` tinyint (1) DEFAULT '0',
	`id_images` bigint unsigned DEFAULT NULL,
	`id_casino` bigint unsigned DEFAULT NULL,
	`id_parameter` bigint unsigned DEFAULT NULL,
	PRIMARY KEY (`id`),
	UNIQUE KEY `id` (`id`),
	KEY `rewards_id_images_images_id_fk` (`id_images`),
	KEY `rewards_id_casino_kiosks_id_fk` (`id_casino`),
	KEY `rewards_id_parameter_parameters_id_fk` (`id_parameter`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE `tickets` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`id_casino` bigint unsigned DEFAULT NULL,
	`id_reward` bigint unsigned DEFAULT NULL,
	`ticket_number` varchar(256) NOT NULL,
	`client_phone` varchar(100) DEFAULT NULL,
	`claim` tinyint (1) DEFAULT '0',
	`entry_scan` timestamp DEFAULT CURRENT_TIMESTAMP,
	`exit_scan` timestamp DEFAULT NULL,
	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	UNIQUE KEY `id` (`id`),
	UNIQUE KEY `tickets_ticket_number_unique` (`ticket_number`),
	KEY `ticket_number_idx` (`ticket_number`),
	KEY `tickets_id_casino_casinos_id_fk` (`id_casino`),
	KEY `tickets_id_reward_rewards_id_fk` (`id_reward`),
	KEY `idx_tickets_id_reward` (`id_reward`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

-- Create view
CREATE VIEW `kiosk_view` AS
SELECT
	`kiosks`.`id`,
	`parameters`.`id` AS `id_parameter`,
	`casinos`.`id` AS `id_casino`,
	`kiosks`.`name`,
	`kiosks`.`macadress_wifi`,
	`kiosks`.`macadress_ethernet`,
	`kiosks`.`location`,
	`casinos`.`name` AS `casino_name`,
	`casinos`.`location` AS `casino_location`,
	`parameters`.`lotery_name`,
	`parameters`.`date_start`,
	`parameters`.`date_end`,
	`parameters`.`status`,
	`parameters`.`client_data`,
	`parameters`.`general_rules`,
	`specific_rules`.`specific_rule` AS `specific_rule`,
	`parameters`.`home_page` AS `home_page`,
	`parameters`.`client_page` AS `client_page`,
	`parameters`.`result_page` AS `result_page`,
	`parameters`.`secret`,
	`parameters`.`secret_length`,
	`kiosks`.`updated_at`,
	`parameters`.`updated_at` AS `updated_at_parameters`
FROM
	`kiosks`
	LEFT JOIN `parameters` ON `parameters`.`id` = `kiosks`.`id_parameter`
	LEFT JOIN `casinos` ON `casinos`.`id` = `kiosks`.`id_casino`
	LEFT JOIN `specific_rules` ON `specific_rules`.`id_casino` = `casinos`.`id`;

-- Create reward view for dl the reward in the db
CREATE VIEW `reward_view` AS
SELECT
	`rewards`.`id` AS `reward_id`,
	`rewards`.`name` AS `reward_name`,
	`rewards`.`big_win`,
	`images`.`id` AS `image_id`,
	`images`.`name` AS `image_name`,
	`images`.`format` AS `image_format`,
	`images`.`url` AS `image_url`,
	`casinos`.`id` AS `casino_id`,
	`rewards`.id_parameter
FROM
	`tickets`
	LEFT JOIN `casinos` ON `casinos`.`id` = `tickets`.`id_casino`
	LEFT JOIN `rewards` ON `tickets`.`id_reward` = `rewards`.`id`
	LEFT JOIN `images` ON `rewards`.`id_images` = `images`.`id`
WHERE
	`tickets`.`id_reward` IS NOT NULL;

-- Filter to only include tickets with rewards
-- Add foreign key constraints
ALTER TABLE `publicity_images` ADD CONSTRAINT `publicity_images_image_id_images_id_fk` FOREIGN KEY (`image_id`) REFERENCES `images` (`id`);

ALTER TABLE `publicity_images` ADD CONSTRAINT `publicity_images_parameter_id_parameters_id_fk` FOREIGN KEY (`id_parameter`) REFERENCES `parameters` (`id`);

ALTER TABLE `publicity_images` ADD CONSTRAINT `publicity_images_kiosk_id_kiosks_id_fk` FOREIGN KEY (`kiosk_id`) REFERENCES `kiosks` (`id`);

ALTER TABLE `rewards` ADD CONSTRAINT `rewards_id_images_images_id_fk` FOREIGN KEY (`id_images`) REFERENCES `images` (`id`);

ALTER TABLE `parameters` ADD CONSTRAINT `parameters_home_page_images_id_fk` FOREIGN KEY (`home_page`) REFERENCES `images` (`id`);

ALTER TABLE `parameters` ADD CONSTRAINT `parameters_result_page_images_id_fk` FOREIGN KEY (`result_page`) REFERENCES `images` (`id`);

ALTER TABLE `parameters` ADD CONSTRAINT `parameters_client_page_images_id_fk` FOREIGN KEY (`client_page`) REFERENCES `images` (`id`);

ALTER TABLE `tickets` ADD CONSTRAINT `tickets_id_reward_rewards_id_fk` FOREIGN KEY (`id_reward`) REFERENCES `rewards` (`id`);

ALTER TABLE `tickets` ADD CONSTRAINT `tickets_id_casino_casinos_id_fk` FOREIGN KEY (`id_casino`) REFERENCES `casinos` (`id`);

ALTER TABLE `kiosks` ADD CONSTRAINT `kiosks_id_parameters_parameters_id_fk` FOREIGN KEY (`id_parameter`) REFERENCES `parameters` (`id`);

ALTER TABLE `specific_rules` ADD CONSTRAINT `specific_rules_id_casino_casinos_id_fk` FOREIGN KEY (`id_casino`) REFERENCES `casinos` (`id`);

ALTER TABLE `rewards` ADD CONSTRAINT `rewards_id_parameter_parameters_id_fk` FOREIGN KEY (`id_parameter`) REFERENCES `parameters` (`id`);

ALTER TABLE `rewards` ADD CONSTRAINT `rewards_id_casino_kiosks_id_fk` FOREIGN KEY (`id_casino`) REFERENCES `casinos` (`id`);
