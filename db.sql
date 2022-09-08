DROP TABLE IF EXISTS `card`;
CREATE TABLE IF NOT EXISTS `card` (
  `card_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `kind` tinyint(4) NOT NULL,
  `front` varchar(100) NOT NULL,
  `back` varchar(100) NOT NULL,
  PRIMARY KEY (`card_id`) USING BTREE,
  UNIQUE KEY `kind_front` (`kind`,`front`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `card_deck`;
CREATE TABLE IF NOT EXISTS `card_deck` (
  `card_deck_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `card_id` int(10) unsigned NOT NULL,
  `deck_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`card_deck_id`) USING BTREE,
  UNIQUE KEY `card_id_deck_id` (`card_id`,`deck_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `deck`;
CREATE TABLE IF NOT EXISTS `deck` (
  `deck_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `kind` tinyint(4) NOT NULL,
  `deck_name` varchar(100) NOT NULL,
  PRIMARY KEY (`deck_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
