#include <Arduino.h>

static const char *stateA = "closed"; // text to send, if switch changes from closed to open
static const char *stateB = "open"; // text to send, if switch changes from open to closed
static const uint8_t INPUT_PIN = D5; // pin of the switch
static const int WS_PORT = 80; // port for the websocket server

#include <ESP8266WiFi.h>
#include <WebSocketsServer.h>
#include <secrets.h>

WiFiClient wifiClient;
WebSocketsServer webSocket = WebSocketsServer(WS_PORT);

bool state;

void setup_wifi() {
  WiFi.hostname("esp8266-door-monitor");

  Serial.println();
  Serial.print("Connecting to ");
  Serial.println(ssid);

  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password);

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }

  randomSeed(micros());

  Serial.println("");
  Serial.println("Wifi connected!");
}

void webSocketHandler(uint8_t num, WStype_t type, uint8_t * payload, size_t length) {
  switch(type) {
    case WStype_DISCONNECTED:
      Serial.printf("[%u] Disconnected!\n", num);
      break;
    case WStype_CONNECTED: 
      {
        IPAddress ip = webSocket.remoteIP(num);
        Serial.printf("[%u] Connected from %d.%d.%d.%d url: %s\n", num, ip[0], ip[1], ip[2], ip[3], payload);
        if (state) {
          webSocket.sendTXT(num, stateA);
        } else {
          webSocket.sendTXT(num, stateB);
        }
      }
      break;
    default:
      break;
  }
}

void setup() {
  Serial.begin(115200);

  setup_wifi();

  pinMode(INPUT_PIN, INPUT_PULLUP);
  state = digitalRead(INPUT_PIN) == HIGH;

  webSocket.begin();
  webSocket.onEvent(webSocketHandler);
}

void loop() {
  webSocket.loop();

  bool oldState = state;
  state = digitalRead(INPUT_PIN) == HIGH;

  if (state != oldState) {
    if (state) {
      webSocket.broadcastTXT(stateA);
    } else {
      webSocket.broadcastTXT(stateB);
    }
  }
}