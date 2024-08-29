-- Create tables
CREATE TABLE `images` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(100) DEFAULT NULL,
  `format` varchar(100) DEFAULT NULL,
  `url` varchar(256) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT (now()),
  `updated_at` timestamp NULL DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `parameters` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name_lotery` varchar(100) DEFAULT NULL,
  `name_casino` varchar(100) DEFAULT NULL,
  `date_start` varchar(100) DEFAULT NULL,
  `date_end` varchar(100) DEFAULT NULL,
  `status` enum('scan','draw') DEFAULT NULL,
  `client_data` tinyint(1) DEFAULT '0',
  `home_page` bigint unsigned DEFAULT NULL,
  `scan_page` bigint unsigned DEFAULT NULL,
  `result_page` bigint unsigned DEFAULT NULL,
  `general_rules` text,
  `specific_rules` text,
  `secret` varchar(256) DEFAULT NULL,
  `secret_length` int DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT (now()),
  `updated_at` timestamp NULL DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `parameters_home_page_images_id_fk` (`home_page`),
  KEY `parameters_scan_page_images_id_fk` (`scan_page`),
  KEY `parameters_result_page_images_id_fk` (`result_page`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `kiosks` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `macadress_wifi` varchar(100) DEFAULT NULL,
  `macadress_ethernet` varchar(100) DEFAULT NULL,
  `location` varchar(256) DEFAULT NULL,
  `id_parameters` bigint unsigned DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT (now()),
  `updated_at` timestamp NULL DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `kiosks_id_parameters_parameters_id_fk` (`id_parameters`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `publicity_images` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `parameter_id` bigint unsigned DEFAULT NULL,
  `image_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `publicity_images_parameter_id_parameters_id_fk` (`parameter_id`),
  KEY `publicity_images_image_id_images_id_fk` (`image_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `rewards` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `big_win` tinyint(1) DEFAULT '0',
  `amount` int DEFAULT NULL,
  `id_images` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `rewards_id_images_images_id_fk` (`id_images`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `tickets` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `kiosk_id` bigint unsigned DEFAULT NULL,
  `id_reward` bigint unsigned DEFAULT NULL,
  `ticket_number` varchar(256) NOT NULL,
  `claim` tinyint(1) DEFAULT '0',
  `entry_scan` datetime DEFAULT NULL,
  `exit_scan` datetime DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT (now()),
  `updated_at` timestamp NULL DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  UNIQUE KEY `tickets_ticket_number_unique` (`ticket_number`),
  KEY `ticket_number_idx` (`ticket_number`),
  KEY `tickets_kiosk_id_kiosks_id_fk` (`kiosk_id`),
  KEY `tickets_id_reward_rewards_id_fk` (`id_reward`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `email` varchar(256) NOT NULL,
  `password` varchar(256) NOT NULL,
  `created_at` timestamp NULL DEFAULT (now()),
  `updated_at` timestamp NULL DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  UNIQUE KEY `users_email_unique` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Create view
CREATE VIEW `kiosk_view` AS
SELECT
    `kiosks`.`id`,
    `kiosks`.`name`,
    `kiosks`.`macadress_wifi`,
    `kiosks`.`macadress_ethernet`,
    `kiosks`.`location`,
    `parameters`.`name_lotery`,
    `parameters`.`name_casino`,
    `parameters`.`date_start`,
    `parameters`.`date_end`,
    `parameters`.`status`,
    `parameters`.`client_data`,
    GROUP_CONCAT(DISTINCT `publicity_imgs`.`url` SEPARATOR ', ') AS `publicity`,
    GROUP_CONCAT(DISTINCT CASE WHEN `images`.`id` = `parameters`.`home_page` THEN `images`.`url` ELSE NULL END) AS `home_page`,
    GROUP_CONCAT(DISTINCT CASE WHEN `images`.`id` = `parameters`.`scan_page` THEN `images`.`url` ELSE NULL END) AS `scan_page`,
    GROUP_CONCAT(DISTINCT CASE WHEN `images`.`id` = `parameters`.`result_page` THEN `images`.`url` ELSE NULL END) AS `result_page`,
    `parameters`.`general_rules`,
    `parameters`.`specific_rules`,
    `parameters`.`secret`,
    `parameters`.`secret_length`,
    `kiosks`.`updated_at`,
    `parameters`.`updated_at` AS `updated_at_parameters`
FROM
    `kiosks`
    LEFT JOIN `parameters` ON `parameters`.`id` = `kiosks`.`id_parameters`
    LEFT JOIN `images` ON `images`.`id` IN (`parameters`.`home_page`, `parameters`.`scan_page`, `parameters`.`result_page`)
    LEFT JOIN `publicity_images` ON `publicity_images`.`parameter_id` = `parameters`.`id`
    LEFT JOIN `images` AS `publicity_imgs` ON `publicity_imgs`.`id` = `publicity_images`.`image_id`
    GROUP BY `kiosks`.`id`;

-- Add foreign key constraints
ALTER TABLE `publicity_images`
ADD CONSTRAINT `publicity_images_image_id_images_id_fk`
FOREIGN KEY (`image_id`) REFERENCES `images` (`id`);

ALTER TABLE `publicity_images`
ADD CONSTRAINT `publicity_images_parameter_id_parameters_id_fk`
FOREIGN KEY (`parameter_id`) REFERENCES `parameters` (`id`);

ALTER TABLE `rewards`
ADD CONSTRAINT `rewards_id_images_images_id_fk`
FOREIGN KEY (`id_images`) REFERENCES `images` (`id`);

ALTER TABLE `parameters`
ADD CONSTRAINT `parameters_home_page_images_id_fk`
FOREIGN KEY (`home_page`) REFERENCES `images` (`id`);

ALTER TABLE `parameters`
ADD CONSTRAINT `parameters_result_page_images_id_fk`
FOREIGN KEY (`result_page`) REFERENCES `images` (`id`);

ALTER TABLE `parameters`
ADD CONSTRAINT `parameters_scan_page_images_id_fk`
FOREIGN KEY (`scan_page`) REFERENCES `images` (`id`);

ALTER TABLE `tickets`
ADD CONSTRAINT `tickets_id_reward_rewards_id_fk`
FOREIGN KEY (`id_reward`) REFERENCES `rewards` (`id`);

ALTER TABLE `tickets`
ADD CONSTRAINT `tickets_kiosk_id_kiosks_id_fk`
FOREIGN KEY (`kiosk_id`) REFERENCES `kiosks` (`id`);

ALTER TABLE `kiosks`
ADD CONSTRAINT `kiosks_id_parameters_parameters_id_fk`
FOREIGN KEY (`id_parameters`) REFERENCES `parameters` (`id`);
