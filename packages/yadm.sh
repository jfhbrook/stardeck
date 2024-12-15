DNF_PACKAGES=(yadm)

function install_yadm {
  sudo dnf config-manager --add-repo https://download.opensuse.org/repositories/home:TheLocehiliosan:yadm/Fedora_40/home:TheLocehiliosan:yadm.repo
  sudo dnf install -y yadm
}
