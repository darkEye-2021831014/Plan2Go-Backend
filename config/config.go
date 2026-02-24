package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)
type Config struct{
	Version string
	ServiceName string
	HttpPort int
	Jwt_SecretKey string
	OpenWeatherApiKey string
}
var configuration *Config
func LoadConfig() {
	err:=godotenv.Load()
	if err!=nil{
		fmt.Println("Error loading .env file",err) 
		os.Exit(1)
	}
	 version:=os.Getenv("VERSION")
	 if version==""{
		fmt.Println("VERSION not set in .env file")
		os.Exit(1)
	 }
	 serviceName:=os.Getenv("SERVICE_NAME")
	 if serviceName==""{
		fmt.Println("SERVICE_NAME not set in .env file")
		os.Exit(1)
	 }
	 httpPort:=os.Getenv("HTTP_PORT")
	 if httpPort==""{
		fmt.Println("HTTP_PORT not set in .env file")
		os.Exit(1)	
	 }
	 httpPortInt, err := strconv.Atoi(httpPort)
	 if err != nil {
		fmt.Println("HTTP_PORT must be a valid integer")
		os.Exit(1)
	 }
	 jwtSecretkey:=os.Getenv("JWT_SECRETKEY")
	 if jwtSecretkey==""{
		fmt.Println("JWT_SECRETKEY not set in .env file")
		os.Exit(1)	
	 }
	 openWeatherApiKey:=os.Getenv("OPENWEATHER_API_KEY")
	 if openWeatherApiKey==""{
		fmt.Println("OPENWEATHER_APIKEY not set in .env file")
		os.Exit(1)	
	 }
	 configuration=&Config{
		Version:version,
		ServiceName:serviceName,
		HttpPort:httpPortInt,
		Jwt_SecretKey: jwtSecretkey,
		OpenWeatherApiKey: openWeatherApiKey,
	}
}
func GetConfig() *Config{
	if(configuration==nil){
		LoadConfig()
	}
	return configuration
}