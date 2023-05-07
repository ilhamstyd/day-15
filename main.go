package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	connection "personal-web/connnection"
	"strconv"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AddProject struct {
	ID          int
	StartDate   string
	EndDate     string
	Description string
	Image       string
	Author      string
	Name        string
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type SessionData struct {
	Islogin bool
	Name    string
}

var userData = SessionData{}

func main() {
	connection.DatabaseConnect()

	e := echo.New()

	//serve static files from public directory
	e.Static("/public", "public")

	// initialitation to use session
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

	// ROuting
	e.GET("/Hello", helloworld)
	e.GET("/about", about)
	e.GET("/", home)
	e.GET("/contact", contactMe)
	e.GET("/addProject", addProject)
	e.GET("/add-project-detail/:id", addProjectDetail)
	e.GET("/delete-addProject/:id", deleteaddProject)
	e.GET("/form-register", formRegister)
	e.GET("/form-login", formLogin)
	e.POST("/register", register)
	e.POST("/login", login)
	e.POST("/add-add-project-detail", addaddprojectdetail)

	e.Logger.Fatal(e.Start("localhost:3000"))

}

func helloworld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func about(c echo.Context) error {
	return c.String(http.StatusOK, "ini adalah about anjay")
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		// fmt.Println("tidak ada")
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	data, _ := connection.Conn.Query(context.Background(), "SELECT id, start_date, end_date, description, image, name FROM tb_projects")
	fmt.Println(data)

	var result []AddProject

	for data.Next() {
		var each = AddProject{}

		err := data.Scan(&each.ID, &each.StartDate, &each.EndDate, &each.Description, &each.Image, &each.Name)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		each.Author = "ilham setyadji"

		result = append(result, each)
	}

	addProjects := map[string]interface{}{
		"AddProject": result,
	}

	return tmpl.Execute(c.Response(), addProjects)

	sess, _ := session.Get("session", c)

	flash := map[string]interface{}{
		"FlashStatus":  sess.Values["isLogin"],
		"FlashMessage": sess.Values["message"],
		"FlashName":    sess.Values["name"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), flash)
}

func contactMe(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/contact-me.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func addProject(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/addProject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func addProjectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tmpl, err := template.ParseFiles("views/add-project-detail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}
	var AddProjectDetail = AddProject{}

	err = connection.Conn.QueryRow(context.Background(), "SELECT id, start_date, end_date, description, image, name FROM tb_projects WHERE id = $1", id).Scan(&AddProjectDetail.ID, &AddProjectDetail.StartDate, &AddProjectDetail.EndDate, &AddProjectDetail.Description, &AddProjectDetail.Image, &AddProjectDetail.Name)

	AddProjectDetail.Author = "Ilham"

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	data := map[string]interface{}{
		"AddProject": AddProjectDetail,
	}

	return tmpl.Execute(c.Response(), data)
}

func addaddprojectdetail(c echo.Context) error {
	start_date := c.FormValue("start_date")
	end_date := c.FormValue("end_date")
	name := c.FormValue("name")
	description := c.FormValue("description")
	image := "image.png"

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_projects(start_date, end_date, name, description, image) VALUES ($1, $2, $3, $4, $5)", start_date, end_date, name, description, image)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteaddProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id")) // id = 0 string => 0 int

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_projects WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/addProject")
}

func formRegister(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/register.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func formLogin(c echo.Context) error {
	sess, _ := session.Get("session", c)
	flash := map[string]interface{}{
		"FlashStatus":  sess.Values["alertStatus"], // true / false
		"FlashMessage": sess.Values["message"],     // "Register success"
	}

	delete(sess.Values, "message")
	delete(sess.Values, "alertStatus")

	tmpl, err := template.ParseFiles("views/login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	return tmpl.Execute(c.Response(), flash)
}

func login(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	email := c.FormValue("email")
	password := c.FormValue("password")

	user := User{}
	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_user WHERE email=$1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		return redirectWithMessage(c, "Email Salah !", false, "/form-login")
	}

	fmt.Println(user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return redirectWithMessage(c, "Password Salah !", false, "/form-login")
	}

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800 // 3 jam
	sess.Values["message"] = "Login Success"
	sess.Values["status"] = true // show alert
	sess.Values["name"] = user.Name
	sess.Values["id"] = user.ID
	sess.Values["isLogin"] = true // access login
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func register(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	// generate password
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_user (name, email, password) VALUES ($1, $2, $3)", name, email, passwordHash)
	if err != nil {
		redirectWithMessage(c, "Register failed, please try again :)", false, "/form-register")
	}

	return redirectWithMessage(c, "Register success", true, "/form-login")
}

func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Values["isLogin"] = false
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func redirectWithMessage(c echo.Context, message string, status bool, path string) error {
	sess, _ := session.Get("session", c)
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, path)
}
