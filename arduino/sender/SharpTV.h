#ifndef MKFSN_SHARP_TV_H
#define MKFSN_SHARP_TV_H

#include <IRremote.h>
#include "device.h"

class SharpTV: public Device {
  void pressPower();
  void pressInput();
  void pressLeft();
  void pressRight();
  void pressEnter();
  void pressVolumeUp();
  void pressVolumeDown();
public:
  int executeCommand(const char* command);
};

void SharpTV::pressPower() {
  IrSender.sendSharp(0x10, 0x68, 1);
}

void SharpTV::pressInput() {
  IrSender.sendSharp(0x10, 0xC8, 1);
}

void SharpTV::pressLeft() {
  IrSender.sendSharp(0x10, 0xEB, 1);
}

void SharpTV::pressRight() {
  IrSender.sendSharp(0x10, 0x1B, 1);
}

void SharpTV::pressEnter() {
  IrSender.sendSharp(0x10, 0x4A, 1);
}

void SharpTV::pressVolumeUp() {
  IrSender.sendSharp(0x10, 0x28, 1);
}

void SharpTV::pressVolumeDown() {
  IrSender.sendSharp(0x10, 0xA8, 1);
}

int SharpTV::executeCommand(const char* command) {
  if (strcmp(command, "pressPower") == 0) {
    this->pressPower();
  } else if (strcmp(command, "pressInput") == 0) {
    this->pressInput();
  } else if (strcmp(command, "pressLeft") == 0) {
    this->pressLeft();
  } else if (strcmp(command, "pressRight") == 0) {
    this->pressRight();
  } else if (strcmp(command, "pressEnter") == 0) {
    this->pressEnter();
  } else if (strcmp(command, "pressVolumeUp") == 0) {
    this->pressVolumeUp();
  } else if (strcmp(command, "pressVolumeDown") == 0) {
    this->pressVolumeDown();
  } else {
    return -1;
  }
  return 0;
}

#endif
