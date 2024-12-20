#!/usr/bin/env python

#
# A half-done port of the tool at https://www.ampedrftech.com/cod.htm to
# Python and outputting hexademical.
#
# It turns out that there are a LOT of minor classes to enumerate, and my
# use for this kind of script is limited. But also, this tool works much
# better:
#
#     https://bluetooth-pentest.narod.ru/software/bluetooth_class_of_device-service_generator.html
#
# My use for this kind of thing is somewhat limited. But if I, for example,
# wanted a DSL for specifying the class in the future, this script might be
# an OK place to start.
#
# See also this document defining the classes:
#
#     https://www.ampedrftech.com/datasheets/cod_definition.pdf
#

service_class = 0
major_dev_class = 0
minor_dev_class = 0

minor_dev_class &= ~3

def service(bitshift):
    global service_class

    bitmask = 1 << bitshift
    service_class |= bitmask

def major_dev(value):
    global major_dev_class
    global minor_dev_class

    major_dev_class = value << 8
    minor_dev_class = 0


def minor_dev(bitmask, bitshift, value):
    global minor_dev_class
    minor_dev_class &= ~(bitmask << bitshift)
    minor_dev_class |= value << bitshift


# service classes
limited_discoverable_mode = 13
positioning = 16
networking = 17
rendering = 18
capturing = 19
object_transfer = 20
audio = 21
telephony = 22
information = 23

# major device classes
misc = 0
computer = 1
phone = 2
network = 3
audio = 4
peripheral = 5
imaging = 6
wearable = 7
toy = 8
health = 9
uncategorized = 0x1f

# minor device classes

# computer
c_uncategorized = (0x1f, 2, 0)
c_desktop = (0x1f, 2, 1)
c_server = (0x1f, 2, 2)
c_laptop = (0x1f, 2, 3)
c_handheld = (0x1f, 2, 4)
c_palm = (0x1f, 2, 5)
c_wearable = (0x1f, 2, 6)

# phone
p_uncategorized = (0x3f, 2, 0)
p_cell = (0x3f, 2, 1)
p_cordless = (0x3f, 2, 2)
p_smart = (0x3f, 2, 3)
modem = (0x3f, 2, 4)
isdn = (0x3f, 2, 5)

#lan
# from order of most to least available
lan = [(7, 5, i) for i in [0, 1, 2, 3, 4, 5, 6, 7]]
