CREATE DATABASE IF NOT EXISTS appdb;
USE appdb;
--
-- Table structure for table `customers`
--

DROP TABLE IF EXISTS customers;
CREATE TABLE customers (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(255) DEFAULT NULL,
  address varchar(255) DEFAULT NULL,
  PRIMARY KEY (id)
); 



DROP TABLE IF EXISTS `Firmen`;
CREATE TABLE `Firmen` (  
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,  
  `enabled` int(1) NOT NULL,
  PRIMARY KEY (`id`)
);


DROP TABLE IF EXISTS `Personal`;
CREATE TABLE `Personal` (
`id`	int(11) NOT NULL AUTO_INCREMENT,
`firmaId` int(11) NOT NULL,
`persnr` int(11) NOT NULL,
`nachname`	varchar(255) NOT NULL,
`vorname`	varchar(255) NOT NULL,
`titel`	varchar(255) DEFAULT NULL ,
`geschlecht`	varchar(255) NOT NULL,
`eintritt`	varchar(255) NOT NULL,
`austritt`	varchar(255) ,
`abteilungId`	int(11) NOT NULL,
`gebdat`	varchar(255),
`dienstart`	varchar(255) NOT NULL,
`berufsbezeichnung` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `Abteilungen`;
CREATE TABLE `Abteilungen` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `firmaId` int(11) NOT NULL,
  `notifyMail` varchar(255) NOT NULL,  
  `enabled` int(1) NOT NULL,  
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `Items`;
CREATE TABLE `Items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `processorId` int(1) NOT NULL,
  `text` varchar(255) DEFAULT NULL,  
  `type` varchar(255) DEFAULT NULL,  
  `enabled` int(1) NOT NULL,  
  `depId` int(11) NOT NULL DEFAULT '0',  
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `Processors`;
CREATE TABLE `Processors` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name`  varchar(255) NOT NULL,
  `notifyMail` varchar(255) DEFAULT NULL, 
`enabled` int(1) NOT NULL,  
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `Person_Processor_Status`;
CREATE TABLE `Person_Processor_Status` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `personId` int(11) NOT NULL ,
  `processorId`  int(11) NOT NULL,
  `status`  int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `Person_Abteilung_Status`;
CREATE TABLE `Person_Abteilung_Status` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `personId` int(11) NOT NULL,
  `abteilungId`  int(11) NOT NULL,
  `status`  int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `Person_Item`;
CREATE TABLE `Person_Item` (
  `personId` int(11) NOT NULL,
  `itemId`  int(11) NOT NULL,  
  `value` varchar(255) DEFAULT NULL  
);

--
-- Dumping data for table `customers`
--

LOCK TABLES `Firmen` WRITE;
/*!40000 ALTER TABLE `Firmen` DISABLE KEYS */;
INSERT INTO `Firmen` (`id`, `name`, `enabled`) VALUES ('0', 'ALL', '1');
INSERT INTO `Firmen` (`id`, `name`, `enabled`) VALUES ('1', 'MCB', '1');
INSERT INTO `Firmen` (`name`, `enabled`) VALUES ('Klinikum FN','1');
INSERT INTO `Firmen` (`name`, `enabled`) VALUES ('Klinik TET','1');
/*!40000 ALTER TABLE `Firmen` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `Abteilungen` WRITE;
/*!40000 ALTER TABLE `Abteilungen` DISABLE KEYS */;
INSERT INTO `Abteilungen` (`id`,`name`,`FirmaId`,`notifyMail`,`enabled`) VALUES ('0','ALL','0','jojo@ulewu.de','1');
INSERT INTO `Abteilungen` (`name`,`FirmaId`,`notifyMail`,`enabled`) VALUES ('Station 4','1','st4@ulewu.de','1');
INSERT INTO `Abteilungen` (`name`,`FirmaId`,`notifyMail`,`enabled`) VALUES ('MVZ I','1','mvz@ulewu.de','1');
/*!40000 ALTER TABLE `Abteilungen` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `Items` WRITE;
/*!40000 ALTER TABLE `Items` DISABLE KEYS */;
INSERT INTO `Items` (`processorId`,`text`,`type`,`enabled`) VALUES ('1','Email Adresse anlegen','bool','1');
/*!40000 ALTER TABLE `Items` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `Processors` WRITE;
/*!40000 ALTER TABLE `Processors` DISABLE KEYS */;
INSERT INTO `Processors` (`name`,`notifyMail`,`enabled`) VALUES ('EDV-Sachbearbeitung','ulewu@ulewu.de','1');
/*!40000 ALTER TABLE `Processors` ENABLE KEYS */;
UNLOCK TABLES;


LOCK TABLES `Personal` WRITE;
/*!40000 ALTER TABLE `Personal` DISABLE KEYS */;
INSERT INTO `Personal` 
(`firmaId`,`persnr`,	`nachname`,`vorname`,`titel`,`geschlecht`,`eintritt`,`austritt`,`abteilungId`,`gebdat`,`dienstart`,`berufsbezeichnung`)
VALUES
(1,003673,'Kleinmann','Brigitte', 'Dr.','weiblich','01.09.2018','31.08.2019','1','20.06.2001','Verwaltungsdienst','Kaufm. Mitarbeiter/in');

INSERT INTO `Personal` 
(`firmaId`,`persnr`,	`nachname`,`vorname`,`titel`,`geschlecht`,`eintritt`,`austritt`,`abteilungId`,`gebdat`,`dienstart`,`berufsbezeichnung`)
VALUES
(1,003578,'Kreutzer','Helmut','','männlich','01.09.2018','31.08.2021','2','04.07.1986','Verwaltungsdienst','Auszubildende/r');
/*!40000 ALTER TABLE `Personal` ENABLE KEYS */;
UNLOCK TABLES;

-- LOCK TABLES `Person_Processor_Status` WRITE;
-- /*!40000 ALTER TABLE `Person_Processor_Status` DISABLE KEYS */;
-- INSERT INTO `Person_Processor_Status` (personId,processorId,status) VALUES (`1`,`1`,`0`);
-- /*!40000 ALTER TABLE `Person_Processor_Status` ENABLE KEYS */;
-- UNLOCK TABLES;

-- LOCK TABLES `Person_Abteilung_Status` WRITE;
-- /*!40000 ALTER TABLE `Person_Abteilung_Status` DISABLE KEYS */;
-- INSERT INTO `Person_Abteilung_Status` (personId,abteilungId,status) VALUES (`1`,`1`,`0`);
-- /*!40000 ALTER TABLE `Person_Abteilung_Status` ENABLE KEYS */;
-- UNLOCK TABLES;

--- Person-Item -> Soll beim ausfüllen des Formulars gespeichert werden.

--- Person -> Soll beim Anlegen einer Person gespeichert werden.




