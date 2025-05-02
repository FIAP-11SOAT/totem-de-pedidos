package handlers

import "github.com/labstack/echo/v4"

type DocsHandler struct{}

func NewDocsHandler() *DocsHandler {
	return &DocsHandler{}
}

func (h *DocsHandler) Docs(c echo.Context) error {
	return c.String(200, "OpenAPI documentation goes here")
}
