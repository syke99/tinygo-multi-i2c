# tinygo-multi-i2c
A reimplementation of the TinyGo drivers package for communicating with multiples of the same (supported) devices on one individual I2C bus.

Supported Devices
====
- ADXL345 - I2C Accelerometor by iMEMS
- AMG88XX series - I2C High Precision Infrared Array Sensor  by Panasonic
- BH1750 - I2C Light Sensor by SMAKN
- BLINKM: I2C-controllable LEDs 
- BME280 - I2C Humidity and Pressure Sensor by Bosch Sensortec
- BMP280 - I2C Barometric Pressure Sensor by Bosch Sensortec
- DS3231 - I2C RTC (Real-Time-Clock) by Adafruit
- LIS3DH - I2C Acclerometer by STMicroelectronics
- MPU6050 - I2C (Gyro + Acclerometer) MEMS Motion Tracking Device by InvenSense/TDK

Pre-requisites
====
- you must have Go installed (recommended v1.17+)
- you must have TinyGo installed (quick install guides found [here](https://tinygo.org/getting-started/install/)) and configured for the IDE you are using (guides for both VSCode, IntelliJ IDEA, and other IDEs found [here](https://tinygo.org/docs/guides/ide-integration/))

Why was tinygo-multi-i2c built?
====
tinygo-multi-2c was built for the purpose of simplifying the process of setting up multiples of one or more devices that communicate over the I2C protocol using pre-written drivers from TinyGo (found [here](https://tinygo.org/docs/concepts/drivers/)).

What problem does tinygo-multi-i2c solve?
=====
While using the pre-written drivers from TinyGo is extremely useful, it takes a handful of lines to set up one device. And that's only if you're using the default address that specifc device uses. Whereas using the tinygo-multi-i2c package, you can create a device, configure it, set the address specifically to the one you want, and test the connection for that device all in as little as two lines after initializing the your I2C bus, allowing for cleaner, more concise code. Initializing the I2C bus; which will be the same as doing so with the drivers package from TinyGo (an example of can be found on TinyGo's page for Divers [here](https://github.com/tinygo-org/drivers); there's also an example down below). The error returned, if nil isn't returned, will help point you in the direction of where your device set-up is failing to allow for smooth debugging.

How do I use tinygo-multi-i2c?
=====
After installing the pre-requisites and configuring your IDE, simply import tinygo-multi-i2c using go get

```bash
$ go get github.com/syke99/tinygo-multi-i2c
```

Then you can import the package in any go file you'd like

```go
import multi (

    "machine"

    "github.com/syke99/go-c2dmc"
)
```

### Basic usage

Initialize your I2C bus by initializing and configuring an machine.I2C0 to be passed into the NewDevice() method. Ex:

```go
    i2c := machine.I2C0
    err := i2c.Configure(machine.I2CConfig{
        SCL: machine.P0_30,
        SDA: machine.P0_31,
    })
    if err != nil {
        println("could not configure I2C:", err)
        return
    }
```

Declare a variable to hold the returned Devices struct that will be used to access the device you create, and then call multi.NewDevices(params) and pass in your parameters, the first being the name of the device you would like to create in all lowercase. Ex:

```go
devices, error := multi.NewDevice(i2c, "bmp280", 0, 0, 0, 0, 0)
```

**NOTE:** if you are setting up a BMP280 device, you must provide 5 arguments (of type uint) to the NewDevice method to configure the BMP280's Standby, Filter, Temperature, Pressure, and Mode. If not, simply pass in all 0's. This example just uses all 0's with a bmp280 for brevity's sake.

After that, you can use the variable you declared, followed by the name of the device you just created using dot notation, followed by the function you would like to use. (A list of available devices can be found above. Just change any letters to their respective lowercase.) Ex:

```go
pressure, error := devices.bmp.ReadPressure()
```

This process can be repeated by simply repeating the line to create a new device, just with a new name of a variable to hold the Devices struct for each device you wish to create. Ex:

```go
// Device 1, a BMP280 sensor
d1, error := multi.NewDevice(i2c, "bmp280", 0, 0, 0, 0, 0)

// Device 2, an MPU6050 motion tracking device
d2, error := multi.NewDevice(i2c, "mpu6050", 0, 0, 0, 0, 0)

// Device 3, a BH1750 light sensor
d3, error := multi.NewDevice(i2c, "bh1750", 0, 0,0,0,0)

pressure, error := d1.bmp280.ReadPressure()
// handle error and/or using pressure reading value here

acceleration, error := d2.mpu6050.ReadAcceleration()
// handle error and/or using acceleration reading value here

illuminance, error := d3.bh1750.Illuminance()
// handle error and/or using illuminance reading value here
```

(screenshot will go here)
