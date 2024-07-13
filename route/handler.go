package route

import (
	"github.com/Bartosz-D3V/recruitment-task-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HandleGetNumber(searchSvc service.SearchSvc, context *gin.Context) {
	numberStr := context.Param("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		body := ErrorBody{
			ErrorMessage: "Provided value is not a valid number",
		}

		context.JSON(http.StatusOK, &body)
		return
	}

	index, err := searchSvc.BinarySearch(number)
	if err != nil {
		body := ErrorBody{
			ErrorMessage: "Provided value was not found",
		}

		context.JSON(http.StatusOK, &body)
		return
	}
	body := FoundIndexBody{
		FoundIndex: index,
	}
	context.JSON(http.StatusOK, &body)
}
