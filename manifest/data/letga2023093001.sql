-- MySQL dump 10.13  Distrib 8.0.32, for Win64 (x86_64)
--
-- Host: localhost    Database: letga
-- ------------------------------------------------------
-- Server version	8.0.32

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '唯一ID',
  `account` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '账号',
  `mobile` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '手机号',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '电子邮箱',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `salt` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码盐',
  `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '昵称',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '头像',
  `signature` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '个性签名',
  `signin_role` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '登录角色',
  `signin_failure` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '失败次数',
  `signin_ip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '登录IP',
  `signin_at` datetime DEFAULT NULL COMMENT '登录日期',
  `status` enum('normal','disabled') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'normal' COMMENT '状态',
  `create_at` datetime DEFAULT NULL COMMENT '创建日期',
  `update_at` datetime DEFAULT NULL COMMENT '更新日期',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除日期',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_un_uuid` (`uuid`),
  UNIQUE KEY `user_un_account` (`account`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'3y4eyz08zg0cv5yjfwldz5s100tawyhg','letga','13800138001','letga@qq.com','edd19d3651d570cd3037b4060e5c9df2','hDCr','Letga','/upload/20230904/cv9y7efhuuh4a4o5z5.png','','root',0,NULL,'2023-09-29 07:27:33','normal','2023-08-30 22:43:19','2023-09-29 07:27:33',NULL),(2,'3y4eyz09ig0cvcbb7ob7o00100e5c1o3','letgaTest1','13800138101','13800138101@qq.com','f4c56e9358a373325fb60e88c9e51a16','qHJV','刘一','','',NULL,0,NULL,NULL,'normal','2023-09-07 09:59:48','2023-09-07 10:37:01',NULL),(3,'3y4eyz09ig0cvcc0q9akhxs200ivqbju','letgaTest2','','','5cb995aefdefdf9c8b68dc1d205cf1f1','BAOg','陈二','','',NULL,0,NULL,NULL,'normal','2023-09-07 10:33:08','2023-09-07 10:37:11',NULL),(4,'3y4eyz09ig0cvcc12dchzj4300ho0csb','letgaTest3','','','6736b4f803c37c88f47dd8563961b138','KcdC','张三','','',NULL,0,NULL,NULL,'normal','2023-09-07 10:33:34','2023-09-07 10:37:33',NULL),(5,'3y4eyz09ig0cvcc1b2v3244400gspt1u','letgaTest4','','','a175697d946e1da10691663f5087f399','EgWb','李四','','',NULL,0,NULL,NULL,'normal','2023-09-07 10:33:53','2023-09-07 10:37:45',NULL),(6,'3y4eyz09ig0cvcc1l2t46zk500p2jgrn','letgaTest5','','','43972f61d0ff035fb83ef75d9c001f85','SeRL','王五','','',NULL,0,NULL,NULL,'normal','2023-09-07 10:34:15','2023-09-07 10:37:59',NULL),(7,'3y4eyz09ig0cvcc1wv50a2w600qmkog6','letgaTest6','','','9931d616ef615736b8bb03e4fb351d1b','dBRi','赵六','','',NULL,0,NULL,NULL,'normal','2023-09-07 10:34:40','2023-09-07 10:38:13',NULL),(8,'3y4eyz09ig0cvcc2ukh6yrc700fk2exj','letgaTest7','','','9a69fbb3e66dc7706f348a61a932880d','yuza','孙七','','',NULL,0,NULL,NULL,'normal','2023-09-07 10:35:54','2023-09-07 10:38:25',NULL),(9,'3y4eyz09ig0cvcc31pfop3w800fs768h','letgaTest8','','','622d184ccc5c99c7844548116fb95251','QALb','周八','','',NULL,0,NULL,NULL,'normal','2023-09-07 10:36:09','2023-09-07 10:38:41',NULL),(10,'3y4eyz09ig0cvcc5l88nhco900d6nmuf','letgaTest9','','','7cc328245b1821021cf68b5ca13010fa','pKEI','吴九','',NULL,NULL,0,NULL,NULL,'normal','2023-09-07 10:39:29','2023-09-07 10:39:29',NULL),(11,'3y4eyz09ig0cvcc60g1u3r0a00j00td8','LetgaTest10','','','5922e95e72bac469eacc5356e7e0ef6f','GyoB','郑十','',NULL,NULL,0,NULL,NULL,'normal','2023-09-07 10:40:02','2023-09-07 10:40:02',NULL),(12,'3y4eyz09ig0cvcc6m7hhskgb00rm36up','Menterma','','','2be67fe4fe05c2c334289a7f871a66bc','ZObR','Menterma','',NULL,NULL,0,NULL,NULL,'normal','2023-09-07 10:40:49','2023-09-07 10:40:49',NULL),(13,'3y4eyz09ig0cvcc6t0d5xlgc000js50e','Mibargu','','','51fadd6d6b135bd7dfe8d3155914bff0','Gcik','Mibargu','',NULL,NULL,0,NULL,NULL,'normal','2023-09-07 10:41:04','2023-09-07 10:41:04',NULL),(14,'3y4eyz09ig0cvcc6yjmwz00d000t2dp6','Comfyre','','','f4f1cd79db717d90650ff74fc217d61f','wzCR','Comfyre','',NULL,NULL,0,NULL,NULL,'normal','2023-09-07 10:41:16','2023-09-07 10:41:16',NULL),(15,'3y4eyz09ig0cvcc7eu6hxwoe00zibj70','Sityqsou','','','5c8e36e777f66c01521437083e67a9a5','wuoy','Sityqsou','',NULL,NULL,0,NULL,NULL,'normal','2023-09-07 10:41:51','2023-09-07 10:41:51',NULL),(16,'3y4eyz09ig0cvcc7sqs226of00yplpzl','Mediant','','','72f975a31d723a3c878afdf7c529a7cf','KlPq','Mediant','',NULL,NULL,0,NULL,NULL,'normal','2023-09-07 10:42:22','2023-09-07 10:42:22',NULL),(17,'3y4eyz09ig0cvcc804niojcg003occca','Amazewor','','','244a99fb68f75f9ecfe3c88191487226','HiGy','Amazewor','',NULL,NULL,0,NULL,NULL,'normal','2023-09-07 10:42:38','2023-09-07 10:42:38',NULL),(18,'3y4eyz0f3c0cvu2kxxjcnv4n00xr1ysr','Test008','18195671381','test008@qq.com','05b52a7856b58fba230c0f2ad1ae61fd','lWsK','Test008','http://dummyimage.com/100x100',NULL,'default',0,NULL,'2023-09-28 06:57:20','normal','2023-09-28 06:57:20','2023-09-28 06:57:20',NULL);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menu`
--

DROP TABLE IF EXISTS `menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `menu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `parent_id` int unsigned NOT NULL DEFAULT '0' COMMENT '父ID',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '标题',
  `icon` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '图标',
  `cover_url` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '封面图片',
  `remark` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '描述',
  `status` enum('normal','disabled') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT 'normal' COMMENT '状态',
  `weight` int NOT NULL DEFAULT '0' COMMENT '权重',
  `create_at` datetime DEFAULT NULL COMMENT '创建日期',
  `update_at` datetime DEFAULT NULL COMMENT '更新日期',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='菜单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menu`
