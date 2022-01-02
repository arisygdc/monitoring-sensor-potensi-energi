#include <ArduinoJson.h>
#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>

#define SERVER_IP "192.168.1.123"
#define METHOD_POST "POST"

#ifndef STASSID
#define STASSID "ssid"
#define STAPSK  "password"
#endif

int id_sensor = 0;

struct HTTP_RESULT {
  int code;
  String payload;
};

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

  String setupJson = "{\"sensor\": {\"tipe_sensor\":\"angin\",\"identity\":\"cx7\"},\"lokasi\":{\"nama\":\"wit gedhang\",\"provinsi\":\"jawa timur\",\"kecamatan\":\"candi\",\"desa\":\"sidoarjo\"}}";
  if ((WiFi.status() == WL_CONNECTED)) {
    HTTP_RESULT getResponse = HTTPsend(METHOD_POST, "http://" SERVER_IP "/api/v1/setup", setupJson);
    Serial.printf("[HTTP] POST SETUP DEVICE... code: %d\n", getResponse.code);
    if (getResponse.code == 202) {
      const String& payload = getResponse.payload;
      DynamicJsonDocument doc(1024);
      deserializeJson(doc, payload);
      Serial.println("received payload:\n<<");
      Serial.println(payload);
      Serial.println(">>");
      id_sensor = doc["id_sensor"];
    }
  }
  if (id_sensor == 0) {
    while (1) {
      Serial.println("Device Not Setup");
      delay(10000);
    }
  }
}

void loop() {
  //   wait for WiFi connection
  HTTP_RESULT getResponse;
  if ((WiFi.status() == WL_CONNECTED)) {
    String valueSensorJson = "{\"id_sensor\": 4,\"data\": 76}";
    getResponse = HTTPsend(METHOD_POST, "http://" SERVER_IP "/api/v1/sensor/data", valueSensorJson);
    Serial.printf("[HTTP] POST SETUP DEVICE... code: %d\n", getResponse.code);
  }

  delay(10000);
}

HTTP_RESULT HTTPsend(String http_method, String url, String jsonBody) {
  WiFiClient client;
  HTTPClient http;

  // configure traged server and url
  http.begin(client, url); //HTTP
  http.addHeader("Content-Type", "application/json");
  int httpCode;
  if (http_method == METHOD_POST) {
    // start connection and send HTTP header and body
    httpCode = http.POST(jsonBody);
  }
  // httpCode will be negative on error
  HTTP_RESULT getResp = {httpCode, http.getString()};
  http.end();
  return getResp;
}
