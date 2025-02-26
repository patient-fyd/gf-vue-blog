-- 用户表 (users)
CREATE TABLE `users` (
                         `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                         `username` VARCHAR(50) NOT NULL COMMENT '用户名，唯一',
                         `password` VARCHAR(255) NOT NULL COMMENT '用户密码',
                         `email` VARCHAR(100) DEFAULT NULL COMMENT '用户邮箱，唯一',
                         `created_at` DATETIME(3) DEFAULT NULL COMMENT '用户创建时间',
                         `updated_at` DATETIME(3) DEFAULT NULL COMMENT '用户更新时间',
                         `role` VARCHAR(20) NOT NULL DEFAULT 'user' COMMENT '用户角色，默认是普通用户',
                         `reset_token` VARCHAR(64) DEFAULT NULL COMMENT '密码重置令牌',
                         `reset_token_expiry` DATETIME(3) DEFAULT NULL COMMENT '密码重置令牌过期时间',
                         `verification_code` VARCHAR(6) DEFAULT NULL COMMENT '邮箱验证码',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `idx_users_username` (`username`),
                         UNIQUE KEY `idx_users_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 角色表 (roles)
CREATE TABLE `roles` (
                         `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                         `role_name` VARCHAR(50) NOT NULL COMMENT '角色名称',
                         `created_at` DATETIME(3) DEFAULT NULL COMMENT '创建时间',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `idx_roles_role_name` (`role_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 权限表 (permissions)
CREATE TABLE `permissions` (
                               `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                               `permission_name` VARCHAR(100) NOT NULL COMMENT '权限名称',
                               `created_at` DATETIME(3) DEFAULT NULL COMMENT '创建时间',
                               PRIMARY KEY (`id`),
                               UNIQUE KEY `idx_permissions_permission_name` (`permission_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 角色权限关联表 (role_permissions)
CREATE TABLE `role_permissions` (
                                    `role_id` INT UNSIGNED NOT NULL,
                                    `permission_id` INT UNSIGNED NOT NULL,
                                    PRIMARY KEY (`role_id`, `permission_id`),
                                    FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE,
                                    FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 用户角色关联表 (user_roles)
CREATE TABLE `user_roles` (
                              `user_id` INT UNSIGNED NOT NULL,
                              `role_id` INT UNSIGNED NOT NULL,
                              PRIMARY KEY (`user_id`, `role_id`),
                              FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
                              FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 分类表 (categories)
CREATE TABLE `categories` (
                              `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                              `name` VARCHAR(100) NOT NULL COMMENT '分类名称',
                              `description` TEXT DEFAULT NULL COMMENT '分类描述',
                              `created_at` DATETIME(3) DEFAULT NULL COMMENT '创建时间',
                              `updated_at` DATETIME(3) DEFAULT NULL COMMENT '更新时间',
                              PRIMARY KEY (`id`),
                              UNIQUE KEY `idx_categories_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 标签表 (tags)
CREATE TABLE `tags` (
                        `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                        `name` VARCHAR(100) NOT NULL COMMENT '标签名称',
                        `created_at` DATETIME(3) DEFAULT NULL COMMENT '创建时间',
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `idx_tags_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 文章表 (posts)
CREATE TABLE `posts` (
                         `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                         `title` VARCHAR(255) NOT NULL COMMENT '文章标题',
                         `content` TEXT NOT NULL COMMENT '文章内容',
                         `user_id` INT UNSIGNED NOT NULL COMMENT '作者ID',
                         `category_id` INT UNSIGNED NOT NULL COMMENT '分类ID',
                         `created_at` DATETIME(3) DEFAULT NULL COMMENT '创建时间',
                         `updated_at` DATETIME(3) DEFAULT NULL COMMENT '更新时间',
                         `status` ENUM('draft', 'published', 'archived') NOT NULL DEFAULT 'draft' COMMENT '文章状态',
                         PRIMARY KEY (`id`),
                         FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
                         FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 文章标签表 (post_tags)
CREATE TABLE `post_tags` (
                             `post_id` INT UNSIGNED NOT NULL COMMENT '文章ID',
                             `tag_id` INT UNSIGNED NOT NULL COMMENT '标签ID',
                             PRIMARY KEY (`post_id`, `tag_id`),
                             FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`) ON DELETE CASCADE,
                             FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 签到记录表 (check_ins)
CREATE TABLE `check_ins` (
                             `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                             `user_id` INT UNSIGNED NOT NULL COMMENT '用户ID',
                             `check_in_date` DATE NOT NULL COMMENT '签到日期',
                             `created_at` DATETIME(3) DEFAULT NULL COMMENT '创建时间',
                             PRIMARY KEY (`id`),
                             FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 评论表 (comments)
CREATE TABLE `comments` (
                            `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                            `post_id` INT UNSIGNED NOT NULL COMMENT '文章ID',
                            `user_id` INT UNSIGNED NOT NULL COMMENT '用户ID',
                            `content` TEXT NOT NULL COMMENT '评论内容',
                            `created_at` DATETIME(3) DEFAULT NULL COMMENT '创建时间',
                            PRIMARY KEY (`id`),
                            FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`) ON DELETE CASCADE,
                            FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 操作日志表 (logs)
CREATE TABLE `logs` (
                        `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                        `user_id` INT UNSIGNED NOT NULL COMMENT '用户ID',
                        `action` VARCHAR(255) NOT NULL COMMENT '操作类型',
                        `description` TEXT DEFAULT NULL COMMENT '操作描述',
                        `created_at` DATETIME(3) DEFAULT NULL COMMENT '操作时间',
                        PRIMARY KEY (`id`),
                        FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;