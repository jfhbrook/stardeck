# TODO

## packages

- nerd fonts dnf packages imo
- starship dnf package
  - installed to /usr/local/bin
- watchexec

## Remote Desktop

The things I want to do are pretty dependent on a graphical environment,
especially given I don't have an entire software stack to replace it.
Therefore, _some_ form of remote desktop is more or less crucial.

I finally got `krfb` to work, including autostart and unattended access. But
there are two big problems that currently have me dead in the water:

1. Auto-login is not working. I configured SDDM to allow auto login and I
   confirmed the config file has those settings, but SDDM isn't respecting
   them. As far as I can tell, SDDM is running. Fixing this may require
   hitting up some forums.
2. `krfb` requests permission to screen share from Plasma every time it
   starts. I haven't even begun to investigate this. For all I know, there
   is no reasonable way to solve this.

Note that solutions aside from `krfb`, particularly TightVNC but also
NoMachine, struggle with KDE for Wayland reasons. The way forward here is
almost certainly `krfb`.

### A Small Monitor

If I can't get the Stardeck to work with remote desktop consistently, I'll
need a petite monitor I can hook it up to. My current extra monitor is no
good because the 10:16 aspect ratio makes the Stardeck so confused it won't
consistently display on it.

## Bluetooth Pairing

Bluetooth pairing is a MESS.

First, a lot of the scripts/tools say to use `sbc`, which has more or less
been sunset. Some documentation on that lives here:

<https://github.com/ev3dev/ev3dev/issues/198>

But also, Fedora uses PipeWire by default. It _appears_ that PipeWire handles
this differently - but perhaps better?

<https://www.reddit.com/r/pipewire/comments/s3jth9/has_anyone_ever_been_able_to_play_bluetooth_audio/>

I suspect that this stuff is configurable with... whatever UI-driven tools
pipewire has.

## Mopidy

- [x] Install mopidy base
- [ ] Configure mopidy as a service
- [ ] Install [an extension](https://mopidy.com/ext/)
- [ ] Configure favored frontend as a service

## File Sharing

## Audio Startup Hints

- [x] Install `ocean-sound-theme`
- [ ] POC playing theme sounds with ffmpeg
- [ ] Connect sounds to events in systemd
- [ ] Investigate sounds for bluetooth events
- [ ] Add custom sounds

## PlusDeck Support

1. Develop a dbus interface for Plusdeck
2. Develop a Cockpit application for Plusdeck

## CrystalFontz Support

1. Write a driver/client library
2. Write higher level interface/notification functions
3. Write a dbus service
4. ???
5. Profit
