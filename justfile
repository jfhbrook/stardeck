set dotenv-load := true

releasever := '42'

default:
  @just --list

# Run setup steps
setup:
  cd ./playbook && npm install

# Lint everything
lint:
  find . -name '*.go' -exec dirname {} ';' | sort -u | xargs go vet
  cd ./playbook && npm run lint
  just -f ./kwin/justfile lint
  shellcheck scripts/*.sh

# Format everything
format:
  find . -name '*.go' -exec go fmt {} ';'
  cd ./playbook && npm run format
  just -f ./kwin/justfile format

build:
  mkdir -p bin
  go build -o bin/stardeckd ./cmd/stardeckd/main.go
  go build -o bin/stardeckctl ./cmd/stardeckctl/main.go
  just -f ./kwin/justfile build

stardeckd *argv:
  ./scripts/stardeckd.sh {{ argv }}

stardeckctl *argv:
  @go run ./cmd/stardeckctl/main.go {{ argv }}

# Link tool
link:
  ./scripts/link-justfile.sh ./justfile stardeck

# Install local dependencies
install:
  cd playbook && npm i
  cd kwin && npm i

# Enable development stardeck service
enable:
  cp ./systemd/user/* ~/.config/systemd/user/
  systemctl --user enable stardeck.service
  @just start

# Start stardeck service
start:
  systemctl --user start stardeck.service

# Disable development stardeck service
disable:
  @just stop
  for file in ./systemd/user/*; do rm ~/.config/systemd/user/"$(basename "${file}")"; done

# Stop stardeck service
stop:
  systemctl --user stop stardeck.service

# Run updates
update:
  @bash ./scripts/playbook-dependencies.sh
  @just playbook

# Steps to run before upgrading Fedora
pre-upgrade:
  sudo dnf upgrade --refresh
  sudo reboot

# Upgrade Fedora
upgrade:
  sudo dnf system-upgrade download --releasever={{ releasever }}
  sudo dnf5 offline reboot

# Steps to run after upgrading Fedora
post-upgrade:
  bash ./scripts/post-upgrade.sh

# Clean up extra packages
cleanup-extras:
  bash ./scripts/cleanup-extras.sh

# Clean up old kernels
cleanup-kernels:
  bash ./scripts/cleanup-kernels.sh

# Clean up dangling symlinks in /usr
cleanup-symlinks:
  bash ./scripts/cleanup-symlinks.sh

# Run playbook
playbook *ARGV:
  cd ./playbook && sudo -E node main.mjs {{ ARGV }}

# Scan music for mopidy
scan-music:
  sudo mopidyctl local scan

# Reset LCD brightness and contrast, and clear screen
lcd-reset:
  @just stardeckctl reset lcd

# Display the LCD splash screen
lcd-splash: lcd-reset
  @just stardeckctl set lcd splash

# Logs for a service
logs *argv:
  @just stardeckctl logs {{ argv }}

# Put the computer to sleep
nini:
  sudo systemctl suspend

dbus-services bus:
  if [[ ! '{{ bus }}' == 'session' ]] && [[ ! '{{ bus }}' == 'system' ]]; then echo "bus must be 'session' or 'system'"; exit 1; fi
  dbus-send --{{ bus }} --print-reply --dest=org.freedesktop.DBus  /org/freedesktop/DBus org.freedesktop.DBus.ListNames

dbus-object bus dest:
  if [[ ! '{{ bus }}' == 'session' ]] && [[ ! '{{ bus }}' == 'system' ]]; then echo "bus must be 'session' or 'system'"; exit 1; fi
  dbus-send --{{ bus }} --dest={{ dest }} --print-reply "/" org.freedesktop.DBus.Introspectable.Introspect
