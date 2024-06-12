package controller

import (
	"net/http"
	"strconv"

	"github.com/haerul-umam/capstone-project-mikti/model/web"
	"github.com/haerul-umam/capstone-project-mikti/service"
	"github.com/labstack/echo/v4"
)

type CategoryContollerImpl struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryContollerImpl {
	return &CategoryContollerImpl{
		categoryService: categoryService,
	}
}

func (contoller *CategoryContollerImpl) NewCategory(e echo.Context) error {
	category := new(web.CategoryRequest)

	if err := e.Bind(category); err != nil {
		return e.JSON(http.StatusBadRequest, web.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := e.Validate(category); err != nil {
		return err
	}

	saveCategory, errSaveCategory := contoller.categoryService.CreateCategory(category.Name)

	if errSaveCategory != nil {
		return e.JSON(http.StatusBadRequest, web.ResponseToClient(http.StatusBadRequest, errSaveCategory.Error(), nil))
	}

	return e.JSON(http.StatusOK, web.ResponseToClient(http.StatusOK, "Berhasil Membuat Category", saveCategory))
}

func (controller *CategoryContollerImpl) GetCategoryList(c echo.Context) error {
	getCategorys, errGetCategorys := controller.categoryService.GetCategoryList()

	if errGetCategorys != nil {
		return c.JSON(http.StatusNotFound, web.ResponseToClient(http.StatusNotFound, errGetCategorys.Error(), nil))
	}
	return c.JSON(http.StatusOK, web.ResponseToClient(http.StatusOK, "Succes", getCategorys))
}

func (controller *CategoryContollerImpl) UpdateCategory(c echo.Context) error {
	category := new(web.CategoryUpdateServiceRequest)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, web.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	categoryUpdate, errCategoryUpdate := controller.categoryService.UpdateCategory(*category, id)

	if errCategoryUpdate != nil {
		return c.JSON(http.StatusBadRequest, web.ResponseToClient(http.StatusBadRequest, errCategoryUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, web.ResponseToClient(http.StatusOK, "sukses ubah kategori", categoryUpdate))
}

func (controller *CategoryContollerImpl) DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := controller.categoryService.DeleteCategory(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, web.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, web.ResponseToClient(http.StatusOK, "Kategori Berhasil Dihapus", nil))
}
