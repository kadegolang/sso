package tools

import "golang.org/x/crypto/bcrypt"

// 密码hash加密
func HashPassword(password string) string {
	pass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pass)
}

// 检查密码
func CheckPassword(userpassword, password string) error {
	// fmt.Println("password", password)
	// fmt.Println("password", string(password))
	// fmt.Println("u.Password", string(u.Password))
	return bcrypt.CompareHashAndPassword([]byte(userpassword), []byte(password))
}
