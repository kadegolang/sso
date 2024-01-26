CREATE TABLE `users` (
  `id` int(64) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL COMMENT '用户名, 用户名不允许重复的',
  `password` varchar(255) NOT NULL COMMENT '不能保持用户的明文密码',
  `label` varchar(255) NOT NULL COMMENT '用户标签',
  `role` tinyint(4) NOT NULL COMMENT '用户角色,0表示普通用户,1表示审核用户,2表示管理员',
  `create_time` int(64) NOT NULL COMMENT '创建时间',
  `update_time` int(64) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_user` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4;


CREATE TABLE `tokens` (
  `user_id` int(64) NOT NULL COMMENT '用户的Id',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名, 用户名不允许重复的',
  `access_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户的访问令牌',
  `access_token_expired_at` int(11) NOT NULL COMMENT '令牌过期时间',
  `refresh_token` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '刷新令牌',
  `refresh_token_expired_at` int(11) NOT NULL COMMENT '刷新令牌过期时间',
  `created_at` int(64) NOT NULL COMMENT '创建时间',
  `updated_at` int(64) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`access_token`) USING BTREE,
  UNIQUE KEY `idx_token` (`access_token`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


type Token struct {
	UserID                int64  `json:"user_id" gorm:"column:user_id"`                                   //该token是颁发给谁的
	UserName              string `json:"username" gorm:"column:username"`                                 //人的名字
	AccessToken           string `json:"access_token" gorm:"column:access_token"`                         //颁发给用户的访问令牌（需要用户携带token来访问接口）
	AccessTokenExpiredAt  int    `json:"access_token_expired_at" gorm:"column:access_token_expired_at"`   //过期时间
	RefreshToken          string `json:"refresh_token" gorm:"column:refresh_token"`                       //刷新令牌
	RefreshTokenExpiredAt int    `json:"refresh_token_expired_at" gorm:"column:refresh_token_expired_at"` //刷新令牌过期时间
	CreatedAt             int64  `json:"created_at" gorm:"column:created_at"`                             //创建时间
	UpdatedAt             int64  `json:"updated_at" gorm:"column:updated_at"`                             //刷新时间
	// AccessTokens          string `json:"access_tokens"`                   //颁发给用户的访问令牌（需要用户携带token来访问接口）
	// Role user.Role `gorm:"-"` // 额外补充信息, gorm忽略处理,这样就不需要存在数据库中了
}