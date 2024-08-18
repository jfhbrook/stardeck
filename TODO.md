# TODO

## Mopidy and Mopidy-HTTP

Don't worry about a cool custom UI for now. Just use one of the canned ones.

## Bluetooth Pairing

Get a proof of concept for bluetooth pairing on deck. I want to be able to play
music off my laptop. I can cowpath this POC into scripts and/or a Cockpit app
later.

## COPR Packages

Prerequisite for file sharing...

- [ ] Get my current packages building again
  - Might need a Docker image for local dev on MacOS
- [ ] Create COPR project for `stardeck`

## File Sharing

Do it up with <https://github.com/45Drives/cockpit-file-sharing>.

- [ ] Create `cockpit-file-sharing` project on COPR
- [ ] Install `cockpit-file-sharing` from COPR
- [ ] Configure file sharing
- [ ] Add all my existing MP3s

## Audio Startup Hints

- [ ] Install `ocean-sound-theme`
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
