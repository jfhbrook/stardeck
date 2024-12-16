STARDECK_BIN='stardeck'

function install_stardeck {
  mkdir -p ~/.local/bin

  if [ ! -f ~/.local/bin/stardeck ]; then
    echo "#!/usr/bin/env bash
cd '$(pwd)' && exec just "'"$@"' \
      > ~/.local/bin/${STARDECK_BIN}
    chmod +x ~/.local/bin/${STARDECK_BIN}
  fi
}

function remove_stardeck {
  rm ~/.local/bin/${STARDECK_BIN}
}
