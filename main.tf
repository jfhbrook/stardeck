module "stardeck" {
  source      = "./modules/stardeck"
  hostname    = "localhost"
  config_file = "./stardeck.yml"
  vars = {
    ansible_user = "josh"
  }
}
