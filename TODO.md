# TODO

## File Sharing

Assuming Mopidy is a thing, I should configure samba and get some file sharing going. I can also use this to upload my current mp3 collection. Heck yeah.

I think KDE is already configured for this - just gotta set it up and transfer my files over.

## Mopidy

I'm not actually wild about Mopidy. But it would definitely be nice to have a way to play music headlessly. It's also a relatively easy win. I'd like to spend some time to play with it and get a basic setup going. Minimum is playing off disk, but ideally I can also do YouTube Music.

- [x] Install mopidy base
- [x] Install mopidy-iris
- [ ] Fix/install mopidy-ytmusic
- [ ] Configure mopidy

mopidy-ytmusic is currently broken for... reasons.

## Ansible/Terraform It Up

It's becoming rapidly clear that a lot of my setup should be handled with ansible. I don't like using ansible directly, so doing up Terraform will be valuable.

- [ ] Install ansible and terraform
- [ ] Get stub module working/running
- [ ] Port easy/straightforward stuff to ansible/terraform

### Port from Scripts

- [ ] 1password
  - at least, the package installs
  - separate role for op-cli
- [ ] asdf - options may be limited here
- [ ] audio
  - Check out attic for missing things
- [ ] brass-tacks/cool-kids-cli-tools
  - Publish existing modules to ansible-galaxy
- [ ] cockpit
  - Is there an existing role?
- [ ] gh-cli
  - Is there an existing role?
- [ ] git
  - Use existing role
- [ ] joshiverse
  - COPR
- [ ] joshiverse packages
- [ ] kitty
- [ ] mopidy
  - Including configurence
- [ ] nmtui
- [ ] volta (node-dev)
- [ ] uv (python-dev)
- [ ] rpmfusion
  - Is there an existing role?
- [ ] rustup (rust-dev)
- [ ] starship
- [ ] vim
  - Use existing role
- [ ] neovim
- [ ] watchexec
- [ ] yadm
  - Fork [existing role](https://docs.debops.org/en/stable-2.1/ansible/roles/yadm/)

### Port from Attic

- [ ] ansible.cfg?
- [ ] bluetooth
- [ ] development - a few tools
- [ ] ffmpeg
- [ ] tools - anything missing?

### Open Question

Ansible and terraform need to be installed. But there's a bootstrapping issue. Perhaps these should continue to be scripted?

## SSH Key

1password ssh management is great on desktop, but annoying when ssh'd in. I should just fix this.

## Upgrade to F41

Note that setup will be a little different, as `dnf` had a major upgrade.

## Bluetooth Pairing

I'd love to be able to do bluetooth pairing. But from what I've discovered, it's a MESS.

See [`./notes/bluetooth.md`](./notes/bluetooth.md) for more details.

Bluetooth pairing is a MESS.

First, a lot of the scripts/tools say to use `sbc`, which has more or less
been sunset. Some documentation on that lives here:

<https://github.com/ev3dev/ev3dev/issues/198>

But also, Fedora uses PipeWire by default. It _appears_ that PipeWire handles
this differently - but perhaps better?

<https://www.reddit.com/r/pipewire/comments/s3jth9/has_anyone_ever_been_able_to_play_bluetooth_audio/>

I suspect that this stuff is configurable with... whatever UI-driven tools
pipewire has.

## PlusDeck Support

1. Develop a dbus interface for Plusdeck
2. Develop a Cockpit application for Plusdeck

## CrystalFontz Support

1. Write a driver/client library
2. Write higher level interface/notification functions
3. Write a dbus service
4. ???
5. Profit

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

# Cockpit Loose Ends

If I get this far, I'll want to ensure that anything I'd reasonably want to do (that doesn't otherwise have an interface) can be accessed through Cockpit.
