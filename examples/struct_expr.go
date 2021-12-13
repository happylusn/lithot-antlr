package examples

import (
	"fmt"
	"github.com/happylusn/lithot-antlr/src/expr"
)

type UserRole struct {
	RoleName string
}

func NewUserRole(role string) *UserRole {
	return &UserRole{role}
}

func (this *UserRole) GetRole(prefix string) string {
	return prefix + ":" + this.RoleName
}

type User struct {
	Name string
	Role *UserRole
}

func (this *User) GetName() string {
	return this.Name
}
func (this *User) GetRoleName(role *UserRole) string {
	return role.RoleName
}

//初始化用户实体
func NewUser(name string, role string) *User {
	return &User{Name: name, Role: &UserRole{RoleName: role}}
}

func StructExpr() {
	exprMap2 := map[string]interface{}{
		"NewUserRole": NewUserRole,
		"user":        NewUser("jtthink", "admin"),
	}
	fmt.Println(expr.BeanExpr("user.GetName()", exprMap2))             //方法名 大小写敏感
	fmt.Println(expr.BeanExpr("user.Role.GetRole('当前角色是')", exprMap2)) //方法名 大小写敏感
	//fmt.Println(expr.BeanExpr("user.GetRoleName(NewUserRole('guest'))", exprMap2))
}
