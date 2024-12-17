resource "ansible_host" "stardeck" {
  name      = "localhost"
  variables = var.vars
}

resource "ansible_playbook" "rpmfusion" {
  name      = ansible_host.stardeck.name
  playbook  = "${path.module}/rpmfusion.yml"
  var_files = [var.config_file]
}

resource "ansible_playbook" "update" {
  name      = ansible_host.stardeck.name
  playbook  = "${path.module}/update.yml"
  var_files = [var.config_file]
  depends_on = [
    ansible_playbook.rpmfusion
  ]
}

resource "ansible_playbook" "op" {
  name      = ansible_host.stardeck.name
  playbook  = "${path.module}/1password.yml"
  var_files = [var.config_file]
  depends_on = [
    ansible_playbook.update
  ]
}

resource "ansible_playbook" "git" {
  name      = ansible_host.stardeck.name
  playbook  = "${path.module}/git.yml"
  var_files = [var.config_file]
  depends_on = [
    ansible_playbook.update
  ]
}

resource "ansible_playbook" "asdf" {
  name      = ansible_host.stardeck.name
  playbook  = "${path.module}/asdf.yml"
  var_files = [var.config_file]
  depends_on = [
    ansible_playbook.git
  ]
}

resource "ansible_playbook" "yadm" {
  name      = ansible_host.stardeck.name
  playbook  = "${path.module}/yadm.yml"
  var_files = [var.config_file]
  depends_on = [
    ansible_playbook.git
  ]
}


