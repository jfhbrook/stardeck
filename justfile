set dotenv-load := true

email := 'josh.holbrook@gmail.com'
destination := 'josh@stardeck.local'

default:
  @just --list

# Run setup steps
setup:
  cd ./stardeck-playbook && npm install

# Lint everything
lint:
  just -f ./perl/Stardeck/justfile lint
  cd ./stardeck-playbook && npm run lint
  shellcheck scripts/*.sh

# Format everything
format:
  just -f ./perl/Stardeck/justfile format
  cd ./stardeck-playbook && npm run format

# Link tool
link:
  ./scripts/link-justfile.sh ./justfile stardeck

# Run updates
update:
  @bash ./scripts/playbook-dependencies.sh
  @just playbook

# Run stardeck-playbook
playbook *ARGV:
  cd ./stardeck-playbook && sudo -E node main.mjs {{ ARGV }}

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
