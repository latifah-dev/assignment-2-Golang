package controllers

import (
	"Tugas2/database"
	"Tugas2/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetOrder(c *gin.Context) {
	var DetailOrders []models.Order
	db := database.GetDB()
	if err := db.Preload("Items").Order("ID desc").Find(&DetailOrders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if len(DetailOrders) > 0 {
		c.JSON(http.StatusOK, DetailOrders)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": "Data Kosong",
		})
	}

}
func GetOrderById(c *gin.Context) {
	var DetailOrders models.Order
	db := database.GetDB()
	if err := db.Preload("Items").First(&DetailOrders, c.Param("orderID")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, DetailOrders)
}

func CreateOrder(c *gin.Context) {
	var input models.Orderitems

	if err := c.ShouldBindJSON(&input); err != nil {
		errMess := map[string]string{}
		for _, v := range err.(validator.ValidationErrors) {
			errMess[v.StructField()] = "Must Be " + v.ActualTag()
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errMess,
		})
		return
	}
	waktu := time.Now()
	db := database.GetDB()
	tx := db.Begin()
	data := models.Order{
		Customer_Name: input.Customer_Name,
		Ordered_at:    waktu,
	}
	if err := tx.Create(&data).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// db.Create(&data)
	var ItemCreate []models.Items
	for i, v := range input.Items {
		n := input.Items[i].Quantity
		gatau, _ := strconv.ParseUint(string(n), 10, 64)
		ItemCreates := []models.Items{
			{
				Item_Code:   v.Item_Code.String(),
				Description: v.Description,
				Quantity:    uint(gatau),
				OrderID:     data.ID,
			},
		}
		ItemCreate = append(ItemCreate, ItemCreates...)
	}

	if err := tx.Create(&ItemCreate).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, input)
}

func UdateOrder(c *gin.Context) {
	var (
		input        models.Orderitems
		DetailOrders models.Order
		DetailItems  []models.Items
		err          error
	)
	id := c.Param("orderID")
	db := database.GetDB()
	if err = db.Preload("Items").First(&DetailOrders, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		errMess := map[string]string{}
		for _, v := range err.(validator.ValidationErrors) {
			errMess[v.Field()] = "Must Be " + v.ActualTag()
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errMess,
		})
		return
	}

	DetailOrders.Customer_Name = input.Customer_Name
	DetailOrders.Ordered_at = time.Now()
	db.Where("order_id = ?", id).Find(&DetailItems)
	for i, v := range input.Items {
		n := input.Items[i].Quantity
		gatau, _ := strconv.ParseUint(string(n), 10, 64)
		DetailItems[i].Description = v.Description
		DetailItems[i].Item_Code = v.Item_Code.String()
		DetailItems[i].Quantity = uint(gatau)

	}
	tx := db.Begin()
	if err := tx.Debug().Save(&DetailItems).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	DetailOrders.Items = DetailItems
	if err := tx.Debug().Updates(&DetailOrders).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, DetailOrders)
}

func DeleteOrder(c *gin.Context) {
	var (
		DetailOrders models.Order
		DetailItems  []models.Items
	)
	id := c.Param("orderID")
	db := database.GetDB()
	if err := db.Where("id = ?", id).First(&DetailOrders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}
	tx := db.Begin()
	if err := tx.Debug().Where("order_id = ?", id).Delete(&DetailItems).Error; err != nil {
		tx.Rollback()
		return
	}
	if err := tx.Debug().Delete(&DetailOrders, id).Error; err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{
		"message": "data berhasil di Hapus",
	})
}
