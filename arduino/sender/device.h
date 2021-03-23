#ifndef MKFSN_DEVICE_H
#define MKFSN_DEVICE_H

class Device {
public:
  virtual int executeCommand(const char* command) = 0;
};

#endif
