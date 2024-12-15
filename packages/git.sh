DNF_PACKAGES=(git)

function install_git {
  git config --global user.name "Josh Holbrook"
  git config --global user.email "josh.holbrook@gmail.com"
  git config --global init.defaultBranch main
}
