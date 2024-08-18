# TODO

## Ansible

Set up an Ansible playbook that runs locally. This will form the basis of
automatic updates and system config in the future.

## Developer Environment and Dotfiles

Manage dotfiles with ansible, not yadm. But otherwise, set up the same stuff
I have on lil-nas-x.

One thing I haven't done on lil-nas-x is setting up the SSH agent...
<https://unix.stackexchange.com/questions/132791/have-ssh-add-be-quiet-if-key-already-there>

## File Sharing

1. Do it up with <https://github.com/45Drives/cockpit-file-sharing>
2. Add all my existing MP3s

## Mopidy and Mopidy-HTTP

Don't worry about a cool custom UI for now. Just use one of the canned ones.

## Bluetooth Pairing

Get a proof of concept for bluetooth pairing on deck. I want to be able to play
music off my laptop. I can cowpath this POC into scripts and/or a Cockpit app
later.

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
