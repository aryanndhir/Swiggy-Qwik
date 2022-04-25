package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc" //grpc
	"google.golang.org/grpc/credentials/insecure"
	_ "qwik.in/checkout/docs" //GoSwagger
	"qwik.in/checkout/protos"

	"github.com/gin-gonic/gin"
	"qwik.in/checkout/domain/services"
	"qwik.in/checkout/domain/tools/errs"
	"qwik.in/checkout/domain/tools/logger"
)

//CheckoutHandler ..
type CheckoutHandler struct {
	CheckoutService services.CheckoutService
}

//PaymentModesDTO ..
type PaymentModesDTO struct {
	Mode       string `json:"mode" dynamodbav:"mode"`
	CardNumber int    `json:"cardNumber" dynamodbav:"card_number"`
}

// CheckoutGetShippingAddressFlow ..
// @Summary      Gets Default Shipping Address
// @Description  Returns default Shipping Address
// @Tags         ShippingAddress Service
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /checkout/api/shippingaddress [get]
func (ch CheckoutHandler) CheckoutGetShippingAddressFlow() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("⌛ Getting Shipping Address !")
		var shippingAddressRespBody ShippingAddrDTO
		////INITIAL GET REQUEST
		req, err := http.Get("http://localhost:9002/shipping/api/existing")
		//ADDING HEADER FOR REQUEST
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		//HANDLING ERROR / ELSE -> CONSUME REQUEST
		if err != nil {
			logger.Error("Default Shipping Address does not exist")
			ctx.JSON(http.StatusBadRequest, errs.NewBadRequest("Shipping Address does not exist"))
			return
		}
		data, err2 := ioutil.ReadAll(req.Body)
		if err2 != nil {
			logger.Error("Error during RESPONSE value from GET Request")
			ctx.JSON(http.StatusBadRequest, errs.NewBadRequest("Some shit doesn't work during GET - CONSUME"))
		}
		err3 := json.Unmarshal(data, &shippingAddressRespBody)
		if err3 != nil {
			ctx.JSON(http.StatusBadRequest, errs.NewBadRequest("Default Shipping Address couldn't be UNMARSHALLED."))
			return
		}
		fmt.Println("✅ Default ShippingAddress Fetched ")
		logger.Info("✅ Default ShippingAddress Fetched ", shippingAddressRespBody)
		ctx.JSON(http.StatusAccepted, shippingAddressRespBody)

	}
}

// CheckoutGetCartItemsFlow ..
// @Summary      Gets Cart Items
// @Description  Returns Cart Items using GET Request.
// @Tags         Cart Service
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /checkout/api/cartItems   [get]
func (ch CheckoutHandler) CheckoutGetCartItemsFlow() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("⌛ Getting Cart Items !")
		var cartItemsRespBody CartDTO
		////INITIAL GET REQUEST
		req, err := http.Get("http://localhost:9002/cart/api/cartItems")
		//ADDING HEADER FOR REQUEST
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		//HANDLING ERROR
		if err != nil {
			logger.Error("Cart Items do not exist")
			ctx.JSON(http.StatusBadRequest, errs.NewBadRequest("Cart Items do not exist"))
			return
		}
		data, err2 := ioutil.ReadAll(req.Body)
		if err2 != nil {
			logger.Error("Error during RESPONSE value from GET Request")
			ctx.JSON(http.StatusBadRequest, errs.NewBadRequest("Some shit doesn't work during GET - CONSUME"))
		}
		err3 := json.Unmarshal(data, &cartItemsRespBody)
		if err3 != nil {
			ctx.JSON(http.StatusBadRequest, errs.NewBadRequest("Cart Items couldn't be UNMARSHALLED."))
			return
		}

		fmt.Println("✅ Cart Items Fetched")
		logger.Info("✅ Cart Items Fetched ", cartItemsRespBody)
		ctx.JSON(http.StatusAccepted, cartItemsRespBody)

	}
}

