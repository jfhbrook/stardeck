# TODO

## MP3 Tagging

I have like 80gb of MP3s now, but the tagging and naming is a mess. I've started sorting through this on Windows. I should finish this up. Cool thing, Jupyter and pandas are involved.

- <https://methodmatters.github.io/editing-id3-tags-mp3-meta-data-in-python/>
- <https://mutagen.readthedocs.io/en/latest/>

## PlusDeck Driver

A basic dbus service exists! Just a long tail of small improvements.

## CrystalFontz Driver

The driver is mostly complete, but...

- [ ] Develop a dbus service
- [ ] Flesh out special character support

### stardeck-playbook

This is done. It's in node.

## stardeckd/stardeckctl

Way to get information on open windows: <https://github.com/jinliu/kdotool>

Wrap those up in a dbus service. The service can be used to trigger actions manually - for example, reloading the config.

The most interesting functionality, aside from Ansible, would be controlling the LCD and integrating tape deck events into that output.

This should also handle loopback in a clever way, using <https://pypi.org/project/pulsectl-asyncio/>.

## MP3 Tagging

My MP3s are kinda garbage in mopidy because they aren't tagged properly. This is probably something I could fix with some light scripting.

- <https://methodmatters.github.io/editing-id3-tags-mp3-meta-data-in-python/>
- <https://mutagen.readthedocs.io/en/latest/>

## Plusdeck UI

Two options: embed in Cockpit, or expose as a separate service. Leaning towards the latter, so I can skeumorph it up.

## HEOS Support

This would be hilarious, AND would make the HEOS app work with the Stardeck.

## Familiarize Myself with Audio CLI Tools

If I want to go headless, I'll need to be able to admin audio shenanigans over SSH. I found *some* tools - more to come - but I'll need to get into the habit of actually using them.

- `pamix`
- `pamixer`

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

## Sleep on Power Button

The machine ostensibly will sleep if I hit the power button. But KDE seems to interfere with that. Something to investigate more fully. In practice, I don't put the device to sleep and it runs really quiet.

<https://wiki.archlinux.org/title/Power_management#Power_management_with_systemd>
