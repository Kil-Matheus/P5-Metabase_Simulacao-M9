package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
	"os"
	"net/http"
	"log"
	"strings"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Recebido: %s do tópico: %s com QoS: %d\n", msg.Payload(), msg.Topic(), msg.Qos())

	// Enviar a mensagem recebida para a API HTTP
	sendMessageToAPI(msg.Topic(), string(msg.Payload()), byte(msg.Qos()))
}

func sendMessageToAPI(topic, data string, qos byte) {
	// Montar a mensagem para enviar para a API
	message := fmt.Sprintf(`{"topic": "%s", "data": "%s", "qos": %d}`, topic, data, qos)

	// Fazer uma requisição HTTP POST para a API
	resp, err := http.Post("http://localhost:8080/messages", "application/json", strings.NewReader(message))
	if err != nil {
		log.Println("Erro ao enviar mensagem para a API:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Mensagem enviada para a API com sucesso.")
}

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
	opts.SetClientID("Subscriber")
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("#", 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	fmt.Println("Subscriber está rodando. Pressione CTRL+C para sair.")
	select {} // Bloqueia indefinidamente
	client.Disconnect(250)
}