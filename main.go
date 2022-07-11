package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Global struct {
	Namespace string `json:"namespace,omitempty"`
	EnvName   string `json:"env_name,omitempty"`
}

// ToOtherInterfaceValue 通过json的方式将一个结构体转换成另一个结构体
func ToOtherInterfaceValue(toValue interface{}, fromValue interface{}) error {
	fromBytes, err := json.Marshal(fromValue)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(fromBytes, toValue); err != nil {
		return err
	}
	return nil
}

func main() {
	r := gin.Default()

	viper.SetConfigName("app")
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	global := Global{}
	ToOtherInterfaceValue(&global, viper.GetStringMap("global"))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, global)
	})

	r.Run()
}
