-- MySQL dump 10.13  Distrib 8.0.17, for macos10.14 (x86_64)
--
-- Host: 127.0.0.1    Database: bitcoin_db
-- ------------------------------------------------------
-- Server version	8.0.22

--
-- Table structure for table `bitcoins`
--
USE `bitcoin_db`;
DROP TABLE IF EXISTS `bitcoins`;
CREATE TABLE `bitcoins` (
  `id` int NOT NULL AUTO_INCREMENT,
  `base` varchar(45) NOT NULL,
  `currency` varchar(45) DEFAULT NULL,
  `amount` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `bitcoins`
--

LOCK TABLES `bitcoins` WRITE;
INSERT INTO `bitcoins` VALUES (1,'BTC','MXN','368100.0982136'),(2,'BTC','MXN','387100.0982136'),(3,'BTC','MXN','1234.56'),(4,'BTC','MXN','109774'),(5,'BTC','MXN','109974'),(6,'BTC','MXN','109974'),(7,'BTC','MXN','109974'),(8,'BTC','MXN','109974'),(9,'BTC','MXN','360416.509189');
UNLOCK TABLES;

