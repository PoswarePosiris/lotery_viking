-- MySQL dump 10.13  Distrib 9.0.1, for macos14.4 (arm64)
--
-- Host: localhost    Database: lotery_viking
-- ------------------------------------------------------
-- Server version	9.0.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


--
-- Table structure for table `images`
--

DROP TABLE IF EXISTS `images`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `images` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `format` varchar(100) DEFAULT NULL,
  `url` varchar(256) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT (now()),
  `updated_at` timestamp NULL DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;


--
-- Table structure for table `parameters`
--

DROP TABLE IF EXISTS `parameters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `parameters` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name_lotery` varchar(100) DEFAULT NULL,
  `name_casino` varchar(100) DEFAULT NULL,
  `date_start` varchar(100) DEFAULT NULL,
  `date_end` varchar(100) DEFAULT NULL,
  `status` enum('scan','draw') DEFAULT NULL,
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
  KEY `parameters_result_page_images_id_fk` (`result_page`),
  CONSTRAINT `parameters_home_page_images_id_fk` FOREIGN KEY (`home_page`) REFERENCES `images` (`id`),
  CONSTRAINT `parameters_result_page_images_id_fk` FOREIGN KEY (`result_page`) REFERENCES `images` (`id`),
  CONSTRAINT `parameters_scan_page_images_id_fk` FOREIGN KEY (`scan_page`) REFERENCES `images` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;


--
-- Table structure for table `kiosks`
--

DROP TABLE IF EXISTS `kiosks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
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
  KEY `kiosks_id_parameters_parameters_id_fk` (`id_parameters`),
  CONSTRAINT `kiosks_id_parameters_parameters_id_fk` FOREIGN KEY (`id_parameters`) REFERENCES `parameters` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `publicity_images`
--

DROP TABLE IF EXISTS `publicity_images`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `publicity_images` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `parameter_id` bigint unsigned DEFAULT NULL,
  `image_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `publicity_images_parameter_id_parameters_id_fk` (`parameter_id`),
  KEY `publicity_images_image_id_images_id_fk` (`image_id`),
  CONSTRAINT `publicity_images_image_id_images_id_fk` FOREIGN KEY (`image_id`) REFERENCES `images` (`id`),
  CONSTRAINT `publicity_images_parameter_id_parameters_id_fk` FOREIGN KEY (`parameter_id`) REFERENCES `parameters` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `rewards`
--

DROP TABLE IF EXISTS `rewards`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rewards` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `big_win` tinyint(1) DEFAULT '0',
  `amount` int DEFAULT NULL,
  `id_images` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `rewards_id_images_images_id_fk` (`id_images`),
  CONSTRAINT `rewards_id_images_images_id_fk` FOREIGN KEY (`id_images`) REFERENCES `images` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `tickets`
--

DROP TABLE IF EXISTS `tickets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
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
  KEY `tickets_id_reward_rewards_id_fk` (`id_reward`),
  CONSTRAINT `tickets_id_reward_rewards_id_fk` FOREIGN KEY (`id_reward`) REFERENCES `rewards` (`id`),
  CONSTRAINT `tickets_kiosk_id_kiosks_id_fk` FOREIGN KEY (`kiosk_id`) REFERENCES `kiosks` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
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
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

DROP VIEW IF EXISTS `kiosk_view`;

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
    -- Concatenate all publicity images related to the parameters
    GROUP_CONCAT(DISTINCT `publicity_imgs`.`url` SEPARATOR ', ') AS `publicity`,
    
    -- Use CASE statements to selectively aggregate the correct image URLs
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
    LEFT JOIN `images` AS `publicity_imgs` ON `publicity_imgs`.`id` = `publicity_images`.`image_id`;


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-08-27 10:37:00
