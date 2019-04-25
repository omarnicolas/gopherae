USE `gopherae`;

CREATE TABLE `gopher` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `short_description` longtext COLLATE utf8_unicode_ci NOT NULL,
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `review` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gopher_id`int(11) NOT NULL,
  `first_name` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `last_name` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `text` longtext COLLATE utf8_unicode_ci NOT NULL,
  `score`int(11) NOT NULL,
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`gopher_id`) REFERENCES gopher(`id`)
);