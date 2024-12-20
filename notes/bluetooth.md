# Bluetooth Pairing

## Configuring Bluetooth

Configuration for Bluetooth lives in `/etc/bluetooth/`. The configuration I'm using comes from this Reddit post:

<https://www.reddit.com/r/linuxquestions/comments/sqg220/how_do_i_make_my_device_appear_as_a_bluetooth/>

I was able to confirm the device class setting using this device type calculator:

<https://bluetooth-pentest.narod.ru/software/bluetooth_class_of_device-service_generator.html>

For more details, see [`./playbooks/audio/files/source.conf`](./playbooks/audio/files/source.conf).

### Changing Device Name

The currently configured name is just the hostname (`%h`). It would be nice to fix this up to be `Stardeck` instead.

Hints: <https://askubuntu.com/questions/80960/how-to-change-bluetooth-broadcast-device-name>

## Pairing to a Device

TK

Link dump:

- <https://unix.stackexchange.com/questions/381342/bluetooth-a2dp-pulseaudio-source-to-play-sound-from-phone-to-linux-with-bluez-5>

## Configuring Pulseaudio and/or Pipewire

This blog post shows how to configure pulseaudio:

<https://fam-ribbers.com/blog/2019-11-17-share-a-sound-system-between-multiple-devices-using-a-raspberry-pi/>

This will probably work. Fedora uses pipewire, not pulseaudio, but I also have `pipewire-pulseaudio` installed. I'm actually already using that to expose a TCP sound server, so I can probably continue to follow this approach. Basically, I would want to go to [`./playbooks/audio/files/pipewire-pulse.conf`](./playbooks/audio/files/pipewire-pulse.conf) and add that module call to the `pulse.cmd` block. Though, I'd probably have to install those modules, too.

However, it *might* be possible to make this happen natively with pipewire. The documentation online is probably not as good and they *might* delegate this to pipewire-pulse. But if I can get it to show up as a source, I should be able to prototype the patch with Helvium, and productionize it with wireplumber (`/etc/wireplumber`).

Link dump:

- <https://gist.github.com/zw/3349078>
