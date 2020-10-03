package handler

import (
	"net/http"
	
	service "OrderManagerHex/service"

	model "OrderManagerHex/domain/model"

	"github.com/gin-gonic/gin"
	
    "github.com/gin-contrib/cors"
	"strconv"
	"fmt"
)


type OrderHandler interface {
	MainHttpd()
	ShowOrders() gin.HandlerFunc
	AddOrders() gin.HandlerFunc
}

type orderHandler struct {
	orderManagementService service.OrderManagementService
}

func NewOrderHandler(cs service.OrderManagementService) OrderHandler {
	return &orderHandler{
		orderManagementService :cs,
	}
}


func (ch orderHandler) MainHttpd(){
	r := gin.Default()

    // CORS 対応
    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"*"}
    r.Use(cors.New(config))

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
	})
	
	r.POST("/showoders", ch.ShowOrders())
	r.POST("/addorders", ch.AddOrders())

    r.Run(":8081")
}

type GetItemsRequest struct {
    User string `json:"user"`
    OrderDate string `json:"order_date"`
}

func (ch orderHandler) ShowOrders() gin.HandlerFunc {
    return func(c *gin.Context) {
        requestBody := GetItemsRequest{}
        c.Bind(&requestBody)
		result, err := ch.orderManagementService.ShowOrders(requestBody.User, requestBody.OrderDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, result)
		}
        c.JSON(http.StatusOK, result)
    }
}

type AddOrderRequest struct {
	User string `json:"user"`
	Item string `json:"item"`
	Price string `json:"price"`
	Quantity string `json:"quantity"`
}

func (ch orderHandler)AddOrders() gin.HandlerFunc {
    return func(c *gin.Context) {
        requestBody := []AddOrderRequest{}
        c.Bind(&requestBody)
        orders := []model.Order{}
        for _, v := range requestBody {   
            p, err := strconv.ParseInt(v.Price, 10, 64)    
            if err != nil {
                fmt.Println(err)
                c.Status(http.StatusBadRequest)
                return
            }
            q, err := strconv.ParseInt(v.Quantity, 10, 64)    
            if err != nil {
                fmt.Println(err)
                c.Status(http.StatusBadRequest)
                return
            }

            order := model.Order{
                User: &v.User,
                Item: &v.Item,
                Price: &p,
                Quantity: &q,
            }
            orders = append(orders, order)
        }
        ch.orderManagementService.AddOrder(&orders)

        c.Status(http.StatusNoContent)
    }
}