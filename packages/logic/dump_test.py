from unittest.mock import patch
from dump import getBatteryDump
import pprint


def Test_getBatteryDump():
    with patch("dump.subprocess") as mock_subprocess:
        mock_subprocess.call.return_value = """Battery stat: Current Battery Service state:
  AC powered: true
  USB powered: false
  Wireless powered: false
  Max charging current: 0
  Max charging voltage: 0
  Charge counter: 0
  status: 2
  health: 2
  present: true
  level: 68
  scale: 100
  voltage: 0
  temperature: 0
  technology: Li-ion
"""
        pprint.pprint(getBatteryDump())


def main():
    Test_getBatteryDump()


if __name__ == "__main__":
    main()
