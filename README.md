# P5-Metabase_Simulacao-M9

# P5-Metabase_Simulacao-M9

## Documentação da Aplicação MQTT e API

Esta documentação descreve a estrutura e funcionalidade de uma aplicação MQTT (Message Queuing Telemetry Transport) que inclui um Publisher, um Subscriber e uma API HTTP. O Subscriber consome a API para enviar os dados recebidos pelos tópicos MQTT para a API HTTP que envia para um banco de dados que é utilizado para alimentar um Dashboard no MetaBase.

### Publisher

O Publisher é responsável por simular a leitura de sensores e publicar os dados gerados em tópicos MQTT. Ele estabelece uma conexão com um broker MQTT e periodicamente publica mensagens contendo dados simulados de sensores em um tópico específico.

### Subscriber

O Subscriber é responsável por se conectar a um broker MQTT, subscrever a todos os tópicos disponíveis e lidar com as mensagens recebidas. Quando recebe uma mensagem em qualquer tópico, ele invoca um handler que envia os dados recebidos para a API HTTP. O Subscriber é configurado para enviar as mensagens recebidas pela MQTT para a API HTTP, que posteriormente armazena esses dados no banco de dados SQLite.

### Broker

A tecnologia utilizada é o HiveMQ, uma plataforma que permite subir em cloud um broker gratuito, na qual usando as credeciais do servidor, é possível fazer uma conexão que permite se inscrever e publicar.

### API

A API HTTP é responsável por receber mensagens HTTP contendo dados de sensores. Ela decodifica as mensagens recebidas, armazena-as no banco de dados SQLite e envia uma resposta HTTP adequada para o remetente. A API HTTP está configurada para lidar com solicitações POST no endpoint "/messages".

### Fluxo de Dados

1. O Publisher gera dados simulados de sensores.
2. Os dados são publicados em tópicos MQTT pelo Publisher.
3. O Subscriber recebe as mensagens MQTT em todos os tópicos e as encaminha para a API HTTP.
4. A API HTTP recebe os dados dos sensores, os decodifica, armazena no banco de dados SQLite e responde com um status HTTP apropriado.

Este fluxo de dados permite a comunicação bidirecional entre os sistemas de sensoriamento e os servidores de dados por meio do protocolo MQTT e da API HTTP.

## Link do Broker

Link: https://console.hivemq.cloud/

## Vídeo de Execução

Link: https://drive.google.com/file/d/1dGFxFk6ON_zURzbdB0N92RD2f_9Brils/view?usp=sharing
