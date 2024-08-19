set dotenv-load := true

default:
  @just --list

download-win3x-sounds:
  mkdir -p sounds/win3x
  curl -L https://winsounds.com/downloads/Windows3x.zip -o sounds/win3x/Windows3x.zip
  cd sounds/win3x && unzip Windows3x.zip
  rm -f sounds/win3x/Windows3x.zip

download-winxp-sounds:
  mkdir -p sounds/winxp
  curl -L 'https://archive.org/compress/windowsxpstartup_201910/formats=VBR%20MP3&file=/windowsxpstartup_201910.zip' -o sounds/winxp/windowsxp.zip
  cd sounds/winxp && unzip windowsxp.zip
  rm -f sounds/winxp/windowsxp.zip

download-bt-sounds:
  mkdir -p sounds/bt
  curl -L https://www.myinstants.com/media/sounds/the-bluetooth-device-is-ready-to-pair.mp3 -o sounds/bt/the-bluetooth-device-is-ready-to-pair.mp3
  curl -L https://www.myinstants.com/media/sounds/the-bluetooth-device-its-connected-succesfull.mp3 -o sounds/bt/the-bluetooth-device-its-connected-succesfull.mp3

download-floppy-sounds:
  mkdir -p sounds/floppy
  yt-dlp --extract-audio https://www.youtube.com/watch?v=o_quPha61D0 --audio-format wav --output sounds/floppy/full.wav
  @just cut-floppy-sounds

cut-floppy-sounds:
  ffmpeg -y -ss 23 -t 7 -i sounds/floppy/full.wav sounds/floppy/start.mp3

download-videlectrix-sounds:
  mkdir -p sounds/videlectrix
  yt-dlp --extract-audio https://www.youtube.com/watch?v=xBmxHT2SUXg --audio-format wav --output sounds/videlectrix/full.wav
  @just cut-videlectrix-sounds

cut-videlectrix-sounds:
  ffmpeg -y -ss 3 -t '5.5' -i sounds/videlectrix/full.wav sounds/videlectrix/start.mp3

download-sounds:
  @just download-win3x-sounds
  @just download-winxps-ounds
  @just download-bt-sounds
  @just download-floppy-sounds
  @just download-videlectrix-sounds

play FILE:
  ffplay '{{FILE}}' -nodisp -autoexit

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
