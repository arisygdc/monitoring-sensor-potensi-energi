/**
   PostHTTPClient.ino

    Created on: 21.11.2016

*/

#include <ArduinoJson.h>
#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>

#define SERVER_IP "192.168.1.234"

#ifndef STASSID
#define STASSID "ssid"
#define STAPSK  "password"
#endif

int id_sensor = 0;

void setup() {

  Serial.begin(9600);

  Serial.println();
  Serial.println();
  Serial.println();

  WiFi.begin(STASSID, STAPSK);

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("");
  Serial.print("Connected! IP address: ");
  Serial.println(WiFi.localIP());
  delay(2000);
  
  
  if ((WiFi.status() == WL_CONNECTED)) {

    WiFiClient client;
    HTTPClient http;

    Serial.print("[HTTP] begin...\n");
    // configure traged server and url
    http.begin(client, "http://" SERVER_IP "/api/v1/setup"); //HTTP
    http.addHeader("Content-Type", "application/json");

    Serial.print("[HTTP] POST http://" SERVER_IP "/api/v1/setup ...\n");
    // start connection and send HTTP header and body
    int httpCode = http.POST("{\"sensor\": {\"tipe_sensor\":\"angin\",\"identity\":\"cx7\"},\"lokasi\":{\"nama\":\"wit gedhang\",\"provinsi\":\"jawa timur\",\"kecamatan\":\"candi\",\"desa\":\"sidoarjo\"}}");
    Serial.printf("before %d\n", id_sensor);
    
    // httpCode will be negative on error
    if (httpCode > 0) {
      // HTTP header has been send and Server response header has been handled
      Serial.printf("[HTTP] POST... code: %d\n", httpCode);
      // file found at server;
      if (httpCode == 202) {
        const String& payload = http.getString();
        DynamicJsonDocument doc(1024);
        deserializeJson(doc, payload);
        Serial.println("received payload:\n<<");
        Serial.println(payload);
        Serial.println(">>");
        id_sensor = doc["id_sensor"];
      }
    } else {
      Serial.printf("[HTTP] POST... failed, error: %s\n", http.errorToString(httpCode).c_str());
    }
    Serial.printf("id sensor : %d", id_sensor);
    http.end();
  }
}

void loop() {
//   wait for WiFi connection
  if ((WiFi.status() == WL_CONNECTED) && (id_sensor != 0) {

    WiFiClient client;
    HTTPClient http;

    Serial.print("[HTTP] begin...\n");
    // configure traged server and url
    http.begin(client, "http://" SERVER_IP "/api/v1/sensor/data"); //HTTP
    http.addHeader("Content-Type", "application/json");

    Serial.print("[HTTP] POST...\n");
    // start connection and send HTTP header and body
    int httpCode = http.POST("{\"id_sensor\": 4,\"data\": 76}");

    // httpCode will be negative on error
    if (httpCode > 0) {
      // HTTP header has been send and Server response header has been handled
      Serial.printf("[HTTP] POST... code: %d\n", httpCode);

    } else {
      Serial.printf("[HTTP] POST... failed, error: %s\n", http.errorToString(httpCode).c_str());
    }

    http.end();
  }

  delay(10000);
}