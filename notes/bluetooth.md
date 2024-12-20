# Bluetooth Pairing

## Configuring Bluetooth

Configuration for Bluetooth lives in `/etc/bluetooth/`. The configuration I'm using comes from this Reddit post:

<https://www.reddit.com/r/linuxquestions/comments/sqg220/how_do_i_make_my_device_appear_as_a_bluetooth/>

I was able to confirm the device class setting using this device type calculator:

<https://bluetooth-pentest.narod.ru/software/bluetooth_class_of_device-service_generator.html>

For more details, see [`./playbooks/audio/files/source.conf`](./playbooks/audio/files/source.conf).

**NOTE:** The main config is at `/etc/bluetooth/main.conf`, and it has a lot of handy comments. Worth referencing.

### Changing Device Name

The currently configured name is just the hostname (`%h`). It would be nice to fix this up to be `Stardeck` instead.

Hints: <https://askubuntu.com/questions/80960/how-to-change-bluetooth-broadcast-device-name>

## Pairing to a Device

**NOTE:** This also seems possible with KDE's bluetooth widget. But I'll have to try and reproduce it there.

Here are the steps I did that led to me successfully pairing my iPhone. I'm not sure what reference I followed - TODO.

I ran `bluetoothctl`, which put me in a REPL. I then ran:

```
discoverable on
pairable on
scan on  # See available devices
scan off  # Stop scanning
pair 44:90:BB:BC:FF:DF  # My iPhone
```

At this point, it should show up in my iPhone's bluetooth menu as a thing I can pair with. Once I click on it, things will go through a flow, and it will show up as connected.

From there, I ran:

```
connect 44:90:BB:BC:FF:DF
trust 44:90:BB:BC:FF:DF
```

I think all these steps are necessary, but I don't fully understand what they all mean. Figuring out what's fully necessary here is a TODO.

I *believe* the primary documentatio I was following here was from this post:

    <https://unix.stackexchange.com/questions/381342/bluetooth-a2dp-pulseaudio-source-to-play-sound-from-phone-to-linux-with-bluez-5>

## Forgetting a Device

For testing, forgetting a device is useful too. In `bluetoothctl`:

```
devices Paired  # Show what devices are paired
remove FD:13:1F:49:36:14  # This was my trackball, lol
```

## Configuring Pipewire

Pipewire with the default configuration Just Works. If you go into pipewire's lua scripts, you can see how it works - it listens for an audio source to pop up from bluez, and then automatically wires it to audio out. Easy-peasy - all you need to do is configure bluetooth correctly.

### Pulseaudio

On OS's other than Fedora, you *may* need to do this with pulseaudio. This post shows how to do it:

<https://fam-ribbers.com/blog/2019-11-17-share-a-sound-system-between-multiple-devices-using-a-raspberry-pi/>
