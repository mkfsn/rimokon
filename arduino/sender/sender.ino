#include <IRremote.h> // 3.0.3
#include <ArduinoJson.h> // 6.17.3

#include "device.h"

#include "SharpTV.h"
#include "UnblockTV.h"
#include "LogiSpeaker.h"

#define IR_SEND_PIN 5

SharpTV sharpTV;
UnblockTV unblockTV;
LogiSpeaker logiSpeaker;

void setup() {
  Serial1.begin(57600);
}

void loop() {
  StaticJsonDocument<256> request;

  if (Serial1.available() > 0) {
    String serialData = Serial1.readStringUntil('\n');

    DeserializationError error = deserializeJson(request, serialData);
    if (error) {
      Serial1.print("{\"error\":\"");
      Serial1.print(error.f_str());
      Serial1.println("\"}");
      return;
    }

    const char* device = request["device"];
    if (!device) {
      Serial1.println("{\"error\":\"no device\"}");
      return;
    }

    Device *d;
    if (strcmp(device, "SharpTV") == 0) {
      d = &sharpTV;
    } else if (strcmp(device, "LogiSpeaker") == 0) {
      d = &logiSpeaker;
    } else if (strcmp(device, "UnblockTV") == 0) {
      d = &unblockTV;
    }

    if (!d) {
      Serial1.println("{\"error\":\"unknown device\"}");
      return;
    }

    const char* command = request["command"];
    if (!command){
      Serial1.println("{\"error\":\"no command\"}");
      return;
    }

    if (d->executeCommand(command) == 0) {
      Serial1.println("{\"ok\":true}");
    } else {
      Serial1.println("{\"error\":\"unknown command\"}");
    }
  }
}
