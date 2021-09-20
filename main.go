package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type article struct {
	ID        int `json:"id"`
	Name      string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

var articles = []article{
	{
		ID:        1,
		Name:      "Article_1",
		CreatedAt: time.Now(),
	},
	{
		ID:        2,
		Name:      "Article_2",
		CreatedAt: time.Now().AddDate(0, 0, -2),
	},
}


func main() {
	echoApp := echo.New()

	echoApp.GET("/articles",getArticleController)
	echoApp.GET("/article",getArticleByQueryController)
	echoApp.POST("/article",createArticleController)
	echoApp.POST("/article2",createArticleController2)
	echoApp.GET("/article/:id",getArticleByIdController)

	echoApp.Start(":8080")
}

func getArticleController(ctx echo.Context) error{
	return ctx.JSON(http.StatusOK,articles)
}

func getArticleByIdController(ctx echo.Context) error{
	//url params
	id,_ := strconv.Atoi(ctx.Param("id"))
	return ctx.JSON(http.StatusOK,articles[id-1])
}

func getArticleByQueryController(ctx echo.Context) error{
	//query params
	id,_ := strconv.Atoi(ctx.QueryParam("id"))
	return ctx.JSON(http.StatusOK,articles[id-1])
}

//with form value
func createArticleController(ctx echo.Context)error{
	var newArticle article

	formID,_ := strconv.Atoi(ctx.FormValue("id"))
	formName := ctx.FormValue("name")

	newArticle.ID = formID
	newArticle.Name = formName
	newArticle.CreatedAt = time.Now()

	articles = append(articles,newArticle)

	return ctx.JSON(http.StatusOK,map[string]interface{}{
		"message": "succes create article",
		"data": newArticle,
	})
}

//with binding
func createArticleController2(ctx echo.Context)error{
	newArticle := article{}

	ctx.Bind(&newArticle)

	articles = append(articles,newArticle)

	return ctx.JSON(http.StatusOK,map[string]interface{}{
		"message": "succes create article",
		"data": newArticle,
	})
}
