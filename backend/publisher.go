package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
	"os"
	"time"
	"math/rand"
)

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("Publisher")
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		//zl1, zl2 := simularLeituraSensor()
		zl1 := simularLeituraSensor()
		//text := fmt.Sprintf("%.2f,%.2f", zl1, zl2)
		text := fmt.Sprintf("%.2f", zl1)
		token := client.Publish("Zona_Leste/topic", 1, false, text)
		token.Wait()
		fmt.Println("Publicado:", text)
		time.Sleep(2 * time.Second)
	}
	client.Disconnect(250)
}

func simularLeituraSensor() (float64){
	zl1 := rand.Float64() * 100
	//zl2 := rand.Float64() * 100
	return zl1
	//return zl1, zl2
}