package handler

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	service2 "github.com/yunussandikci/shareterm/common/service"
	"github.com/yunussandikci/shareterm/server/service"
	"net/http"
)

type WebHandler struct {
	fileService       service.FileService
	encryptionService service2.EncryptionService
}

func NewWebHandler() *WebHandler {
	return &WebHandler{
		service.NewFileService(),
		service2.NewEncryptionService(),
	}
}

func (h WebHandler) Read(c echo.Context) error {
	pasteFileName := c.Param("name")
	key := c.QueryParam("key")
	file, fileReadErr := h.fileService.ReadFile(pasteFileName)
	if fileReadErr != nil {
		log.Error(fileReadErr)
		return echo.ErrNotFound
	}
	decrypt, decryptErr := h.encryptionService.Decrypt(file, []byte(key))
	if decryptErr != nil {
		log.Error(decryptErr)
		return echo.ErrUnauthorized
	}
	return c.Render(http.StatusOK, "ReadView", string(decrypt))
}
