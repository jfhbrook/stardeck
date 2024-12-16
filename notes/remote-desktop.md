# Notes on Remote Desktop

The things I want to do are pretty dependent on a graphical environment, especially given I don't have an entire software stack to replace it. Therefore, remote desktop would be *really* nice to have.

This whole thing is a fiasco because all the VNC and most of the RDP implementations were for X11. Wayland has some implementations, but they tend to be desktop-specific. In the case of KDE, this means `krfb`.

I actually have `krfb` installed and working, including autostart and unattended access. But there are two big problems that currently make it *really* limted:

1. Auto-login is not working. I configured SDDM to allow auto login and I
   confirmed the config file has those settings, but SDDM isn't respecting
   them. As far as I can tell, SDDM is running. Fixing this may require
   hitting up some forums.
2. `krfb` requests permission to screen share from Plasma every time it
   starts. I haven't even begun to investigate this. For all I know, there
   is no reasonable way to solve this.

So it's possible to do remote desktop, but you need to set it up from a non-remote desktop.

Also, in practice, remote desktop is laggy. Some of that is probably wifi being janky because I'm an antenna short. But even so... it's not fun. It works, but just walking up to it is easier.
