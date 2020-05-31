/*
 Navicat Premium Data Transfer

 Source Server         : mysql57
 Source Server Type    : MySQL
 Source Server Version : 50721
 Source Host           : localhost:3307
 Source Schema         : jt

 Target Server Type    : MySQL
 Target Server Version : 50721
 File Encoding         : 65001

 Date: 16/05/2020 16:08:54
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for course_topic_reply
-- ----------------------------
DROP TABLE IF EXISTS `course_topic_reply`;
CREATE TABLE `course_topic_reply`  (
  `id` int(11) NOT NULL,
  `topic_id` int(11) NOT NULL COMMENT '短评ID',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL COMMENT '回复内容',
  `user_id` int(11) NULL DEFAULT NULL COMMENT '回复用户ID',
  `likes` int(11) NULL DEFAULT NULL COMMENT '认同',
  `unlikes` int(11) NULL DEFAULT NULL COMMENT '不认同',
  `addtime` timestamp(0) NULL DEFAULT NULL COMMENT '入库时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
