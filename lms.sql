/*
Navicat MySQL Data Transfer

Source Server         : 172.29.4.87
Source Server Version : 50629
Source Host           : 172.29.4.87:3306
Source Database       : lms

Target Server Type    : MYSQL
Target Server Version : 50629
File Encoding         : 65001

Date: 2017-05-19 19:21:09
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `members`
-- ----------------------------
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
) ENGINE=MyISAM AUTO_INCREMENT=59 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of members
-- ----------------------------
INSERT INTO `members` VALUES ('38', 'student3', 'RE7+OOlq', '学生', '123456@qq.com', '13512341341', '1494329831', '1494329831');
INSERT INTO `members` VALUES ('21', 'yeshulin1', 'RE7+OOlq', '测试1', '329801746@qq.com', '13543214321', '1493906345', '1493906345');
INSERT INTO `members` VALUES ('27', 'student', 'RE7+OOlq', '学生', '1914875404@qq.com', '13512341234', '1494248324', '1494248324');
INSERT INTO `members` VALUES ('30', 'student2', 'RE7+OOlq', '测试', '191@qq.com', '13512341234', '1494248539', '1494248539');
INSERT INTO `members` VALUES ('29', 'student1', 'RE7+OOlq', '姓名', '1914875404@qq.com', '13512341234', '1494248404', '1494248404');
INSERT INTO `members` VALUES ('46', 'student1', 'RE7+OOlq', 'student1', '194878787@qq.com', '13512341234', '1494662534', '1494662534');
INSERT INTO `members` VALUES ('40', 'yeshulin1', 'RE7+OOlq', '叶树林', '1914875404@qq.com', '13512341234', '1494651921', '1494651921');
INSERT INTO `members` VALUES ('41', 'yeshulin2', 'RE7+OOlq', '测试', '1914875404@qq.com', '13512341234', '1494652079', '1494652079');
INSERT INTO `members` VALUES ('42', 'sobey', 'RE7+OOlq', 'sobey', '194@qq.com', '13512341234', '1494661127', '1494661127');
INSERT INTO `members` VALUES ('47', 'student2', 'RE7+OOlq', '叶树林', '1914875404@qq.com', '13512341234', '1494662853', '1494662853');
INSERT INTO `members` VALUES ('48', 'student3', 'RE7+OOlq', '叶树林', '1914875404@qq.com', '13512341234', '1494663703', '1494663703');
INSERT INTO `members` VALUES ('49', 'student4', 'RE7+OOlq', 'student4', '1914875404@qq.com', '13512341234', '1494664152', '1494664152');
INSERT INTO `members` VALUES ('51', 'student5', 'RE7+OOlq', '12456', '1914875404@qq.com', '13512341234', '1494664321', '1494664321');
INSERT INTO `members` VALUES ('52', 'yeshulin4', 'RE7+OOlq', '叶树林', '1914875404@qq.com', '13512341234', '1494847284', '1494847284');
INSERT INTO `members` VALUES ('53', 'yeshulin23', 'RE7+OOlq', '张三323', '1914875404@qq.com', '13558872536', '1494848244', '1494848244');
INSERT INTO `members` VALUES ('54', '张三', 'RE7+OOlq', '叶树林', '154878@qq.com', '13512341234', '1494848308', '1494848308');
INSERT INTO `members` VALUES ('55', '李四', 'RE7+OOlq', '李四', '1914875404@qq.com', '18012341234', '1494848499', '1494848499');
INSERT INTO `members` VALUES ('56', 'sobey', 'RE7+OOlq', 'sobey123', '144878@qq.com', '18048554844', '1494848543', '1494848543');
INSERT INTO `members` VALUES ('57', 'sobey123', 'RE7+OOlq', 'jkiuiur', '123@qq.com', '13512341234', '1494848616', '1494848616');
INSERT INTO `members` VALUES ('58', 'sobey1234', 'RE7+OOlq', 'SOBEY', '1914875404@qq.com', '13512341234', '1494849266', '1494849266');

-- ----------------------------
-- Table structure for `news`
-- ----------------------------
DROP TABLE IF EXISTS `news`;
CREATE TABLE `news` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `catid` smallint(5) unsigned DEFAULT '0' COMMENT '分类ID',
  `title` varchar(80) NOT NULL DEFAULT '' COMMENT '标题',
  `thumb` varchar(100) NOT NULL DEFAULT '' COMMENT '缩略图',
  `keywords` char(40) NOT NULL DEFAULT '' COMMENT '关键词',
  `description` mediumtext NOT NULL COMMENT '描述',
  `content` text COMMENT '内容',
  `sort` tinyint(3) unsigned DEFAULT '0' COMMENT '排序',
  `status` tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '状态',
  `username` char(20) NOT NULL COMMENT '添加用户',
  `addtime` int(11) unsigned NOT NULL DEFAULT '0',
  `updatetime` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `status` (`status`,`sort`,`id`),
  KEY `listorder` (`catid`,`status`,`sort`,`id`),
  KEY `catid` (`catid`,`status`,`id`)
) ENGINE=MyISAM AUTO_INCREMENT=1050 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of news
-- ----------------------------
INSERT INTO `news` VALUES ('1047', '1', '测试内容1', '', '测试内容2', '测试内容3343443', '<p>测试内容4</p>', '1', '1', '', '1495098531', '1495103815');
INSERT INTO `news` VALUES ('1049', '1', '转转为什么没有变得更好？', '23.jpg', '交网络', '的社交网络也确实有利于转转用户之间进行交易，但这种交易更多是基于熟人社交关系之间，但二手交易最根本核心在于撬动闲置物品需求，而这绝不是一种熟人关系的交易，就连姚劲波自己也承认了这个观点。\n\n姚劲波曾撰文谈为何', '<p class=\"text\">4月18日，58同城宣布与腾讯达成协议，后者将向58集团旗下的二手交易平台「转转」投资2亿美元。腾讯的入局加剧了二手交易市场的竞争局势，在一个月后的今天，或许我们是时候来复盘了。</p><p class=\"text\"><span>营收停止增长后，转转成为一根救命稻草</span></p><p class=\"text\"><br></p>', '2', '1', '', '1495104001', '1495105292');

-- ----------------------------
-- Table structure for `newstype`
-- ----------------------------
DROP TABLE IF EXISTS `newstype`;
CREATE TABLE `newstype` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `catname` varchar(50) DEFAULT NULL COMMENT '分类名称',
  `pid` int(11) DEFAULT '0' COMMENT '父ID',
  `sort` tinyint(4) DEFAULT NULL,
  `addtime` int(11) DEFAULT NULL,
  `updatetime` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of newstype
-- ----------------------------
INSERT INTO `newstype` VALUES ('1', '新闻动态', '0', '1', '1495028014', '1495028014');

-- ----------------------------
-- Table structure for `node`
-- ----------------------------
DROP TABLE IF EXISTS `node`;
CREATE TABLE `node` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(20) DEFAULT NULL COMMENT '节点名称',
  `title` varchar(20) DEFAULT NULL COMMENT '节点标题',
  `status` tinyint(4) DEFAULT NULL COMMENT '节点状态',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `sort` tinyint(4) DEFAULT NULL COMMENT '排序',
  `pid` tinyint(4) DEFAULT '0' COMMENT '父ID',
  `level` tinyint(4) DEFAULT NULL COMMENT '级别',
  `type` varchar(20) DEFAULT NULL COMMENT '类型',
  `group_id` tinyint(4) DEFAULT NULL COMMENT '分组',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of node
-- ----------------------------
INSERT INTO `node` VALUES ('1', 'user', '用户管理', '1', '用户管理', '1', '0', '0', 'model', '0');
INSERT INTO `node` VALUES ('2', 'role', '角色管理', '1', '角色管理', '2', '0', '0', 'model', '0');

-- ----------------------------
-- Table structure for `node_access`
-- ----------------------------
DROP TABLE IF EXISTS `node_access`;
CREATE TABLE `node_access` (
  `role_id` int(11) DEFAULT NULL,
  `node_id` int(11) DEFAULT NULL,
  `level` tinyint(4) DEFAULT NULL,
  `pid` tinyint(4) DEFAULT NULL,
  `module` varchar(10) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of node_access
-- ----------------------------

-- ----------------------------
-- Table structure for `role`
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) DEFAULT NULL COMMENT '角色名称',
  `status` tinyint(4) DEFAULT NULL COMMENT '状态',
  `remark` varchar(100) DEFAULT NULL COMMENT 'P备注',
  `addtime` int(11) DEFAULT NULL,
  `updatetime` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES ('6', '管理员', '1', '管理员角色1232', '1494643853', '1494643853');
INSERT INTO `role` VALUES ('18', '内容管理员', '1', '内容管理员', '1494650587', '1494650587');

-- ----------------------------
-- Table structure for `role_member`
-- ----------------------------
DROP TABLE IF EXISTS `role_member`;
CREATE TABLE `role_member` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of role_member
-- ----------------------------
INSERT INTO `role_member` VALUES ('1', '6', '42');
INSERT INTO `role_member` VALUES ('2', '6', '43');
INSERT INTO `role_member` VALUES ('4', '6', '46');
INSERT INTO `role_member` VALUES ('5', '6', '47');
INSERT INTO `role_member` VALUES ('6', '18', '48');
INSERT INTO `role_member` VALUES ('8', '6', '51');
INSERT INTO `role_member` VALUES ('9', '6', '52');
INSERT INTO `role_member` VALUES ('10', '6', '53');
INSERT INTO `role_member` VALUES ('11', '6', '58');
INSERT INTO `role_member` VALUES ('12', '18', '56');
INSERT INTO `role_member` VALUES ('13', '18', '54');