// CheckoutGetPaymentsFlow ..
// @Summary      Gets Payment Mode
// @Description  Returns Payment Mode using GET Request.
// @Tags         Payment Mode Service
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /checkout/api/payments   [get]
func (ch CheckoutHandler) CheckoutGetPaymentsFlow() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("⌛ Getting Payment Mode Items !")
		var paymentModeResponseBody PaymentModeDTO
		////INITIAL GET REQUEST
		req, err := http.Get("http://localhost:9002/payment/api/payments")
		//ADDING HEADER FOR REQUEST
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		//HANDLING ERROR
		if err != nil {
			logger.Error("Payment Mode do not exist")
			ctx.JSON(http.StatusBadRequest, errs.NewBadRequest("Payment Mode do not exist"))
			return
		}
		data, err2 := ioutil.ReadAll(req.Body)
		if err2 != nil {
			logger.Error("Error during RESPONSE value from GET Request")
			ctx.JSON(http.StatusBadRequest, errs.NewBadRequest("Some shit doesn't work during GET - CONSUME"))
		}
		err3 := json.Unmarshal(data, &paymentModeResponseBody)
		if err3 != nil {
			ctx.JSON(http.StatusBadRequest, errs.NewBadRequest("Payment Mode couldn't be UNMARSHALLED."))
			return
		}

		fmt.Println("✅ Payment Mode Fetched")
		logger.Info("✅ Payment Mode Fetched", paymentModeResponseBody)
		ctx.JSON(http.StatusAccepted, paymentModeResponseBody)

	}
}

// CheckoutPayStatusFlow ..
// @Summary      Gets Payment Mode
// @Description  Returns Payment Mode using GET Request.
// @Tags         Payment Mode Service
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /checkout/api/pay  [get]
func (ch CheckoutHandler) CheckoutPayStatusFlow() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		readData, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			logger.Error("Empty Request Body", err)
		}
		paymentModesDTOData := PaymentModesDTO{}
		_ = json.Unmarshal(readData, &paymentModesDTOData)
		status, result, err := IsCheckoutPaymentSuccessful(paymentModesDTOData)

		if status {
			ctx.JSON(http.StatusAccepted, result)
		} else {
			ctx.JSON(http.StatusInternalServerError, err)
		}
	}
}

// CheckoutShippingAddressFlow ..
// @Summary      Gets Payment Mode
// @Description  Returns Payment Mode using GET Request.
// @Tags         Payment Mode Service
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /checkout/api/existing  [get]
func CheckoutShippingAddressFlow() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		readData, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			logger.Error("Empty Request Body", err)
		}
		paymentModesDTOData := PaymentModesDTO{}
		_ = json.Unmarshal(readData, &paymentModesDTOData)
		conn, err := grpc.Dial("localhost:9003", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Connection failed : %v", err)
		}
		defer func(conn *grpc.ClientConn) {
			err := conn.Close()
			if err != nil {
				logger.Error("Error : ", err)
			}
		}(conn)
		c := protos.NewShippingAddressProtoFuncClient(conn)
		shippingRequest := &protos.ShippingAddressRequest{
			DefaultAddress: true,
		}
		result, err := c.GetDefaultShippingAddress(ctx, shippingRequest)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			log.Printf("Shipping failed : %v", err)
		} else {
			fmt.Println(result.ShippingAddress)
			ctx.JSON(http.StatusAccepted, result.ShippingAddress)
		}
	}
}

//IsCheckoutPaymentSuccessful ..
func IsCheckoutPaymentSuccessful(userPaymentData PaymentModesDTO) (bool, *protos.PaymentResponse, error) {
	conn, err := grpc.Dial("localhost:9004", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Connection failed : %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			logger.Error("Error : ", err)
		}
	}(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c := protos.NewPaymentClient(conn)
	paymentRequest := &protos.PaymentRequest{
		UserId:  "bb912edc-50d9-42d7-b7a1-9ce66d459abcd",
		Amount:  1,
		OrderId: "OA-123",
		PaymentMode: &protos.PaymentMode{
			Mode:       userPaymentData.Mode,
			CardNumber: int64(userPaymentData.CardNumber),
		},
	}
	//fmt.Println("hello")
	result, err := c.CompletePayment(ctx, paymentRequest)
	if err != nil {
		//ctx.JSON(http.StatusInternalServerError, err)
		logger.Error("Payment failed : %v", err)
		return false, nil, err
	} else {
		logger.Info("Payment Successful !", result.GetIsPaymentSuccessful())
		return true, result, nil
	}
}
