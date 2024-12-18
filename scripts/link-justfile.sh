#!/usr/bin/env bash

set -euxo pipefail

JUSTFILE="$(realpath "${1}")"
NAME="${2}"

if [ ! -f "${JUSTFILE}" ]; then
  echo "Justfile not found: ${JUSTFILE}"
  exit 1
fi

if [ -f ~/.local/bin"${NAME}" ]; then
  echo "Script already installed."
  exit 1
fi

# shellcheck disable=SC1078,SC1079,SC2027,SC2086
echo "#!/usr/bin/env bash
exec just -f "${JUSTFILE}" "'"$@"' \
  > ~/.local/bin/"${NAME}" \
  && chmod +x ~/.local/bin/"${NAME}"
