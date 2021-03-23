#ifndef MKFSN_LOGI_SPEAKER_H
#define MKFSN_LOGI_SPEAKER_H

#include <IRremote.h>
#include "device.h"

class LogiSpeaker: public Device {
  void pressPower();
public:
  int executeCommand(const char* command);
};

void LogiSpeaker::pressPower() {
  IrSender.sendNEC(0x7484, 0xFF, 1);
}

int LogiSpeaker::executeCommand(const char* command) {
  if (strcmp(command, "pressPower") == 0) {
    this->pressPower();
  } else {
    return -1;
  }
  return 0;
}

#endif
