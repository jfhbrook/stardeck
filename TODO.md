# TODO

## SSH Key

1password ssh management is great on desktop, but annoying when ssh'd in. I should just fix this.

## Bluetooth Pairing

I'd love to be able to do bluetooth pairing. But from what I've discovered, it's a MESS.

See [`./notes/bluetooth.md`](./notes/bluetooth.md) for more details.

- [x] Initial pairing
- [ ] Write down steps to pair
- [-] Configure bluetooth correctly
- [ ] Research bluetooth source in pipewire
- [ ] ???

## Sleep on Power Button

Can I get the laptop to sleep when I hit the power button?

## HEOS Support for Mopidy

This would be hilarious, AND would make the HEOS app work with the Stardeck.

## Upgrade to F41

Note that setup will be a little different, as `dnf` had a major upgrade.

## PlusDeck Support

1. Develop a dbus interface for Plusdeck
2. Develop a Cockpit application for Plusdeck

## MP3 Tagging

My MP3s are kinda garbage in mopidy because they aren't tagged properly. This is probably something I could fix with some light scripting.

## CrystalFontz Support

1. Write a driver/client library
2. Write higher level interface/notification functions
3. Write a dbus service
4. ???
5. Profit

## More COPR packages

- [ ] yamlfmt

## Separate Development Ansible

I want the essential tasks to be separate from the "I regularly log onto this" tasks.

## Familiarize Myself with Audio CLI Tools

If I want to go headless, I'll need to be able to admin audio shenanigans over SSH. I found *some* tools - more to come - but I'll need to get into the habit of actually using them.

- `pamix`
- `pamixer`

## Auto Login

This currently isn't working at all, though it should be possible. This will be critical if I want to go headless and can't find a way to ditch dependency on a desktop environment.

## Audio Hints

It will be pretty important to have audio indications that things are up and running, if the Stardeck is going to be headless. I *think* this is a relatively easy win? But low priority until I'm serious about going headless.

I already have files downloaded, which is good.

There are basically two angles here. One is customizing KDE's standard sounds and/or enabling any sounds it doesn't have going by default. The other is hooking audio hints into whatever stack I have going to do bluetooth pairing - a whole nother thing.

- [x] Install `ocean-sound-theme`
- [ ] POC playing theme sounds with ffmpeg
- [ ] Connect sounds to events in systemd
- [ ] Investigate sounds for bluetooth events
- [ ] Add custom sounds

## Fix Mopidy-YTMusic

mopidy-ytmusic is broken. But I think that's because it hasn't been updated in a while. Some patches and it would probably work fine.

# Cockpit Loose Ends

If I get this far, I'll want to ensure that anything I'd reasonably want to do (that doesn't otherwise have an interface) can be accessed through Cockpit.

## Bootstrapping

Once I have this buttoned up, I'll want a way to bootstrap a fresh install of Fedora. This implies a bootstrap script, and potentially some functionality pulled out of Ansible. But we can cross that bridge later.
