package classes

import "github.com/gin-gonic/gin"

//注意  文件名我们以后统一使用大写（首字母大写）

type UserClass struct {
	*gin.Engine
}
//右击class自动生成所谓的构造函数
func NewUserClass(engine *gin.Engine) *UserClass {
	return &UserClass{Engine: engine}
}

func(this *UserClass) Build() {
    this.Handle("GET", "/", this.xxoo())
}