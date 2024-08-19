# TODO

## Get Audio Working

I installed PulseAudio, but it's completely unconfigured.

Some research suggests PipeWire is the default audio server, rather than
PulseAudio - so I should switch to that. But it *also* looks like you're
supposed to install ALSA under the hood.

I'm testing that audio works by running `just play sounds/win3x/TADA.WAV`.

## Bluetooth Pairing

Bluetooth pairing is a MESS.

First, a lot of the scripts/tools say to use `sbc`, which has more or less
been sunset. Some documentation on that lives here:

<https://github.com/ev3dev/ev3dev/issues/198>

But also, Fedora uses PipeWire by default. It *appears* that PipeWire handles this differently - but perhaps better?

<https://www.reddit.com/r/pipewire/comments/s3jth9/has_anyone_ever_been_able_to_play_bluetooth_audio/>

I suspect that this stuff is configurable with... whatever UI-driven tools pipewire has.

## Mopidy

- [X] Install mopidy base
- [ ] Configure mopidy as a service
- [ ] Install [an extension](https://mopidy.com/ext/)
- [ ] Configure favored frontend as a service

## File Sharing

Do it up with <https://github.com/45Drives/cockpit-file-sharing>.

- [ ] Create `cockpit-file-sharing` project on COPR
- [ ] Install `cockpit-file-sharing` from COPR
- [ ] Configure file sharing
- [ ] Add all my existing MP3s

## Audio Startup Hints

- [X] Install `ocean-sound-theme`
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
