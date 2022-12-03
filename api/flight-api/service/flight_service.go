package flight_handler

import (
	flight_request "mock-golang/api/flight-api/request"
	flight_response "mock-golang/api/flight-api/response"
	"mock-golang/protobuf"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlightHandler interface {
	CreateFlight(c *gin.Context)
	UpdateFlight(c *gin.Context)
	SearchFlight(c *gin.Context)
	SearchFlightById(c *gin.Context)
}

type flightHandler struct {
	flightClient protobuf.RPCFlightClient
}

func (h *flightHandler) SearchFlightById(c *gin.Context) {
	id := c.Param("id")
	if len(id) < 1 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  "id invalid",
		})

		return
	}
	pReqFlight := &protobuf.FlightParamId{
		Id: id,
	}
	pResFlight, err := h.flightClient.FindById(c.Request.Context(), pReqFlight)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": pResFlight,
	})
}

func NewFlightHandler(flightClient protobuf.RPCFlightClient) FlightHandler {
	return &flightHandler{
		flightClient: flightClient,
	}
}

func (h *flightHandler) CreateFlight(c *gin.Context) {
	req := flight_request.CreateFlightRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}

	DepartDateTime, _ := time.Parse("2006/01/02 15:04:05", req.DepartDate+" "+req.DepartTime)

	pReq := &protobuf.Flight{
		Name:          req.Name,
		From:          req.From,
		To:            req.To,
		DepartDate:    timestamppb.New(DepartDateTime),
		Status:        req.Status,
		AvailableSlot: req.AvailableSlot,
	}

	pRes, err := h.flightClient.CreateFlight(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &flight_response.CreateFlightResponse{
		Id:            pRes.Id,
		Name:          pRes.Name,
		From:          pRes.From,
		To:            pRes.To,
		DepartDate:    pRes.DepartDate.AsTime().Format("2006/01/02"),
		DepartTime:    pRes.DepartDate.AsTime().Format("15:04:05"),
		Status:        pRes.Status,
		AvailableSlot: pRes.AvailableSlot,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *flightHandler) UpdateFlight(c *gin.Context) {
	req := flight_request.UpdateFlightRequest{}
	//fmt.Printf("id", c.Param("id"))
	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}

	pReq := &protobuf.Flight{
		Id:            req.Id,
		Name:          req.Name,
		From:          req.From,
		To:            req.To,
		Status:        req.Status,
		AvailableSlot: req.AvailableSlot,
	}

	if len(strings.TrimSpace(req.DepartDate)) > 0 && len(strings.TrimSpace(req.DepartTime)) > 0 {
		DepartDateTime, _ := time.Parse("2006/01/02 15:04:05", req.DepartDate+" "+req.DepartTime)
		pReq.DepartDate = timestamppb.New(DepartDateTime)
	}

	pRes, err := h.flightClient.UpdateFlight(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": ToApiResponse(pRes),
	})
}

func (h *flightHandler) SearchFlight(c *gin.Context) {

	req := flight_request.SearchFlightRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}

	// pFromDate, _ := time.Parse("2006/01/02 15:04:05", req.FromDate)
	// pToDate, _ := time.Parse("2006/01/02 15:04:05", req.ToDate)

	pReq := &protobuf.SearchFlightRequest{
		Name: req.Name,
		From: req.From,
		To:   req.To,
		// FromDate: timestamppb.New(pFromDate),
		// ToDate:   timestamppb.New(pToDate),
	}

	pRes, err := h.flightClient.SearchFlight(c.Request.Context(), pReq)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dtos := make([]*flight_response.FlightResponse, 0)

	for _, v := range pRes.Flight {
		dto := ToApiResponse(v)

		dtos = append(dtos, dto)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dtos,
	})
}

func ToApiResponse(pRes *protobuf.Flight) *flight_response.FlightResponse {
	res := &flight_response.FlightResponse{
		Id:            pRes.Id,
		Name:          pRes.Name,
		From:          pRes.From,
		To:            pRes.To,
		Status:        pRes.Status,
		AvailableSlot: pRes.AvailableSlot,
		DepatureDate:  pRes.DepartDate.AsTime().Format("2006-01-02 15:04:05"),
		CreatedAt:     pRes.CreatedAt.AsTime().Format("2006-01-02 15:04:05"),
		UpdatedAt:     pRes.UpdatedAt.AsTime().Format("2006-01-02 15:04:05"),
	}

	return res
}
