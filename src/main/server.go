package main

import (
	"net/http"
	"github.com/labstack/echo"
)

type User struct {
	Name string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
	e := echo.New()	
	e.POST("/api/users", save)
	e.GET("/api/users", findAll)
	e.GET("/api/users/:id", findById)
    e.PUT("/api/users/:id", updateById)
	e.DELETE("/api/users/:id", deleteById)
	e.Logger.Fatal(e.Start(":8081"))
}

func save(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil{
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func findAll(c echo.Context) error {
	user := new(User)
	user.Name  = "jaruwit"
	user.Email = "kingbom_300432@hotmail.com"
	return c.JSON(http.StatusOK, user)
}

func findById(c echo.Context) error {
	 id := c.Param("id")
	 return c.JSON(http.StatusOK, id)
}

func updateById(c echo.Context) error {
    id := c.Param("id")
	 return c.JSON(http.StatusOK, id)
}

func deleteById(c echo.Context) error {
    id := c.Param("id")
	 return c.JSON(http.StatusOK, id)
}