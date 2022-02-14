0.19.0
---
- **new devices**
    - ft6336: add support for ft6336
    - pca9685: PCA9685 driver
    - shtc3: Sensirion SHTC3 Relative Humidity / Temperature i2c sensor
    - sx126x: Driver for Semtech sx126x radio modules
    - xpt2046: XPT2046 Touch driver (#350)
- **enhancements**
    - **hd44780i2c** 
        - clean up for go fmt
        - Needed fixes and update hd44780i2c.go
    - **ili9341, ili9342**
        - add support for m5stack
        - add support for m5stack-core2
    - **wifi**
        - modify to use shared net.Adapter interface for all supported wifi devices
    - wifinina: remove busy wait
- **bugfixes**
    - **hd44780** 
        - fix 4-bit data length flag
        - Reset data pins to output mode after reading
    - Nano 33 BLE drivers (#351)
- **docs**
    - examples/wifi: add unified example for tcpclient that compiles for all supported wifi adaptors

0.18.0
---
- **new devices**
    - apds9960: add support for APDS-9960 Digital Proximity sensor
    - axp192: add support for AXP192 single Cell Li-Battery and power system management IC
    - hts221: add support for HTS221 capacitive digital sensor for relative humidity and temperature
    - i2csoft: add support for software I2C
    - image: add support for image/jpeg and image/png
    - lps22hb: add support for LPS22HB MEMS nano pressure sensor
    - lsm6dox: add support for lsm6dox accelerometer
    - lsm9ds1: add support for lsm9ds1 accelerometer
- **enhancements**
    - ili9341: change to use drivers.SPI interface
    - **ws2812**
        - generate assembly instead of handwriting it
        - improve timings to be compatible with the WS2811
        - add support for 168MHz (e.g. Adafruit Feather STM32F405)
        - add support for RISC-V
    - wifinina: control nina pins, for example leds
- **docs**
    - rtl8720dn: examples for tcpclient, udpstation, mqtt, and webserver
    - **wifinina** 
        - nina-fw update docs
        - examples/wifinina/http-get
    - ili9341: refactor examples
    - Fix broken link for SHT3x datasheet
- **core**
    - all: use build directives for both Go1.17 and earlier versions
- **bugfixes**
    - net: fix raddr of tcp conn
    - mcp3008: fix bitshift bug

0.17.1
---
- To correct an error in the release process. Same as 0.17.0.

0.17.0
---
- **new devices**
    - rtl8720dn: add support for rtl8720dn
    - sdcard: add support for spi sdcard driver, along with fatfs
- **enhancements**
    - apa102: use 4-byte buffer to improve speed
    - bmi160: avoid heap allocations
    - ili9341: add standard SPI driver
    - wifinina
        - avoid fmt package
        - Fix RSSI command for WiFiNINA + Print current SSID + Wait for correct time before printing it out + Cleanup
    - ws2812
        - rename the pin to ws2812
        - add tag for nrf52833
        - Disable interrupts before sending ws2812 data
        - add support for qtpy and atsame5x
- **core**
    - modules: switch to use tinygo-org version of tinyfs package
    - all: use machine.Serial as the default output

0.16.0
---
- **new devices**
    - aht20: add device
    - ina260: add new i2c device
    - keypad: add 4x4 keypad driver (#226)
    - max7219: add driver support
    - mcp2515: add support for mcp2515 CAN device
    - p1am: support the P1AM-100 hardware watchdog
    - pcf8563: add support for pcf8563 real time clock
    - servo: add driver using PWM
    - tm1637: add support for tm1637 7-segment LED
    - tone: add package for producing tones using the PWM interface
- **enhancements**
    - pwm: update drivers with PWM to use new interface
    - wifinina: Make TLS work over WiFiNINA Verified on Arduino Nano33 IoT and nina fw v1.4.5
    - ssd1306: Enable reset screen for SSD1306 via I2C
    - st7789: add scrolling functions to match st7735
- **bugfixes**
    - wifinina:
        - fix getMACAddress and getTime
        - fix println + cleanup
        - remove debug flag and remove unnecessary padding call
        - fix padding and implement missing functions
    - flash: fix EraseBlocks method which is erasing sectors instead
- **core**
    - all: use interfaces for UART objects
    - all: do not take the pointer of an I2C object
    - adc: update drivers with ADC to use new config struct
- **testing**
    - tester:
        - add a mock for command-oriented i2c devices
        - add 16-bit register mock device

- **docs**
    - ssd1306: example of ssd1306 with 128x64 display over I2C
    - wifinina:
        - add information about Adafruit boards with ESP32 wifi coprocessors, and modify examples to remove code that was both not being used, and also prevented many Adafruit boards from being able to be targeted by the examples
        - update docs to simplify the nina-fw update process
        - example that connects to AP and prints ip addresses, time and mac
    - p1am: documentation and example program
    - add missing new drivers added since last release

0.15.0
---
- **new devices**
    - dht: add DHTXX thermometer
    - mcp23017: new driver for MCP23017 (I2C port expander)
    - bmp388: Add bmp388 support (#219)
- **enhancements**
    - hd44780: add a mode to work with boards where the RW pin is grounded
    - st7789: add scrolling functions to match st7735
    - microbitmatrix: matrix now working on microbit v2
    - ds1307: Better interface "ReadTime" instead of "Time"
    - ws2812: make AVR support more robust
- **bugfixes**
    - all: fix main package in examples
- **core**
    - adc: update all drivers with ADC to use new config struct
    - spi: remove machine.SPI and replace with drivers.SPI interface for almost all SPI drivers
- **testing**
    - test: run unit tests against i2c drivers and any spi drivers without direct gpio
- **docs**
    - st7789: correct errors on various godoc comments

0.14.0
---
- **new devices**
    - lis2mdl: add LIS2MDL magnetometer (#187)
    - waveshare: add Waveshare 4.2in B/W e-paper driver (#183)
- **enhancements**
    - adt7410: add connection test and for that matter connection method
    - gps
        - add speed and heading to fix, as parsed from RMC NMEA sentence
        - improvements and bugfixes (#186)
    - ili9341
        - add support for setting framerate, vsync pause, and reading scanline data.
        - renamed NewSpi() to NewSPI() in accordance with Go naming conventions
    - ws2812
        - add support for ESP8266
        - add support for ESP32
- **bugfixes**
    - ili9341
        - rix setWindow bug, add CS pin for Clue compatibility. (#180)
        - bugfix for RAMWR bug
    - lis2mdl: turn on read mode on every read, to ensure that magnetometer data is updated
- **core**
    - i2c
        - switch all i2c drivers definitions to use i2c bus interface type instead of machine package concrete type
        - correct interface definition for I2C Tx function
- **testing**
    - fix smoke-test unless avr-gcc installed
    - add very basic mock structs for testing i2c devices, based on work done by @rogpeppe
    - improve API surface and implement one more test function in lis2mdl driver
- **docs**
    - replace README badge for godocs with pkgdocs

0.13.0
---
- **new devices**
    - bmi160: add initial support
    - bmp280: added support for the Bosch BMP280 temperature and pressure sensor. (#158)
    - lsm303agr: add lsm303agr (#162)
    - ssd1351: add SSD1351 OLED display driver (#146)
- **enhancements**
    - hd44780: add Hd44780i2c driver (#173)
    - ili9341
        - add ILI9341 TFT driver (SPI) for ATSAMD2x (#174)
        - cache address window to prevent sending unnecessary commands (#171)
        - ILI9341 TFT driver (SPI) (#153)
        - improve performance of ILI9341 on ATSAMD5X
    - ST77xx: fix DrawFastHLine for ST77xx, SSD1331 and SSD1351 DrawFastHLine uses FillRectangle(x,y,width,height,c), so height must be 1 to draw a horizontal line
    - tmp102: add Connected func to check for device
    - wifinina: added UDP support
    - ws2812: update ws2812_avr_16m.go
- **bugfixes**
    - apa102: avoid creating garbage
    - bmp180: fix temperature type conversion
- **core**
    - all
        - added custom import path (#161)
        - changeover to eliminate all direct use of master/slave terminology
    - build: try vendor in working directory to match expected module path
    - ci: support Go modules
    - modules: update go version and dependency
- **docs**
    - docs: reorder to correct alpha and adjust count of supported drivers

0.12.0
---
- **new devices**
    - hcsr04: Added HC-SR04 ultrasonic distance sensor. (#143)
    - spi/qspi: Low-level IO driver for serial flash memory via SPI and QSPI (#124)
    - tmp102: TMP102 low-power digital temperature sensor (#141)
    - amg88xx: AMG88xx thermal camera module
- **bugfixes**
    - mqtt: reduce use of goroutines in router to not start a new goroutine for each invocation of each callback

0.11.0
---
- **new devices**
    - shiftregister: Support for various shift register chips (#135)
- **enhancements**
    - shifter: simplify API surface for PyBadge (#137)
    - shifter: new API for shifter driver
    - mqtt: use buffered channels for incoming messages to handle bursts
    - ili9341: Adding scroll functionality (#121)
- **bugfixes**
    - wifinina: fix typo on StartScanNetworks
    - ili9341: various bugfixes for display
- **examples**
    - semihosting: add example
- **docs**
    - readme: Use degree sign instead of ordinal
    - all: fix celsius symbol in all code comments

0.10.0
---
- **new devices**
    - adt7410: Support for ADT7410 temperature sensor (#109)
    - ili9341: ILI9341 TFT driver (#115)
    - l293x: added support for h-bridge motor controller
    - l9110x: add support for L9110x h-bridge motor driver
    - resistive: Adding driver for four-wire resistive touchscreen (#118)
- **enhancements**
    - st7735: added scroll functionality to st7735
    - st7735: remove default offsets
    - st7789: remove default offsets
    - ws2812: Added nrf52840 tag to ws2812
    - ws2812: work-arounds to allow Digispark to control WS2812 LEDs
- **docs**
    - readme: update README to include list of all 44 drivers
    - wifinina: update docs and add Dockerfile to build firmware
    - wifinina: update docs and info on how to install WiFiNINA driver

0.9.0
---
- **new devices**
    - net: shared implementation of net package for serial wifi devices
    - shifter: add support for bit Parallel In Serial Out (PISO) shifter
    - stepper: add support for dual stepper motor
    - wifinina: add implementation for WiFiNINA firmware
- **enhancements**
    - st7735: improvements in st7735 driver
    - st7789: improvements in st7789 driver
    - ws2812: add support for 120Mhz Cortex-M4
    - ws2812: added Feather M0 and Trinket M0 to build tags for WS2812
    - ws2812: add support for simulation
- **bugfixes**
    - ws2812: fix "invalid symbol redefinition" error
- **examples**
    - Add examples for wifinina drivers

0.8.0
---
- **new devices**
    - mcp3008: add implementation for MCP3008 ADC with SPI interface
    - semihosting: initial implementation of ARM semihosting
- **enhancements**
    - espat: refactor response processing for greater speed and efficiency
    - espat: implement mqtt subscribe functionality via blocking select/channels (experiemental)
- **bugfixes**
    - st7789: fix index out of bounds error
- **examples**
    - Add espat driver example for mqtt subscribe

0.7.0
---
- **new devices**
    - veml6070: add Vishay UV light sensor
- **enhancements**
    - lis3dh: example uses I2C1 so requires config to specify pins since they are not default
    - ssd1331: make SPI TX faster
    - st7735: make SPI Tx faster
- **docs**
    - complete missing GoDocs for main and sub-packages
- **core**
    - add Version string for support purposes
- **examples**
    - Change all espat driver examples to use Arduino Nano33 IoT by default

0.6.0
---
- **new devices**
    - Support software SPI for APA102 (Itsy Bitsy M0 on-board "Dotstar" LED as example)

0.5.0
---
- **new devices**
    - LSM6DS3 accelerometer
- **bugfixes**
    - ws2812: fix timings for the nrf51
- **enhancements**
    - ws2812: Add build tag for Arduino Nano33 IoT

0.4.0
---
- **new devices**
    - SSD1331 TFT color display
    - ST7735 TFT color display
    - ST7789 TFT color display
- **docs**
    - espat
        - complete list of dependencies for flashing NINA-W102 as used in Arduino Nano33 IoT board.

0.3.0
---
- **new devices**
    - Buzzer for piezo or small speaker
    - PDM MEMS microphone support using I2S interface
- **enhancements**
    - epd2in13: added rotation
    - espat
        - add built-in support for MQTT publish using the Paho library packets, alongside some modifications needed for the AT protocol.
        - add DialTLS and Dial methods, update MQTT example to allow both MQTT and MQTTS connections
        - add example that uses MQTT publish to open server
        - add README with information on how to flash ESP32 or ESP8266 with AT command set firmware.
        - add ResolveUDPAddr and ResolveTCPAddr implementations using AT command for DNS lookup
        - change Response() method to use a passed-in timeout value instead of fixed pauses.
        - implement TCPConn using AT command set
        - improve error handling for key TCP functions
        - refactor net and tls interface compatible code into separate sub-packages
        - update MQTT example for greater stability
        - use only AT commands that work on both ESP8266 and ESP32
        - add documentation on how to use Arduino Nano33 IoT built-in WiFi NINA-W102 chip.
- **bugfixes**
    - core: Error strings should not be capitalized (unless beginning with proper nouns or acronyms) or end with punctuation, since they are usually printed following other context.
    - docs: add note to current/future contributors to please start by opening a GH issue to avoid duplication of efforts
    - examples: typo in package name of examples
    - mpu6050: properly scale the outputs of the accel/gyro

0.2.0
---
- **new devices**
    - AT24C32/64 2-wire serial EEPROM
    - BME280 humidity/pressure sensor
- **bugfixes**
    - ws2812: better support for nrf52832

0.1.0
---
- **first release**
    - This is the first official release of the TinyGo drivers repo, matching TinyGo 0.6.0. The following devices are supported:
        - ADXL345
        - APA102
        - BH1750
        - BlinkM
        - BMP180
        - DS1307
        - DS3231
        - Easystepper
        - ESP8266/ESP32
        - GPS
        - HUB75
        - LIS3DH
        - MAG3110
        - microbit LED matrix
        - MMA8653
        - MPU6050
        - PCD8544
        - SHT3x
        - SSD1306
        - Thermistor
        - VL53L1X
        - Waveshare 2.13"
        - Waveshare 2.13" (B & C)
        - WS2812
