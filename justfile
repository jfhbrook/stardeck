set dotenv-load := true

default:
  @just --list

download-win3xsounds:
  curl -L https://winsounds.com/downloads/Windows3x.zip -o sounds/Windows3x.zip
  cd sounds && unzip Windows3x.zip
  rm -f sounds/Windows3x.zip

download-winxpsounds:
  curl -L 'https://archive.org/compress/windowsxpstartup_201910/formats=VBR%20MP3&file=/windowsxpstartup_201910.zip' -o sounds/windowsxp.zip
  cd sounds && unzip windowsxp.zip
  rm -f sounds/windowsxp.zip

download-btsounds:
  curl -L https://www.myinstants.com/media/sounds/the-bluetooth-device-is-ready-to-pair.mp3 -o sounds/the-bluetooth-device-is-ready-to-pair.mp3
  curl -L https://www.myinstants.com/media/sounds/the-bluetooth-device-its-connected-succesfull.mp3 -o sounds/the-bluetooth-device-its-connected-succesfull.mp3

download-floppysounds:
  yt-dlp --extract-audio https://www.youtube.com/watch?v=o_quPha61D0 --audio-format wav --output sounds/floppy-sounds.wav
  @just cut-floppysounds

cut-floppysounds:
  ffmpeg -y -ss 23 -t 7 -i sounds/floppy-sounds.wav sounds/floppy-start.mp3

download-sounds: download-win3xsounds download-winxpsounds download-btsounds download-floppysounds

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
