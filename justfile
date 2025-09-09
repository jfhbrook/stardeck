set dotenv-load := true

email := 'josh.holbrook@gmail.com'
destination := 'josh@stardeck.local'

releasever := '42'

default:
  @just --list

# Run setup steps
setup:
  cd ./playbook && npm install

# Lint everything
lint:
  cd ./playbook && npm run lint
  shellcheck scripts/*.sh

# Format everything
format:
  cd ./playbook && npm run format
  just -f stardeckd/justfile format
  just -f stardeckctl/justfile format

# Link tool
link:
  ./scripts/link-justfile.sh ./justfile stardeck

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

# Control loopback
loopback ACTION:
  @perl ./scripts/loopback.pl '{{ ACTION }}'

# Reset LCD brightness and contrast, and clear screen
lcd-reset:
  @bash ./scripts/lcd-reset.sh

# Display the LCD splash screen
lcd-splash: lcd-reset
  @bash ./scripts/lcd-splash.sh

# Stream notifications as newline separated JSON
notifications:
  @bash ./scripts/notifications.sh

# Logs for a service
logs SERVICE:
  journalctl -b -u '{{ SERVICE }}.service'

# Put the computer to sleep
nini:
  sudo systemctl suspend

dbus-services bus:
  if [[ ! '{{ bus }}' == 'session' ]] && [[ ! '{{ bus }}' == 'system' ]]; then echo "bus must be 'session' or 'system'"; exit 1; fi
  dbus-send --{{ bus }} --print-reply --dest=org.freedesktop.DBus  /org/freedesktop/DBus org.freedesktop.DBus.ListNames

dbus-object bus dest:
  if [[ ! '{{ bus }}' == 'session' ]] && [[ ! '{{ bus }}' == 'system' ]]; then echo "bus must be 'session' or 'system'"; exit 1; fi
  dbus-send --{{ bus }} --dest={{ dest }} --print-reply "/" org.freedesktop.DBus.Introspectable.Introspect
