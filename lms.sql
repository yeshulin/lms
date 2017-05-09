/*
Navicat MySQL Data Transfer

Source Server         : 172.29.4.87
Source Server Version : 50629
Source Host           : 172.29.4.87:3306
Source Database       : lms

Target Server Type    : MYSQL
Target Server Version : 50629
File Encoding         : 65001

Date: 2017-05-09 21:24:39
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
) ENGINE=MyISAM AUTO_INCREMENT=40 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of members
-- ----------------------------
INSERT INTO `members` VALUES ('38', 'student3', 'RE7+OOlq', '学生', '123456@qq.com', '13512341341', '1494329831', '1494329831');
INSERT INTO `members` VALUES ('21', 'yeshulin1', 'RE7+OOlq', '测试1', '329801746@qq.com', '13543214321', '1493906345', '1493906345');
INSERT INTO `members` VALUES ('27', 'student', 'RE7+OOlq', '学生', '1914875404@qq.com', '13512341234', '1494248324', '1494248324');
INSERT INTO `members` VALUES ('30', 'student2', 'RE7+OOlq', '测试', '191@qq.com', '13512341234', '1494248539', '1494248539');
INSERT INTO `members` VALUES ('29', 'student1', 'RE7+OOlq', '姓名', '1914875404@qq.com', '13512341234', '1494248404', '1494248404');
INSERT INTO `members` VALUES ('39', '34234', 'R0/5P+hv', '234432', '342324@qq.com', '13812341234', '1494335376', '1494335376');

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
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES ('6', '管理员', '1', '管理员角色', '1494336042', '1494336042');
