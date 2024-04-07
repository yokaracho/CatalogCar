package handler

import (
	"CatalogCar/pkg/logger"
	"CatalogCar/pkg/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var getLogger = logger.GetLogger()

// InsertCars godoc
// @Summary Добавить автомобили
// @Description Добавляет новые автомобили в базу данных.
// @Accept json
// @Produce json
// @Param cars body object true "Список регистрационных номеров автомобилей"
// @Param cars.reg_num array string true "Список регистрационных номеров автомобилей"
// @Success 201 {object} map[string]interface{} "Данные об автомобилях вставлены"
// @Failure 400 {object} map[string]string "Неверные данные в запросе"
// @Failure 500 {object} map[string]string "Ошибка при вставке данных об автомобилях"
// @Router /api/data/create [POST]
func (h *Handler) InsertCars(c echo.Context) error {
	req := struct {
		RegNums []string `json:"reg_num"`
	}{}

	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		getLogger.Infof("Неверные данные в запросе: %s", err.Error())
		getLogger.Debugf("Неверные данные в запросе: %s", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверные данные в запросе"})
	}

	rowsAffected, err := h.service.InsertCars(c.Request().Context(), req.RegNums)
	if err != nil {
		getLogger.Infof("Ошибка при вставке данных об автомобилях: %s", err.Error())
		getLogger.Debugf("Ошибка при вставке данных об автомобилях: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при вставке данных об автомобилях"})
	}

	getLogger.Infof("Данные об автомобилях вставлены, количество строк: %d", rowsAffected)
	getLogger.Debugf("Данные об автомобилях вставлены, количество строк: %d", rowsAffected)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("Данные об автомобилях вставлены, количество строк: %d", rowsAffected),
	})
}

// GetCars godoc
// @Summary Получить автомобили
// @Description Получает список автомобилей из базы данных.
// @Param reg_num query string false "Регистрационный номер"
// @Param mark query string false "Марка"
// @Param model query string false "Модель"
// @Param year query int false "Год выпуска"
// @Param offset query int false "Смещение"
// @Param limit query int false "Ограничение"
// @Produce json
// @Response 200 OK {array} []Car
// @Failure 500 {Invalid year parameter}
// @Router /api/data [GET]
func (h *Handler) GetCars(c echo.Context) error {
	var filters model.CarFilter
	filters.RegNum = c.QueryParam("reg_num")
	filters.Mark = c.QueryParam("mark")
	filters.Model = c.QueryParam("model")

	yearParam := c.QueryParam("year")
	if yearParam != "" {
		year, err := strconv.Atoi(yearParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid year parameter"})
		}
		filters.Year = year
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}
	filters.Offset = strconv.Itoa(offset)

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 10
	}
	filters.Limit = strconv.Itoa(limit)

	cars, err := h.service.GetCars(context.Background(), filters)
	if err != nil {
		getLogger.Infof("Ошибка: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error: " + err.Error()})
	}

	return c.JSON(http.StatusOK, cars)
}

// UpdateInfo godoc
// @Summary Обновить информацию об автомобиле
// @Description Обновляет информацию об указанном автомобиле в базе данных.
// @Accept json
// @Produce json
// @Param id path int true "Идентификатор автомобиля для обновления"
// @Success 200 {object} map[string]string "Информация обновлена"
// @Failure 400 {object} map[string]string "Неверный идентификатор или неверные данные в запросе"
// @Failure 500 {object} map[string]string "Ошибка при обновлении информации"
// @Router /api/cars/{id} [PUT]
func (h *Handler) UpdateInfo(c echo.Context) error {
	ctx := context.Background()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		getLogger.Infof("Неверный идентификатор: %s", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверный идентификатор"})
	}

	var updateData model.CarModel
	if err := c.Bind(&updateData); err != nil {
		getLogger.Debugf("Ошибка при извлечении данных из запроса: %s", err.Error())
		getLogger.Infof("Ошибка при извлечении данных из запроса: %s", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверные данные в запросе"})
	}
	updateData.CarID = id

	getLogger.Debugf("Обновляем информацию для ID: %d, новые данные: %+v", updateData.CarID, updateData)

	if err := h.service.UpdateInfo(ctx, &updateData); err != nil {
		getLogger.Debugf("Ошибка при обновлении информации: %s", err.Error())
		getLogger.Infof("Ошибка при обновлении информации: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при обновлении информации"})
	}

	getLogger.Debugf("Информация успешно обновлена для ID: %d", updateData.CarID)
	getLogger.Infof("Информация успешно обновлена для ID: %d", updateData.CarID)
	return c.JSON(http.StatusOK, map[string]string{"message": "Информация обновлена"})
}

func (h *Handler) UpdateOwner(c echo.Context) error {
	ctx := context.Background()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		getLogger.Infof("Неверный идентификатор: %s", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверный идентификатор"})
	}

	var updateData model.PeopleModel
	if err := c.Bind(&updateData); err != nil {
		getLogger.Debugf("Ошибка при извлечении данных из запроса: %s", err.Error())
		getLogger.Infof("Ошибка при извлечении данных из запроса: %s", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверные данные в запросе"})
	}
	updateData.OwnerID = id

	getLogger.Debugf("Обновляем информацию для ID: %d, новые данные: %+v", updateData.OwnerID, updateData)

	if err := h.service.UpdateOwner(ctx, &updateData); err != nil {
		getLogger.Debugf("Ошибка при обновлении информации: %s", err.Error())
		getLogger.Infof("Ошибка при обновлении информации: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при обновлении информации"})
	}

	getLogger.Debugf("Информация успешно обновлена для ID: %d", updateData.OwnerID)
	getLogger.Infof("Информация о владельце успешно обновлена для ID: %d", updateData.OwnerID)
	return c.JSON(http.StatusOK, map[string]string{"message": "Информация о владельце обновлена"})
}

// DeleteCarByID godoc
// @Summary Удалить автомобиль по идентификатору
// @Description Удаляет информацию об автомобиле из базы данных по указанному идентификатору.
// @Accept json
// @Produce json
// @Param id path int true "Идентификатор удаляемого автомобиля"
// @Success 200 {object} map[string]string "Запись удалена"
// @Failure 400 {object} map[string]string "Неверный идентификатор"
// @Failure 404 {object} map[string]string "Запись не найдена"
// @Failure 500 {object} map[string]string "Ошибка при удалении записи"
// @Router /api/cars/{id} [DELETE]
func (h *Handler) DeleteCarByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		getLogger.Debugf("Не совпадает ID:%s", err.Error())
		getLogger.Infof("Не совпадает ID:%s", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	rowsAffected, err := h.service.DeleteCarByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if rowsAffected == 0 {
		getLogger.Infof("Запись не найдена")
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Record not found"})
	}

	getLogger.Debugf("Запись удалена")
	getLogger.Infof("Запись удалена")
	return c.JSON(http.StatusOK, map[string]string{"message": "Record deleted"})
}
