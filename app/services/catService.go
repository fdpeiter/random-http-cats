package services

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"math/rand"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/fdpeiter/random-http-cats/app/logger"

	"github.com/gin-gonic/gin"
)

type CatService struct {
	Logger *logger.Logger
}

func NewCatService(logger *logger.Logger) *CatService {
	return &CatService{Logger: logger}
}

func GetValidHttpStatusCodes() []int {
	return []int{200, 201, 202, 204, 206, 300, 301, 302, 303, 304, 307, 400, 401, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 500, 501, 502, 503, 504, 505}
}

func (cs *CatService) randomHTTPStatusCode() int {
	codes := GetValidHttpStatusCodes()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return codes[rand.Intn(len(codes))]
}

func (cs *CatService) ServeRandomCat(c *gin.Context) {
	statusCode := cs.randomHTTPStatusCode()
	resp, err := http.Get(fmt.Sprintf("https://http.cat/%d", statusCode))
	if err != nil {
		cs.Logger.WithError(err).Error("Error fetching image")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		cs.Logger.WithError(err).Error("Error decoding image")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		cs.Logger.WithError(err).Error("Unable to encode image.")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	cs.Logger.Info(fmt.Sprintf("Successfully generated random cat with status code %d", statusCode))

	c.Data(http.StatusOK, "image/jpeg", buffer.Bytes())
}

func (cs *CatService) ServeSpecificCat(c *gin.Context, statusCode string) {
	codes := GetValidHttpStatusCodes()
	intStatusCode, err := strconv.Atoi(statusCode)
	if err != nil {
		cs.Logger.Error(fmt.Sprintf("Unable to convert status code to number: %s", statusCode))
		c.String(http.StatusBadRequest, fmt.Sprintf("Unable to convert status code to number: %s", statusCode))
		return

	}
	if !slices.Contains(codes, intStatusCode) {
		cs.Logger.Error(fmt.Sprintf("Invalid status code: %s", statusCode))
		c.String(http.StatusBadRequest, fmt.Sprintf("Invalid status code: %s", statusCode))
		return

	}

	resp, err := http.Get(fmt.Sprintf("https://http.cat/%s", statusCode))
	if err != nil {
		cs.Logger.WithError(err).Error("Error fetching image")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		cs.Logger.WithError(err).Error("Error decoding image")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		cs.Logger.WithError(err).Error("Unable to encode image.")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	cs.Logger.Info(fmt.Sprintf("Successfully generated specific cat for: %s", statusCode))

	c.Data(http.StatusOK, "image/jpeg", buffer.Bytes())
}
