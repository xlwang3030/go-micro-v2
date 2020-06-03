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

 Date: 16/05/2020 16:08:48
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for course_topic
-- ----------------------------
DROP TABLE IF EXISTS `course_topic`;
CREATE TABLE `course_topic`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `course_id` int(11) NOT NULL DEFAULT 0 COMMENT '课程ID',
  `course_did` int(11) NOT NULL DEFAULT 0 COMMENT '课时ID',
  `likes` int(11) NULL DEFAULT NULL COMMENT '点赞数',
  `unlikes` int(11) NULL DEFAULT NULL COMMENT '不认同数',
  `title` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '标题',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL COMMENT '评论内容',
  `user_id` int(11) NULL DEFAULT NULL COMMENT '发表者ID',
  `addtime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '发布时间',
  `updatetime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '最后修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

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
