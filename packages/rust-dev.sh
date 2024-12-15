DNF_PACKAGES=(rust-analyzer)

function install_rust-dev {
  curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
}