--

LOCK TABLES `menu` WRITE;
/*!40000 ALTER TABLE `menu` DISABLE KEYS */;
INSERT INTO `menu` VALUES (1,0,'后台系统','','','Letga后台管理系统相关导航','normal',9999,'2023-09-07 05:16:44','2023-09-07 05:16:44'),(2,1,'后台首页','','','','normal',999,'2023-09-07 05:21:01','2023-09-07 08:18:38'),(3,1,'用户管理','','','','normal',998,'2023-09-07 05:21:46','2023-09-07 05:21:46'),(4,1,'权限分组','','','','normal',997,'2023-09-07 05:22:29','2023-09-07 08:17:26'),(5,1,'媒体管理','','','','normal',996,'2023-09-07 05:22:56','2023-09-07 05:23:20'),(6,1,'菜单管理','','','','normal',995,'2023-09-07 05:23:33','2023-09-07 05:23:33'),(7,1,'系统设置','','','','normal',0,'2023-09-07 07:52:33','2023-09-07 08:53:44'),(8,4,'权限角色','','','','normal',99,'2023-09-07 08:50:31','2023-09-07 08:52:28'),(9,4,'权限路由','','','','normal',98,'2023-09-07 08:52:41','2023-09-07 08:52:41'),(10,4,'用户授权','','','','normal',97,'2023-09-07 08:52:57','2023-09-07 08:52:57'),(11,0,'前台API','','','','normal',9990,'2023-09-28 22:36:01','2023-09-28 22:36:25'),(12,0,'演示应用','','','','normal',9980,'2023-09-28 22:38:55','2023-09-28 22:38:55'),(13,11,'注册登录','','','','normal',989,'2023-09-28 22:41:13','2023-09-28 22:44:46'),(14,11,'账户中心','','','','normal',988,'2023-09-28 22:41:46','2023-09-28 22:44:54');
/*!40000 ALTER TABLE `menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `auth_access`
--

DROP TABLE IF EXISTS `auth_access`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `auth_access` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `role_id` int unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `status` enum('normal','disabled') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'normal' COMMENT '状态',
  `create_at` datetime DEFAULT NULL COMMENT '创建日期',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='授权表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `auth_access`
--

LOCK TABLES `auth_access` WRITE;
/*!40000 ALTER TABLE `auth_access` DISABLE KEYS */;
INSERT INTO `auth_access` VALUES (1,1,1,'normal','2023-08-28 20:01:01'),(2,4,2,'normal','2023-09-28 23:10:24'),(3,5,3,'normal','2023-09-28 23:11:42'),(4,3,4,'normal','2023-09-28 23:12:16'),(5,6,5,'normal','2023-09-28 23:12:38'),(6,6,6,'normal','2023-09-29 04:36:17'),(7,3,7,'normal','2023-09-29 04:36:38'),(8,6,7,'normal','2023-09-29 04:36:38');
/*!40000 ALTER TABLE `auth_access` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `auth_role_access`
--

DROP TABLE IF EXISTS `auth_role_access`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `auth_role_access` (
  `role_id` int unsigned NOT NULL,
  `route_id` int unsigned NOT NULL,
  UNIQUE KEY `auth_role_access_un` (`role_id`,`route_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `auth_role_access`
--

LOCK TABLES `auth_role_access` WRITE;
/*!40000 ALTER TABLE `auth_role_access` DISABLE KEYS */;
INSERT INTO `auth_role_access` VALUES (3,4),(4,1),(4,2),(4,3),(4,4),(4,5),(4,6),(4,25),(4,26),(4,27),(4,28),(4,29),(4,30);
/*!40000 ALTER TABLE `auth_role_access` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `auth_role`
--

DROP TABLE IF EXISTS `auth_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `auth_role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `parent_id` int unsigned NOT NULL DEFAULT '0' COMMENT '父ID',
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '标题',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '名称',
  `status` enum('normal','disabled') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'normal' COMMENT '状态',
  `weight` int NOT NULL DEFAULT '0' COMMENT '权重',
  `create_at` datetime DEFAULT NULL COMMENT '创建日期',
  `update_at` datetime DEFAULT NULL COMMENT '修改日期',
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_group_UN` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='权限角色表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `auth_role`
--

LOCK TABLES `auth_role` WRITE;
/*!40000 ALTER TABLE `auth_role` DISABLE KEYS */;
INSERT INTO `auth_role` VALUES (1,0,'超级管理员','root','normal',9999,'2023-09-04 15:14:10','2023-09-28 02:12:06'),(2,1,'默认用户','default','normal',999,'2023-09-04 15:14:50','2023-09-28 01:58:47'),(3,1,'后台用户','SubManger','normal',998,'2023-09-06 19:54:29','2023-09-28 01:58:25'),(4,3,'内容管理员','SubSubManger1','normal',0,'2023-09-08 13:17:49','2023-09-29 13:18:19'),(5,3,'系统管理员','SubSubManger2','normal',0,'2023-09-08 13:21:20','2023-09-28 02:09:55'),(6,1,'测试用户','TestUser','normal',0,'2023-09-28 02:16:33','2023-09-28 02:19:37');
/*!40000 ALTER TABLE `auth_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `auth_route`
--

DROP TABLE IF EXISTS `auth_route`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `auth_route` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `menu_id` int unsigned NOT NULL DEFAULT '0' COMMENT '菜单ID',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题',
  `path` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '路由地址',
  `method` enum('GET','POST','PUT','DELETE','PATCH') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求方法',
  `remark` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '备注',
  `status` enum('normal','disabled') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT 'normal' COMMENT '状态',
  `weight` int NOT NULL DEFAULT '0' COMMENT '权重',
  `create_at` datetime DEFAULT NULL COMMENT '创建日期',
  `update_at` datetime DEFAULT NULL COMMENT '更新日期',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='权限路由表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `auth_route`
--

LOCK TABLES `auth_route` WRITE;
/*!40000 ALTER TABLE `auth_route` DISABLE KEYS */;
INSERT INTO `auth_route` VALUES (1,3,'新增用户','/admin/user','POST','','normal',0,'2023-09-07 08:58:36','2023-09-07 08:58:36'),(2,3,'修改用户','/admin/user','PUT','','normal',0,'2023-09-07 08:59:33','2023-09-07 08:59:33'),(3,3,'查询用户','/admin/user','GET','','normal',0,'2023-09-07 09:00:15','2023-09-07 09:00:15'),(4,3,'删除用户','/admin/user','DELETE','','normal',0,'2023-09-07 09:00:33','2023-09-07 09:00:33'),(5,3,'设置用户状态','/admin/user/status','PUT','','normal',0,'2023-09-07 09:02:35','2023-09-07 09:02:35'),(6,3,'查询用户分页','/admin/users','GET','','normal',0,'2023-09-07 09:03:35','2023-09-07 09:03:35'),(7,8,'创建角色','/admin/auth/role','POST','','normal',0,'2023-09-07 09:05:49','2023-09-07 09:05:49'),(8,8,'修改角色','/admin/auth/role','PUT','','normal',0,'2023-09-07 09:06:05','2023-09-07 09:06:05'),(9,8,'查询角色','/admin/auth/role','GET','','normal',0,'2023-09-07 09:06:38','2023-09-07 09:06:38'),(10,8,'删除角色','/admin/auth/role','DELETE','','normal',0,'2023-09-07 09:07:02','2023-09-07 09:07:02'),(11,8,'设置角色状态','/admin/auth/role/status','PUT','','normal',0,'2023-09-07 09:07:26','2023-09-07 09:07:26'),(12,8,'查询角色树','/admin/auth/role/tree','GET','','normal',0,'2023-09-07 09:08:20','2023-09-07 09:08:20'),(13,9,'新增路由','/admin/auth/route','POST','','normal',0,'2023-09-07 09:09:24','2023-09-07 09:09:32'),(14,9,'修改路由','/admin/auth/route','PUT','','normal',0,'2023-09-07 09:09:49','2023-09-07 09:09:49'),(15,9,'查询路由','/admin/auth/route','GET','','normal',0,'2023-09-07 09:11:02','2023-09-07 09:11:02'),(16,9,'删除路由','/admin/auth/route','DELETE','','normal',0,'2023-09-07 09:11:46','2023-09-07 09:11:46'),(17,9,'设置路由状态','/admin/auth/route/status','PUT','','normal',0,'2023-09-07 09:12:20','2023-09-07 09:12:20'),(18,9,'路由分页','/admin/auth/routes','GET','','normal',0,'2023-09-07 09:12:55','2023-09-28 21:12:42'),(20,10,'授权设置','/admin/auth/access/setup','POST','','normal',0,'2023-09-07 09:27:10','2023-09-28 21:16:57'),(21,10,'设置授权状态','/admin/auth/access/status','PUT','','normal',0,'2023-09-07 09:29:33','2023-09-28 21:17:43'),(22,10,'删除授权','/admin/auth/access','DELETE','','normal',0,'2023-09-28 21:22:05','2023-09-28 21:22:05'),(23,10,'获取授权分页','/admin/auth/accesses','GET','','normal',0,'2023-09-28 21:22:53','2023-09-28 21:23:12'),(25,5,'上传媒体','/admin/media','POST','','normal',0,'2023-09-28 21:30:54','2023-09-28 21:30:54'),(26,5,'设置媒体状态','/admin/media/status','PUT','','normal',0,'2023-09-28 22:25:26','2023-09-28 22:25:26'),(27,5,'获取媒体','/admin/media','GET','','normal',0,'2023-09-28 22:26:11','2023-09-28 22:26:11'),(28,5,'解析媒体','/admin/media/parser','GET','','normal',0,'2023-09-28 22:28:06','2023-09-28 22:28:15'),(29,5,'删除媒体','/admin/meida','DELETE','','normal',0,'2023-09-28 22:29:17','2023-09-28 22:29:17'),(30,5,'获取媒体分页','/admin/meidas','GET','','normal',0,'2023-09-28 22:29:46','2023-09-28 22:29:46'),(31,6,'创建菜单','/admin/menu','POST','','normal',0,'2023-09-28 22:30:45','2023-09-28 22:30:45'),(32,6,'修改菜单','/admin/menu','PUT','','normal',0,'2023-09-28 22:31:11','2023-09-28 22:31:31'),(33,6,'设置才懂得状态','/admin/menu/status','PUT','','normal',0,'2023-09-28 22:32:19','2023-09-28 22:32:19'),(34,6,'获取菜单','/admin/menu','GET','','normal',0,'2023-09-28 22:32:55','2023-09-28 22:32:55'),(35,6,'删除菜单','/admin/menu','DELETE','','normal',0,'2023-09-28 22:33:20','2023-09-28 22:33:20'),(36,6,'获取菜单树','/admin/menu/tree','GET','','normal',0,'2023-09-28 22:34:36','2023-09-28 22:34:36'),(37,13,'用户注册','/api/account/signup','POST','','normal',0,'2023-09-28 22:47:05','2023-09-28 22:47:05'),(38,13,'用户登录','/api/account/signin','POST','','normal',0,'2023-09-28 22:48:28','2023-09-28 22:48:28'),(39,13,'退出登录','/api/account/signout','PUT','','normal',0,'2023-09-28 22:49:18','2023-09-28 22:57:41'),(40,13,'刷新登录','/api/account/token/refresh','PUT','','normal',0,'2023-09-28 22:58:27','2023-09-28 22:59:05'),(41,8,'获取角色过滤树','/admin/auth/role/flitertree','GET','','normal',0,'2023-09-29 04:45:53','2023-09-29 04:45:53'),(42,6,'获取菜单过滤树','/admin/menu/flitertree','GET','','normal',0,'2023-09-29 04:46:54','2023-09-29 04:46:54');
/*!40000 ALTER TABLE `auth_route` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `media`
--

DROP TABLE IF EXISTS `media`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `media` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '文件名',
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '路径',
  `size` int unsigned NOT NULL COMMENT '大小',
  `file_type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '文件类型',
  `mime_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'MIME类型',
  `hash` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '哈希值',
  `extparam` json DEFAULT NULL COMMENT '透传数据',
  `storage` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'local' COMMENT '储存库',
  `status` enum('normal','disabled') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'normal' COMMENT '状态',
  `create_at` datetime DEFAULT NULL COMMENT '创建日期',
  `update_at` datetime DEFAULT NULL COMMENT '更新日期',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除日期',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='媒体资源表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `media`
--

LOCK TABLES `media` WRITE;
/*!40000 ALTER TABLE `media` DISABLE KEYS */;
INSERT INTO `media` VALUES (1,1,'logo.png','/upload/20230904/cv9y7efhuuh4a4o5z5.png',67011,'png','image/png','f7ddf8625ff016f76917de965453ef034e66f21b97e576fb465283eb60957972',NULL,'local','normal','2023-09-04 15:18:15','2023-09-04 15:18:15',NULL);
/*!40000 ALTER TABLE `media` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-09-30 18:48:00
