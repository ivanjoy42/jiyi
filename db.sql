DROP TABLE IF EXISTS `card`;
CREATE TABLE IF NOT EXISTS `card` (
  `card_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `dir_id` int(10) unsigned NOT NULL,
  `front` varchar(100) NOT NULL,
  `back` varchar(100) NOT NULL,
  `helper` varchar(100) NOT NULL,
  `pinyin` varchar(100) NOT NULL,
  PRIMARY KEY (`card_id`),
  UNIQUE KEY `dir_id_front` (`dir_id`,`front`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `card_deck`;
CREATE TABLE IF NOT EXISTS `card_deck` (
  `card_deck_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `card_id` int(10) unsigned NOT NULL,
  `deck_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`card_deck_id`),
  UNIQUE KEY `card_id_deck_id` (`card_id`,`deck_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `deck`;
CREATE TABLE IF NOT EXISTS `deck` (
  `deck_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `dir_id` int(10) unsigned NOT NULL,
  `deck_name` varchar(100) NOT NULL,
  PRIMARY KEY (`deck_id`),
  UNIQUE KEY `dir_id_deck_name` (`dir_id`,`deck_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `dir`;
CREATE TABLE IF NOT EXISTS `dir` (
  `dir_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `dir_name` varchar(100) NOT NULL,
  `kind_id` tinyint(4) NOT NULL,
  PRIMARY KEY (`dir_id`),
  UNIQUE KEY `dir_name` (`dir_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `learn`;
CREATE TABLE IF NOT EXISTS `learn` (
  `learn_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `mode_id` int(10) unsigned NOT NULL,
  `dir_id` int(10) unsigned NOT NULL,
  `deck_id` int(10) unsigned NOT NULL,
  `learn_name` varchar(100) NOT NULL,
  PRIMARY KEY (`learn_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
