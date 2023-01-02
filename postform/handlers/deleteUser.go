package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) DeleteUser(c *gin.Context) {
	var user Users
	if err := h.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Invalido ou já foi apagado. Aqui está uma lista de funcionarios atuais e seus ID'S"})
	} else {
		id := c.Params.ByName("id")
		d := h.DB.Where("id = ?", id).Delete(&user)
		fmt.Println(d)
		c.JSON(200, gin.H{"id #" + id: "deleted"})
	}
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
}
