# Setup

## WiFi

* `nmcli` installed by default
* `nmtui` now installed
* Cockpit's ability to configure wifi very limited.
* A UI on the Crystalfontz is necessary. 

## Cockpit

Just finding my way around for now. Some resources for later:

* <https://github.com/45Drives/cockpit-file-sharing>
* Make an avahi app
* Old but promising: <https://github.com/cyberorg/apsetup-cockpit>

## Avahi

<https://fedoramagazine.org/find-systems-easily-lan-mdns/>

This is more or less plug and play. Relevant settings are mostly going to be
`hostname` (configurable under `overview`)  and the short list of
things in `/etc/avahi/avahi-daemon.conf`.

Note, mDNS needs to be allowed by the firewall before it will work fully.

## Pulseaudio

Installed. Two things to try:

* <https://github.com/patroclos/PAmix> - TUI mixer
* <https://github.com/cdemoulins/pamixer> - CLI mixer

Will eventually need to feed "line in" into output.

## ProjectM

This could be THE way to get visualizations out of PulseAudio:

<https://github.com/projectM-visualizer/projectm>

Unfortunately, it's not SUPER straightforward to capture/stream OpenGL output.
The easiest thing might be to run a headless Wayland and stream to that. But
really I just need to research this more... eventually.

## git, github
