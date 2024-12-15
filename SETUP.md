# Setup

## The Basics

- [x] 1password
  - settings -> developer
    - integrate 1password CLI
    - set up ssh client
- [x] git
  - [x] install with dnf
  - [x] git config --global user.name "Josh Holbrook"
  - [x] git config --global user.email "josh.holbrook@gmail.com"
  - [x] git config --global init.defaultBranch main
- [x] clone stardeck repo
- [x] rpmfusion free/nonfree repos
- [ ] yadm
  - install yadm
  - create new stardeck-dotfiles repo
  - initialize yadm
  - clone old dotfiles repo
  - add dotfiles I currently have going
- [ ] sshd

- [ ] edge
- [ ] edge apps
  - [ ] youtube
  - [ ] youtube music

- [ ] kitty
  - curl -L https://sw.kovidgoyal.net/kitty/installer.sh | sh /dev/stdin
  - probably snag slowpoke config
  - manually create desktop file
- [ ] nerd fonts
  - dnf repo?
  - if manual, just mononoki for now
- [ ] starship
- [ ] neovim
  - set up dotfile based on slowpoke
- [ ] vanilla vim
  - base config off slowpoke
- [ ] CLI tools
  - bat
  - bats
  - eza (exa is unmaintained)
  - fd-find
  - fzf
  - hexedit
  - hexyl
  - htop
  - jq
  - neofetch
  - pandoc
  - ripgrep
  - shellcheck
  - ag
  - just
  - watchexec
- [ ] set hostname
- [ ] coc.nvim extensions
  - typescript
  - python
  - go
- [ ] asdf
  - [ ] follow getting started guide
  - [ ] asdf-ruby
  - [ ] asdf-java
- [ ] uv (python)
- [ ] volta
- [ ] spruce up bashrc
- [ ] set up update scripts

- [ ] podman
  - [ ] install podman
  - [ ] configure socket
  - [ ] podman desktop
- [ ] cool ass wallpaper
- [ ] cool ass lock screen
- [ ] configure ipv6 on wifi to relax/study to
- [ ] autostart
  - [ ] 1password

- [ ] copr
  - [ ] update chroots/repositories
  - [ ] yq
  - [ ] concurrently
- [ ] korbenware
  - [ ] build/install copr package
- [ ] tmtui
- [ ] cockpit

## Avahi

Avahi is the thing that makes mDNS work. It might already Just Work, but in
case it doesn't...

<https://fedoramagazine.org/find-systems-easily-lan-mdns/>

This is more or less plug and play. Relevant settings are mostly going to be
`hostname` (configurable under `overview`)  and the short list of
things in `/etc/avahi/avahi-daemon.conf`.

Note, mDNS needs to be allowed by the firewall before it will work fully.

# Cockpit Stuff

* <https://github.com/45Drives/cockpit-file-sharing>
* Old but promising: <https://github.com/cyberorg/apsetup-cockpit>

## Audio

Things use pipewire by default, but it should pretend to be pulseaudio. Some
things to try:

* <https://github.com/patroclos/PAmix> - TUI mixer
* <https://github.com/cdemoulins/pamixer> - CLI mixer

Will eventually need to feed "line in" into output.
