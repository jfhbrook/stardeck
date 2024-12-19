set dotenv-load := true

destination := 'josh@stardeck.local'

default:
  @just --list

# Set up the environment
setup:
  @just -f ./playbooks/justfile install --force
  # just download-sounds

# Lint everything
lint:
  shellcheck scripts/*.sh

# Format everything
format:
  terraform fmt -recursive
  yamlfmt *.yml

# Link tool
link:
  ./scripts/link-justfile.sh ./justfile stardeck

# Log into stardeck
login:
  ssh '{{ destination }}'

# Get the status of the repo and of yadm
status:
  git status
  yadm status

# Update stardeck
apply:
  @bash ./scripts/apply.sh

# Control loopback
loopback CMD:
  @bash ./scripts/loopback.sh '{{ CMD }}'

# Change samba password
smbpasswd USER:
  sudo smbpasswd -U '{{ USER }}'

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
