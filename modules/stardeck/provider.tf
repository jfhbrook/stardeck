terraform {
  required_version = ">= 1.7.2"

  required_providers {
    ansible = {
      source  = "ansible/ansible"
      version = "1.3.0"
    }
  }
}
