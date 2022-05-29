-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               8.0.28 - MySQL Community Server - GPL
-- Server OS:                    Linux
-- HeidiSQL Version:             11.3.0.6295
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for todo
DROP DATABASE IF EXISTS `todo`;
CREATE DATABASE IF NOT EXISTS `todo` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `todo`;

-- Dumping structure for procedure todo.SP_Task_Create
DROP PROCEDURE IF EXISTS `SP_Task_Create`;
DELIMITER //
CREATE PROCEDURE `SP_Task_Create`(
    IN `_title` TEXT,
    IN `_description` TEXT,
    IN `_create_by` INT
)
BEGIN
    DECLARE _row_count INT DEFAULT 0;
    DECLARE _error_code INT DEFAULT 0;
    DECLARE exit handler for sqlexception
        BEGIN
            ROLLBACK;
            SELECT -99 AS error_code, 'exception' AS msg;
        END;


    START TRANSACTION;
    IF NOT EXISTS (SELECT * FROM tbl_user WHERE tbl_user.id = _create_by) THEN
        SELECT 2 AS error_code, 'khong tim thay user tao task' AS msg;
    END IF;
    INSERT INTO todo.tbl_task(title, description, created_by) VALUES(_title, _description, _create_by);
    SELECT ROW_COUNT() INTO _row_count;
    IF _row_count <= 0 THEN
        ROLLBACK;
        SET _error_code = 1;
        SELECT _error_code AS error_code, 'khong the tao task' AS msg;
    END IF;
    SET _error_code = 0;
    COMMIT;
    SELECT _error_code AS error_code, 'ok' AS msg;
END//
DELIMITER ;

-- Dumping structure for table todo.tbl_task
DROP TABLE IF EXISTS `tbl_task`;
CREATE TABLE IF NOT EXISTS `tbl_task` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` text NOT NULL,
  `description` text,
  `created_by` int DEFAULT NULL,
  `date` date NOT NULL DEFAULT (curdate()),
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `created_by` (`created_by`),
  CONSTRAINT `tbl_task_ibfk_1` FOREIGN KEY (`created_by`) REFERENCES `tbl_user` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table todo.tbl_user
DROP TABLE IF EXISTS `tbl_user`;
CREATE TABLE IF NOT EXISTS `tbl_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `limit_per_day` int NOT NULL,
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username_ind` (`username`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
