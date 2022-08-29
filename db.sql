DROP TABLE IF EXISTS `card`;
CREATE TABLE IF NOT EXISTS `card` (
  `CardId` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `Front` varchar(100) NOT NULL,
  `Back` varchar(100) NOT NULL,
  `Category` tinyint(4) NOT NULL,
  PRIMARY KEY (`CardId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
