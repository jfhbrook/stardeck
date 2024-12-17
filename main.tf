module "stardeck" {
  source      = "./modules/stardeck"
  config_file = "${path.root}/stardeck.yml"
  vars = {
    ansible_user = "josh"
  }
}
