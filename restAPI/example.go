package restapi

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

type UserHandler struct {
	db *gorm.DB
}

func main() {
	r := gin.Default()
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", "postgres", "postgres", "postgres", "5432", "golang101")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("connect database failed %v", err))
	}

	db.AutoMigrate(&User{})
	db.Create(&User{Name: "monkey"})

	userHandler := UserHandler{db: db}
	r.GET("/", testHandler)
	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message" : "test",
	// 	})
	// })
	r.GET("/user", userHandler.GetUser)
	// r.GET("/user", func(ctx *gin.Context) {
	// 	var u User
	// 	db.First(&u)
	// 	ctx.JSON(200, gin.H{
	// 		"message": u,
	// 	})
	// })
	r.GET("/users", userHandler.GetListUser)
	r.POST("/user", userHandler.CreateUser)
	r.DELETE("/user/:id", userHandler.Remove)
	r.PUT("/user/:id", userHandler.Update)

	r.Run(":8080")
}

func testHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "test",
	})
}

func (h UserHandler) GetUser(ctx *gin.Context) {
	var u User
	h.db.First(&u)
	ctx.JSON(200, gin.H{
		"message": u,
	})
}

func (h UserHandler) GetListUser(c *gin.Context) {
	var u []User
	h.db.Find(&u)
	c.JSON(200, u)
}

func (h UserHandler) CreateUser(c *gin.Context) {
	var a User
	if err := c.ShouldBindJSON(&a); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
			})
			return
	}

	r := h.db.Create(&a)
	if err := r.Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
			})
			return
	}

	c.JSON(http.StatusOK, gin.H{
			"ID": a.Model.ID,
			"Name": a.Name,
	})
}

func (h UserHandler) Remove(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
			})
			return
	}

	r := h.db.Delete(&User{}, id)
	if err = r.Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
			})
			return
	}

	c.JSON(http.StatusOK, gin.H{
			"status": "success",
	})
}

func (h UserHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
									"error": "Invalid ID",
					})
					return
	}

	var u User
	if err := h.db.First(&u, id).Error; err != nil {
					c.JSON(http.StatusNotFound, gin.H{
									"error": "User not found",
					})
					return
	}

	if err := c.ShouldBindJSON(&u); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
									"error": err.Error(),
					})
					return
	}

	if err := h.db.Save(&u).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
									"error": err.Error(),
					})
					return
	}

	c.JSON(http.StatusOK, gin.H{
					"status": "success",
					"user":   u,
	})
}
