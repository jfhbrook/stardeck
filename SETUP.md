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
- [x] yadm
  - install yadm
  - create new stardeck-dotfiles repo
  - initialize yadm
- [x] sshd
- [x] set up dotfiles
- [x] vanilla vim
- [x] neovim
- [x] kitty
  - install with dnf
- [x] mononoki nerd fonts
  - <https://blog.khmersite.net/p/installing-nerd-font-on-fedora/>
- [x] starship
  - installed with script
- [x] CLI tools
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
- [x] rustup
  - watchexec
- [ ] asdf
  - [ ] asdf-ruby
  - [ ] asdf-java
- [ ] uv (python)
- [ ] volta
- [ ] coc.nvim extensions
  - typescript
  - python
  - go
- [ ] spruce up bashrc
- [ ] set up update scripts

- [ ] edge
- [ ] edge apps
  - [ ] youtube
  - [ ] youtube music

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
