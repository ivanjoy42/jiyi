DROP TABLE IF EXISTS `card`;
CREATE TABLE IF NOT EXISTS `card` (
  `cardId` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `front` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `back` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `type` tinyint(4) NOT NULL,
  PRIMARY KEY (`cardId`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
