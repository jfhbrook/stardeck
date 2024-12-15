DNF_PACKAGES=(cockpit)

function install_cockpit {
  sudo systemctl enable --now cockpit.socket

  sudo firewall-cmd --add-service=cockpit
  sudo firewall-cmd --add-service=cockpit --permanent
}
