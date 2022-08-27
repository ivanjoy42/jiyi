DROP TABLE IF EXISTS `card`;
CREATE TABLE IF NOT EXISTS `card` (
  `cardId` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `front` varchar(100) NOT NULL,
  `back` varchar(100) NOT NULL,
  `type` tinyint(4) NOT NULL,
  PRIMARY KEY (`cardId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
