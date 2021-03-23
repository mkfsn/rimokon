#ifndef MKFSN_UNBLOCK_TV_H
#define MKFSN_UNBLOCK_TV_H

#include <IRremote.h>
#include "device.h"

class UnblockTV: public Device {
   void pressPower();
   void pressUp();
   void pressDown();
   void pressLeft();
   void pressRight();
   void pressOK();
   void pressHome();
   void pressReturn();
   void pressMenu();
   void pressVolumeUp();
   void pressVolumeDown();
public:
   int executeCommand(const char* command);
};

void UnblockTV::pressPower() {
  IrSender.sendNEC(0xFF40, 0x4D, 1);
}

void UnblockTV::pressUp() {
  IrSender.sendNEC(0xFF40, 0xB, 1);
}

void UnblockTV::pressDown() {
  IrSender.sendNEC(0xFF40, 0xE, 1);
}

void UnblockTV::pressLeft() {
  IrSender.sendNEC(0xFF40, 0x10, 1);
}

void UnblockTV::pressRight() {
  IrSender.sendNEC(0xFF40, 0x11, 1);
}

void UnblockTV::pressOK() {
  IrSender.sendNEC(0xFF40, 0xD, 1);
}

void UnblockTV::pressHome() {
  IrSender.sendNEC(0xFF40, 0x1A, 1);
}

void UnblockTV::pressReturn() {
  IrSender.sendNEC(0xFF40, 0x42, 1);
}

void UnblockTV::pressMenu() {
  IrSender.sendNEC(0xFF40, 0x45, 1);
}

void UnblockTV::pressVolumeUp() {
  IrSender.sendNEC(0xFF40, 0x15, 1);
}

void UnblockTV::pressVolumeDown() {
  IrSender.sendNEC(0xFF40, 0x1C, 1);
}

int UnblockTV::executeCommand(const char* command) {
  if (strcmp(command, "pressPower") == 0) {
    this->pressPower();
  } else if (strcmp(command, "pressUp") == 0) {
    this->pressUp();
  } else if (strcmp(command, "pressDown") == 0) {
    this->pressDown();
  } else if (strcmp(command, "pressLeft") == 0) {
    this->pressLeft();
  } else if (strcmp(command, "pressRight") == 0) {
    this->pressRight();
  } else if (strcmp(command, "pressOK") == 0) {
    this->pressOK();
  } else if (strcmp(command, "pressHome") == 0) {
    this->pressHome();
  } else if (strcmp(command, "pressReturn") == 0) {
    this->pressReturn();
  } else if (strcmp(command, "pressMenu") == 0) {
    this->pressMenu();
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
