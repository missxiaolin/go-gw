# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.24)
# Database: bs_gw
# Generation Time: 2019-07-30 10:17:22 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table article
# ------------------------------------------------------------

DROP TABLE IF EXISTS `article`;

CREATE TABLE `article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '文章表主键',
  `cid` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '分类id',
  `title` char(100) NOT NULL DEFAULT '' COMMENT '标题',
  `author` varchar(15) NOT NULL DEFAULT '' COMMENT '作者',
  `content` mediumtext NOT NULL COMMENT '文章内容',
  `keywords` varchar(255) NOT NULL DEFAULT '' COMMENT '关键字',
  `description` char(255) NOT NULL DEFAULT '' COMMENT '描述',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '文章是否显示 1是 0否',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table category
# ------------------------------------------------------------

DROP TABLE IF EXISTS `category`;

CREATE TABLE `category` (
  `id` tinyint(2) unsigned NOT NULL AUTO_INCREMENT COMMENT '分类主键id',
  `name` varchar(15) NOT NULL DEFAULT '' COMMENT '分类名称',
  `description` varchar(255) DEFAULT '' COMMENT '描述',
  `pid` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '父级栏目id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;

INSERT INTO `category` (`id`, `name`, `description`, `pid`, `created_at`, `updated_at`)
VALUES
	(1,'分类1','分类1',0,'2019-07-30 17:29:30','2019-07-30 17:29:24'),
	(2,'分类2','分类2',0,'2019-07-30 17:29:39','2019-07-30 17:29:37');

/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table config
# ------------------------------------------------------------

DROP TABLE IF EXISTS `config`;

CREATE TABLE `config` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '配置项键名',
  `value` text COMMENT '配置项键值 1表示开启 0 关闭',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table link
# ------------------------------------------------------------

DROP TABLE IF EXISTS `link`;

CREATE TABLE `link` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '链接名',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '链接地址',
  `sort` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '排序',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(60) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(64) NOT NULL DEFAULT '' COMMENT '登录密码；mb_password加密',
  `status` tinyint(1) NOT NULL DEFAULT '2' COMMENT '用户状态 0：禁用； 1：正常 ；2：未验证',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_login_key` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
