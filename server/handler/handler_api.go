package handler

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/yunussandikci/shareterm/common"
	service2 "github.com/yunussandikci/shareterm/common/service"
	"github.com/yunussandikci/shareterm/server/service"
	"github.com/yunussandikci/shareterm/server/utils"
	"net/http"
)

type ApiHandler struct {
	fileService service.FileService
	encryptionService service2.EncryptionService
}

func NewApiHandler() *ApiHandler {
	return &ApiHandler{
		service.NewFileService(),
		service2.NewEncryptionService(),
	}
}

func (h ApiHandler) Create(c echo.Context) error {
	paste := new(common.PasteCreateRequest)
	bindErr := c.Bind(paste)
	if bindErr != nil {
		log.Error(bindErr)
		return echo.ErrBadRequest
	}
	generatedFileName, generateErr := h.fileService.GenerateFileName()
	if generateErr != nil {
		log.Error(generateErr)
		return echo.ErrInternalServerError
	}

	aesKey := common.GenerateAESKey()
	encryptedText, encryptErr := h.encryptionService.Encrypt([]byte(paste.Content), []byte(aesKey))
	if encryptErr != nil {
		log.Error(encryptErr)
		return echo.ErrInternalServerError
	}

	writeFileErr := h.fileService.WriteFile(generatedFileName, encryptedText)
	if writeFileErr != nil {
		log.Error(writeFileErr)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, common.PasteCreateResponse{
		URL: utils.BuildPasteWebReadURL(generatedFileName, aesKey),
	})
}