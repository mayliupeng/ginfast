package classes

import "github.com/gin-gonic/gin"

type IndexClass struct {
	*gin.Engine   //这玩意就是 gin.new()创建出来的
	//嵌套，好比继承，但并不是继承。go并不是纯面向对象语言
}

func NewIndexClass(e *gin.Engine) *IndexClass {  //所谓的构造函数
	return &IndexClass{Engine:e}      //需要赋值，因为是指针
}

func(this *IndexClass) GetIndex() gin.HandlerFunc{ //这就是我们的业务方法，函数名随便
	return func(context *gin.Context) {
		context.JSON(200, gin.H{
			"result": "index ok",
		})
	}	
}

func(this *IndexClass) Build() {
	this.Handle("GET", "/", this.GetIndex())
}