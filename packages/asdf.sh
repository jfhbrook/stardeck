DNF_PACKAGES=(openssl-devel libffi-devel libyaml-devel)

function install_asdf {
  if [ ! -d ~/.asdf ]; then
    git clone https://github.com/asdf-vm/asdf.git ~/.asdf --branch v0.14.1
  fi

  . "$HOME/.asdf/asdf.sh"

  asdf install ruby latest
  asdf global ruby latest
}
