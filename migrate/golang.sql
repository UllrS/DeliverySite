
CREATE TABLE `admins` (
  `id` int(5) NOT NULL,
  `login` varchar(64) NOT NULL,
  `password` varchar(64) NOT NULL,
  `token` varchar(64) DEFAULT ''
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

CREATE TABLE `merchants` (
  `id` int(12) UNSIGNED NOT NULL,
  `name` varchar(100) NOT NULL,
  `type` varchar(100) NOT NULL,
  `addr` varchar(100) NOT NULL DEFAULT '',
  `anons` varchar(1000) NOT NULL DEFAULT '',
  `img` varchar(100) DEFAULT '',
  `date` date NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

CREATE TABLE `orders` (
  `id` int(6) NOT NULL,
  `token` varchar(64) NOT NULL DEFAULT '',
  `user` int(6) DEFAULT '-1',
  `shipping` varchar(100) NOT NULL DEFAULT '',
  `basket` varchar(3000) NOT NULL,
  `tel` varchar(15) NOT NULL DEFAULT '',
  `status` int(12) NOT NULL DEFAULT '0',
  `date` date NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `price` float DEFAULT NULL,
  `category` varchar(100) NOT NULL DEFAULT '',
  `anons` varchar(500) DEFAULT '',
  `merch` int(12) NOT NULL,
  `portion` float NOT NULL,
  `unit` varchar(10) NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

CREATE TABLE `users` (
  `id` int(12) UNSIGNED NOT NULL,
  `name` varchar(64) DEFAULT NULL,
  `password` varchar(64) DEFAULT NULL,
  `token` varchar(64) DEFAULT NULL,
  `basket` json DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

INSERT INTO `admins`(`id`,`login`, `password`) VALUES ('0','admin','admin')