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
- you must have Go installed (recommended v1.17+, for Windows 10, you MUST use Go v1.16 or greater)
- you must have TinyGo installed (quick install guides found [here](https://tinygo.org/getting-started/install/)) and configured for the IDE you are using (guides for both VSCode, IntelliJ IDEA, and other IDEs found [here](https://tinygo.org/docs/guides/ide-integration/))

Why was tinygo-multi-i2c built?
====
tinygo-multi-2c was built for the purpose of simplifying the process of setting up multiples of one or more devices that communicate over the same I2C bus using the I2c protocol and pre-written drivers from TinyGo (found [here](https://tinygo.org/docs/concepts/drivers/)).

What problem does tinygo-multi-i2c solve?
=====
While using the pre-written drivers from TinyGo is extremely useful, it takes a handful of lines to set up one device. And that's only if you're using the default address that specifc device uses. Whereas using the tinygo-multi-i2c package, you can create a device, configure it, set the address specifically to the one you want, and test the connection for that device all in as little as two lines after initializing the your I2C bus, allowing for cleaner, more concise code. Initializing the I2C bus; which will be the same as doing so with the drivers package from TinyGo (an example of can be found on TinyGo's page for Divers [here](https://tinygo.org/docs/concepts/drivers/); there's also an [example](https://github.com/syke99/tinygo-multi-i2c/blob/main/README.md#basic-usage) down below). The error returned, if nil isn't returned, will help point you in the direction of where your device set-up is failing to allow for smooth debugging.

How do I use tinygo-multi-i2c?
=====
After installing the pre-requisites and configuring your IDE, simply import tinygo-multi-i2c using go get

```bash
$ go get github.com/syke99/tinygo-multi-i2c
```

Then you can import the package in any go file you'd like

```go
import (
    "machine"

     multi "github.com/syke99/go-c2dmc"
)
```

### Basic usage

Initialize your I2C bus by initializing and configuring an machine.I2C0 to be passed into the NewDevice() method. Ex:

```go
    i2c := machine.I2C0
    err := i2c.Configure(machine.I2CConfig{
        SCL: machine.P0_30, // <-- These values will be dependent on what microcontroller you're using
        SDA: machine.P0_31, // <-- These values will be dependent on what microcontroller you're using
    })
    if err != nil {
        println("could not configure I2C:", err)
        return
    }
```

Declare a variable to hold the returned Devices struct that will be used to access the device you create, and then call multi.NewDevices(params) and pass in your parameters, the first being the name of the device you would like to create in all lowercase. Ex:

```go
// since this isn't a BMP280, we need to just pass in a slice of uint's that are all 0
// it can also be reused for any other I2C device that you're using that isn't a BMP280
// a BMP280 needs 5 uint's passed to it for its Standby, Filter, Temperature, Pressure, and Mode.
b := [5]uint{0,0,0,0,0}

// the default address for a BLINKM is 0x09, so to dynamically set it, we pass in the address, 
// i.e.: 0x00 in this example
// if you want to use the default address of 0x09, you can just pass in a 0 as the address
devices, error := multi.NewDevice(i2c, "blinkm", 0x00, b)
```

**NOTE:** if you are setting up a BMP280 device, you must provide 5 arguments (of type uint) to the NewDevice method to configure the BMP280's Standby, Filter, Temperature, Pressure, and Mode. If not, simply pass in all 0's. This example just uses all 0's with a BMP280 for brevity's sake.

After that, grab the device you created and save it in a variable by using dot notation with the Devices struct you got returned. Ex:

```go
b := devices.blinkm
```

Then you can simply just call the function that corresponds to how you want to interact with the device. Ex:

```go
error := b.FadeToRGB()
```

This process can be repeated by simply repeating the line to create a new device, just with a new name of a variable to hold the Devices struct for each device you wish to create. Ex:

```go
// These devices are using the default addresses, so passing in 0, and using the []uint, b we initialized above since none of them are a BMP280

// Device 1, a BLINKM RGB light
d1, error := multi.NewDevice(i2c, "blinkm", 0, b)
// Handle or ignore error (advised to not ignore)
bl := d1.blinkm

// Device 2, an MPU6050 motion tracking device
d2, error := multi.NewDevice(i2c, "mpu6050", 0, b)
// Handle or ignore error (advised to not ignore)
m := d2.mpu6050

// Device 3, a BH1750 light sensor
d3, error := multi.NewDevice(i2c, "bh1750", 0, b)
// Handle or ignore error (advised to not ignore)
bh := d3.bh1750

error := bl.FadeToRGB()
// handle error and/or using pressure reading value here

acceleration, error := m.ReadAcceleration()
// handle error and/or using acceleration reading value here

illuminance, error := bh.Illuminance()
// handle error and/or using illuminance reading value here
```

(screenshot will go here)
