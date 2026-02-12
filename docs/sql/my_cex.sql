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

CREATE TABLE `user_transactions` (
 `id` bigint unsigned NOT NULL AUTO_INCREMENT,
 `tx_id` varchar(64) NOT NULL DEFAULT '' COMMENT '交易流水号(唯一)',
 `uid` bigint unsigned NOT NULL COMMENT '用户ID',
 `amount` decimal(30,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT '变动金额',
 `type` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '类型 1:充值 2:提现',
 `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态 1:待处理 2:成功 3:失败',
 `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 PRIMARY KEY (`id`),
 UNIQUE KEY `idx_tx_id` (`tx_id`),
 KEY `idx_uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户账变流水表';