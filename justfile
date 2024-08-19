set dotenv-load := true

default:
  @just --list

download-winsounds:
  curl -L https://winsounds.com/downloads/Windows3x.zip -o sounds/Windows3x.zip
  cd sounds && unzip Windows3x.zip
  rm -f sounds/Windows3x.zip

download-btsounds:
  curl -L https://www.myinstants.com/media/sounds/the-bluetooth-device-is-ready-to-pair.mp3 -o sounds/the-bluetooth-device-is-ready-to-pair.mp3
  curl -L https://www.myinstants.com/media/sounds/the-bluetooth-device-its-connected-succesfull.mp3 -o sounds/the-bluetooth-device-its-connected-succesfull.mp3

download-sounds: download-winsounds download-btsounds

# Set everything up
setup:
  cd playbooks && just install
  just download-sounds

# Lint playbooks and scripts
lint:
  shellcheck scripts/*.sh
  ./scripts/with-all-playbooks.sh ansible-lint

# Get status
status:
  just check main.yml

# Update
update:
  just playbook main.yml

# Run a playbook
playbook *argv:
  ANSIBLE_CONFIG="$(pwd)/ansible.cfg" ansible-playbook -i 'inventory.yml' --ask-become-pass {{ argv }}

# Check a playbook
check *argv:
  @just playbook inventory.yml --check {{ argv }}

# Install ansible roles and collections
install *argv:
  ANSIBLE_CONFIG="$(pwd)/ansible.cfg" ansible-galaxy install -r requirements.yml {{ argv }}

# Run ansible-config
config action *argv:
  ANSIBLE_CONFIG="$(pwd)/ansible.cfg" ansible-config {{ action }} {{ argv }}

# Dump ansible facts for a target
facts target:
  ANSIBLE_CONFIG="$(pwd)/ansible.cfg" ansible -i inventory.yml '{{ target }}' -m ansible.builtin.setup

# Create a new role
new name:
  export ANSIBLE_CONFIG="$(pwd)/ansible.cfg"; cd roles && ansible-galaxy init '{{ name }}'
