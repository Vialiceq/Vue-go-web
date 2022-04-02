package db

import (
	"fmt"
	"net/http"
	"server/models"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
	dbName: = os.Getenv("ENV_DBNAME")
	userName: = os.Getenv("ENV_DBUSER")
	password: = os.Getenv("ENV_DBPASS")
	dbHost: = os.Getenv("ENV_DBHOST")
	mongoDialInfo: = & mgo.DialInfo {
	Addrs: [] string {
	dbHost
	},
	Database: dbName,
	Username: userName,
	Password: password,
	Timeout: 60 * time.Second,
	}
*/
type categories struct {
	Name string `json:"name"`
}

func insert_t(ctx *gin.Context) {
	// 连接mongodb服务
	url := "mongodb://localhost:27017/db_go"
	// 设置数据库一致性模式
	// 连接数据库操作，该操作赋值给session
	// err值必写，用于错误处理
	session, err := mgo.Dial(url)
	// 后边程序执行的err与go程序比对，若有错误则返回错误内容
	if err != nil {
		panic(err)
	} else {
		// 若没有错误，则在页面返回字符串，显示插入成功
		ctx.String(http.StatusOK, "插入成功")
	}
	// defer用法大家自行百度，我解释不清
	defer session.Close()

	// 设置数据库一致性模式，就当作打开数据库
	session.SetMode(mgo.Monotonic, true)
	// 找到某数据库下的某数据表
	c := session.DB("db_go").C("categories")

	// 接收传值，将传值添加到上方定义的数据表结构中
	var form categories
	// 如果传值格式不符合上方定义的结构，则返回错误信息
	if err := ctx.Bind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 此时&form就继承了上方定义的结构格式
	// 插入数据，并将insert状态传值给err
	err = c.Insert(&form)
}

//type category m.Category

const url = "mongodb://localhost:27017/db_go"

// 插入数据
func insert(ctx *gin.Context) {

	// 解析api参数
	resource := ctx.Param("resource")

	// 设置数据库一致性模式
	// 连接数据库操作，该操作赋值给session
	// err值必写，用于错误处理
	session, err := mgo.Dial(url)
	// 后边程序执行的err与go程序比对，若有错误则返回错误内容
	if err != nil {
		panic(err)
	} else {
		// 若没有错误，则在页面返回字符串，显示插入成功
		ctx.String(http.StatusOK, "插入成功")
	}

	defer session.Close()

	// 设置数据库一致性模式，就当作打开数据库
	session.SetMode(mgo.Monotonic, true)
	// 找到某数据库下的某数据表
	c := session.DB("db_go").C(resource)
	// 以上为连接数据库

	if resource == "user" {
		type modelName models.User
		// 使用user数据模型
		var form modelName

		// 如果传值格式不符合上方定义的结构，则返回错误信息
		if err := ctx.Bind(&form); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 此时&form就继承了上方定义的结构格式
		// 插入数据，并将insert状态传值给err
		err = c.Insert(&models.User{"mike", "alice"})
	} else if resource == "category" {
		type modelName models.Category
		// 使用user数据模型
		var form modelName

		// 如果传值格式不符合上方定义的结构，则返回错误信息
		if err := ctx.Bind(&form); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(err)
		fmt.Println(form)
		fmt.Println("here")
		// 此时&form就继承了上方定义的结构格式
		// 插入数据，并将insert状态传值给err
		err = c.Insert(&form)
	}

	fmt.Println(err)
	// ctx.String(http.StatusOK, fmt.Sprintf(resource))
}

// 查询数据
func find(ctx *gin.Context) {

	// 解析api参数
	resource := ctx.Param("resource")
	id := ctx.Param("id")
	// 传来的id值为"/id"，我们要把"/"截去
	id = strings.Trim(id, "/")
	// 将id转化为bson.ObjectId格式
	//var _id bson.ObjectId
	var _id = bson.ObjectIdHex(id)

	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("db_go").C(resource)
	// 以上为数据库连接

	if resource == "user" {
		type modelName models.User

		modelNamer := modelName{}
		// 查找数据
		err = c.Find(bson.M{"_id": _id}).One(&modelNamer)
		ctx.JSON(http.StatusOK, modelNamer)
	} else if resource == "category" {
		type modelName models.Category

		modelNamer := modelName{}
		// 查找数据
		err = c.Find(bson.M{"_id": _id}).One(&modelNamer)
		ctx.JSON(http.StatusOK, modelNamer)
	}
	fmt.Println(err)

}

// 查询全部数据
func findAll(ctx *gin.Context) {

	// 解析api参数
	resource := ctx.Param("resource")

	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("db_go").C(resource)

	if resource == "user" {
		// 使用user数据模型
		type modelName models.User
		// 查找10条数据
		modelNames := make([]modelName, 10)

		// 查找全部
		err = c.Find(nil).All(&modelNames)

		// 返回数据
		ctx.JSON(http.StatusOK, modelNames)
	} else if resource == "category" {
		// 使用user数据模型
		type modelName models.Category
		// 查找10条数据
		modelNames := make([]modelName, 10)

		// 查找全部
		err = c.Find(nil).All(&modelNames)

		// 返回数据
		ctx.JSON(http.StatusOK, modelNames)
	}
	fmt.Println(err)
}

// 删除数据
func delete(ctx *gin.Context) {

	// 解析api参数
	resource := ctx.Param("resource")
	id := ctx.Param("id")
	// 传来的id值为"/id"，我们要把"/"截去
	id = strings.Trim(id, "/")
	// 将id转化为bson.ObjectId格式
	var _id = bson.ObjectIdHex(id)

	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	} else {
		ctx.String(http.StatusOK, id)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("db_go").C(resource)

	// 根据获取到的id删除内容
	err = c.Remove(bson.M{"_id": _id})
	fmt.Println(err)
}

// 修改数据
func update(ctx *gin.Context) {

	// 解析api参数
	resource := ctx.Param("resource")
	id := ctx.Param("id")
	// 传来的id值为"/id"，我们要把"/"截去
	id = strings.Trim(id, "/")
	// 将id转化为bson.ObjectId格式

	var _id = bson.ObjectIdHex(id)

	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	} else {
		ctx.String(http.StatusOK, "修改成功")
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("db_go").C(resource)
	// 以上为连接数据库

	if resource == "user" {
		type modelName models.User
		// 使用user数据模型
		var form modelName
		// 合并数据，如果传值格式不符合上方定义的结构，则返回错误信息
		if err := ctx.Bind(&form); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 如果传值格式不符合上方定义的结构，则返回错误信息
		err = c.Update(bson.M{"_id": _id}, &form)
		ctx.JSON(http.StatusOK, &form)
	} else if resource == "category" {
		type modelName models.Category
		// 使用user数据模型
		var form modelName
		// 合并数据，如果传值格式不符合上方定义的结构，则返回错误信息
		if err := ctx.Bind(&form); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 如果传值格式不符合上方定义的结构，则返回错误信息
		err = c.Update(bson.M{"_id": _id}, &form)
		ctx.JSON(http.StatusOK, &form)
	}
	fmt.Println(err)

}

// middleware
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")                                                                                         // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")                              //header的类型
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                                                          //允许请求方法
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type") //返回数据格式
		c.Header("Access-Control-Allow-Credentials", "true")                                                                                 //设置为true，允许ajax异步请求带cookie信息

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

// Main 配置路由
func Main(e *gin.Engine) {
	e.Use(Cors())
	// 定义路由，调用接口函数
	// 增
	e.POST("/admin/api/rest/:resource", insert)
	//e.POST("/admin/api/rest", insert_t)
	// 删
	e.DELETE("/admin/api/rest/:resource/*id", delete)
	// 改
	e.PUT("/admin/api/rest/:resource/*id", update)
	// 查
	e.GET("/admin/api/rest/:resource", findAll)
	// 根据id查某数据
	e.GET("/admin/api/rest/:resource/*id", find)

}
