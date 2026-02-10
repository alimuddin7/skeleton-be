package helpers

import (
	"test-project/constants"
	"test-project/helpers/models"
	v1Models "test-project/helpers/models" // Assuming health check data models are here
	"strings"
	"github.com/google/uuid"
	"github.com/gofiber/fiber/v3"
)

// GenerateResponseHealthCheck generates a health check response from multiple component statuses.
func GenerateResponseHealthCheck(listData ...v1Models.DataHealthCheck) models.Response {
	res := models.Response{
		Meta: GetMetaResponse(constants.RC_SUCCESS),
	}
	res.Data = listData

	for _, val := range listData {
		if val.StatusCode != 200 {
			res.Meta.Code = "400"
			res.Meta.Message = "Warning"
			return res
		}
	}
	res.Meta.Message = "Healthy"
	return res
}

// GetTraceID retrieves the trace ID from headers or generates a new one.
func GetTraceID(c fiber.Ctx) string {
	traceID := string(c.Get("X-Trace-ID"))
	if traceID == "" {
		traceID = uuid.New().String()
	}
	return traceID
}

// TelnetIP cleans and formats a URL for TCP health checks.
func TelnetIP(urlSvc string) string {
	var urlSvc1, urlSvc2, port string
	var dataURL []string
	if strings.Contains(urlSvc, "http://") {
		urlSvc1 = strings.ReplaceAll(urlSvc, "http://", "")
		dataURL = strings.Split(urlSvc1, ":")
		if len(dataURL) > 1 {
			port = dataURL[1]
		}
		if port == "" {
			port = "80"
		}
	}

	if strings.Contains(urlSvc, "https://") {
		urlSvc2 = strings.ReplaceAll(urlSvc, "https://", "")
		dataURL = strings.Split(urlSvc2, ":")
		if len(dataURL) > 1 {
			port = dataURL[1]
		}
		if port == "" {
			port = "443"
		}
	}

	if len(dataURL) == 0 {
		return urlSvc
	}

	return dataURL[0] + ":" + port
}
