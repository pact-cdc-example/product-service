package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pact-cdc-example/product-service/pkg/cerr"
	"github.com/sirupsen/logrus"
)

type Handler interface {
	SetupRoutes(fr fiber.Router)
	GetProductByID(c *fiber.Ctx) error
	GetProductsByIDs(c *fiber.Ctx) error
	CreateProduct(c *fiber.Ctx) error
}

type handler struct {
	logger  *logrus.Logger
	service Service
}

type NewHandlerOpts struct {
	L *logrus.Logger
	S Service
}

func NewHandler(opts *NewHandlerOpts) Handler {
	return &handler{
		logger:  opts.L,
		service: opts.S,
	}
}

func (h *handler) CreateProduct(c *fiber.Ctx) error {
	h.logger.Infof("Create Product Request Arrived!")

	var req CreateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(cerr.BodyParser())
	}

	if err := req.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	product, err := h.service.CreateProduct(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.JSON(product)
}

func (h *handler) GetProductByID(c *fiber.Ctx) error {
	productID := c.Params("id")
	h.logger.Infof("Get Product By ID request arrived! Product ID: %s", productID)

	product, err := h.service.GetProductByID(c.Context(), productID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.JSON(product)
}

func (h *handler) GetProductsByIDs(c *fiber.Ctx) error {
	h.logger.Infof("Get Product By IDs request arrived!")

	var req GetProductsByIDsRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(cerr.BodyParser())
	}

	if err := req.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	products, err := h.service.GetProductsByIDs(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	err = c.JSON(products)
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return err
}

func (h *handler) SetupRoutes(fr fiber.Router) {
	productsGroup := fr.Group("/products")

	productsGroup.Post("/bulk", h.GetProductsByIDs)
	productsGroup.Get("/:id", h.GetProductByID)
	productsGroup.Post("/", h.CreateProduct)
}
