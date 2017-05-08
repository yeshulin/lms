# Host: localhost  (Version: 5.5.53)
# Date: 2017-05-08 10:12:48
# Generator: MySQL-Front 5.3  (Build 4.234)

/*!40101 SET NAMES utf8 */;

#
# Structure for table "members"
#

DROP TABLE IF EXISTS `members`;
CREATE TABLE `members` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `realname` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `addtime` int(11) DEFAULT NULL,
  `updatetime` int(11) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=MyISAM AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;

#
# Data for table "members"
#

/*!40000 ALTER TABLE `members` DISABLE KEYS */;
INSERT INTO `members` VALUES (20,'yeshulin','RE7+OOlq','叶树林','1914875404@qq.com','13512341234',1493737346,1493737346),(21,'yeshulin1','RE7+OOlq','测试1','329801746@qq.com','13543214321',1493906345,1493906345);
/*!40000 ALTER TABLE `members` ENABLE KEYS */;
