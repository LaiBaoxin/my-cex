CREATE DATABASE IF NOT EXISTS `my_cex` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE `my_cex`;

CREATE TABLE `user_assets` (
   `id` bigint unsigned NOT NULL AUTO_INCREMENT,
   `uid` bigint unsigned NOT NULL COMMENT '用户ID',
   `address` varchar(64) NOT NULL DEFAULT '' COMMENT '钱包地址 0x...',
   `currency` varchar(10) NOT NULL DEFAULT 'ETH' COMMENT '币种',
   `balance` decimal(30,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT '余额',
   `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态 1:启用 2:冻结',
   `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
   `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   PRIMARY KEY (`id`),
   UNIQUE KEY `idx_uid_currency` (`uid`, `currency`),
   UNIQUE KEY `idx_address` (`address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户资产表';