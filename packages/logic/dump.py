import subprocess
import service
import re

BatteryDumpType = {
  "present": bool,
  "ac powered": bool,
  "usb powered": bool,
  "wireless powered": bool,
  "max charging current": int,
  "max charging voltage": int,
  "charge counter": int,
  "temperature": int,
  "voltage": int,
  "status": int,
  "health": int,
  "level": int,
  "scale": int,
  "technology": str
}

class Dump:
  def __init__(self, value):
    self.value = value

  def getValue(self, key):
    match = re.search(f"{key}:\s*(\w|\d)+", self.value, re.I)

    return match.group().split(':')[1].strip() if match else None


def dumpService(name):
  result = subprocess.call(["dumpsys", name])

  return Dump(result)


def initObject(dump, interface):
  result = {}

  for key, ttype in interface.items():
    value = dump.getValue(key)
    result[key] = None if value is None else ttype(value)

  return result


def getBatteryDump():
  dump = dumpService(service.Battery)

  return initObject(dump, BatteryDumpType)
