resource "ansible_host" "dokku" {
  name      = var.hostname
  variables = var.vars
}

resource "ansible_playbook" "dokku" {
  name      = ansible_host.dokku.name
  playbook  = "${path.module}/main.yml"
  var_files = [var.config_file]
}
