-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jul 09, 2024 at 01:50 PM
-- Server version: 8.0.30
-- PHP Version: 7.4.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `map-task`
--
CREATE DATABASE IF NOT EXISTS `map-task` DEFAULT CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci;
USE `map-task`;

DELIMITER $$
--
-- Procedures
--
DROP PROCEDURE IF EXISTS `GetNearestLocationsWithPagination`$$
CREATE DEFINER=`root`@`localhost` PROCEDURE `GetNearestLocationsWithPagination` (IN `limit_count` INT, IN `offset_count` INT, IN `input_lat` FLOAT, IN `input_long` FLOAT)   BEGIN    
    SELECT id, name, latitude, longitude, marker,
           (6371 * acos(cos(radians(input_lat)) * cos(radians(latitude)) * cos(radians(longitude) - radians(input_long)) + sin(radians(input_lat)) * sin(radians(latitude)))) AS distance
    FROM locations
    ORDER BY distance
    LIMIT limit_count OFFSET offset_count;
END$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `locations`
--

DROP TABLE IF EXISTS `locations`;
CREATE TABLE IF NOT EXISTS `locations` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `latitude` double NOT NULL,
  `longitude` double NOT NULL,
  `marker` varchar(45) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_locations_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb3;

--
-- Truncate table before insert `locations`
--

TRUNCATE TABLE `locations`;
--
-- Dumping data for table `locations`
--

INSERT DELAYED IGNORE INTO `locations` (`id`, `name`, `latitude`, `longitude`, `marker`, `created_at`, `updated_at`, `deleted_at`) VALUES
(3, 'Anitakbir', 39.9304, 32.8541, 'white', '2024-07-09 15:34:39.484', '2024-07-09 15:34:39.484', NULL),
(4, 'itu', 41.1053, 29.024, 'white', '2024-07-09 15:35:48.999', '2024-07-09 15:35:48.999', NULL),
(5, 'ytu', 41.0432, 29.0096, 'white', '2024-07-09 15:36:33.263', '2024-07-09 15:36:33.263', NULL),
(6, 'odtu', 39.8999, 32.7806, 'white', '2024-07-09 15:36:55.925', '2024-07-09 15:36:55.925', NULL);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
