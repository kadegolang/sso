package user

type Roles int

const (
	// 创建者, 负责博客创作
	ROLE_CREATE Roles = iota
	// 审核员
	ROLE_AUDITOR
	// 系统管理员
	ROLE_ADMIN
)

// type QueryUserBy int64

// // 通过0 1来确定查询的id或者name条件
// const (
// 	QUERY_USER_ID QueryUserBy = iota
// 	QUERY_USER_NAME
// )
