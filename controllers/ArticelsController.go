package controllers

import (
	"fmt"

	"github.com/PrasGi/learn-golang/initializers"
	"github.com/PrasGi/learn-golang/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var articels []models.Articel

	// Ambil semua data dari tabel Articel
	if err := initializers.DB.Find(&articels).Error; err != nil {
		// Cetak informasi kesalahan pada server logs
		fmt.Println("Error during find:", err)

		// Kirim respon JSON dengan pesan kesalahan
		c.JSON(500, gin.H{"error": err})
		return
	}

	// Kirim respon JSON dengan data Articel
	c.JSON(200, gin.H{
		"message": "success",
		"data":    articels,
	})
}

func Store(c *gin.Context) {
	var body struct {
		Content string `json:"content"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	articel := models.Articel{Content: body.Content}

	result := initializers.DB.Create(&articel)

	if result.Error != nil {
		// Cetak informasi kesalahan pada server logs
		fmt.Println("Error during create:", result.Error)

		// Kirim respon JSON dengan pesan kesalahan
		c.JSON(500, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, gin.H{
		"message": "Berhasil",
		"data":    articel,
	})
}

func Show(c *gin.Context) {
	var articel models.Articel

	if err := initializers.DB.First(&articel, c.Param("id")).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Berhasil",
		"data":    articel,
	})
}

func Destroy(c *gin.Context) {
	var articel models.Articel

	if err := initializers.DB.First(&articel, c.Param("id")).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Delete(&articel).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Berhasil",
	})
}

func Update(c *gin.Context) {
	var articel models.Articel

	if err := initializers.DB.First(&articel, c.Param("id")).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var body struct {
		Content string `json:"content"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	articel.Content = body.Content

	if err := initializers.DB.Save(&articel).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Berhasil",
		"data":    articel,
	})
}
