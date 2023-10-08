package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.File("index.html")
	})

	e.POST("/calculate", func(c echo.Context) error {
		num1, _ := strconv.ParseFloat(c.FormValue("num1"), 64)
		num2, _ := strconv.ParseFloat(c.FormValue("num2"), 64)
		operator := c.FormValue("operator")

		var result float64
		switch operator {
		case "add":
			result = num1 + num2
		case "subtract":
			result = num1 - num2
		case "multiply":
			result = num1 * num2
		case "divide":
			if num2 != 0 {
				result = num1 / num2
			} else {
				return c.HTML(http.StatusBadRequest, "Division by zero is not allowed.")
			}
		default:
			return c.HTML(http.StatusBadRequest, "Invalid operator.")
		}

		// Update the result in the HTML using JavaScript
		script := fmt.Sprintf(`
			<script>
				document.getElementById("calc-result").innerHTML = "Result: %.2f";
			</script>
		`, result)

		return c.HTML(http.StatusOK, script)
	})

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":8080"))
}
