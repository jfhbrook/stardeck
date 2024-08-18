set dotenv-load := true

default:
  @just --list

# Set everything up
setup:
  @just install

# Format playbooks
format:
  yamlfmt *.yml ./group_vars/*.yml ./host_vars/*.yml ./inventory/*.yml ./tasks/*.yml ./roles/raspi-* ./roles/rc2014*

# Lint playbooks and scripts
lint:
  shellcheck scripts/*.sh
  ./scripts/with-all-playbooks.sh just playbook inventory --syntax-check
  ./scripts/with-all-playbooks.sh ansible-lint

# Get status
status:
  just check inventory/lil-nas-x.yml lil-nas-x.yml

# Run a playbook
playbook inventory *argv:
  ANSIBLE_CONFIG="$(pwd)/ansible.cfg" ansible-playbook -i '{{ inventory }}' {{ argv }}

# Check a playbook
check inventory *argv:
  @just playbook '{{ inventory }}' --check {{ argv }}

# Install ansible roles and collections
install *argv:
  ANSIBLE_CONFIG="$(pwd)/ansible.cfg" ansible-galaxy install -r requirements.yml {{ argv }}

# Run ansible-config
config action *argv:
  ANSIBLE_CONFIG="$(pwd)/ansible.cfg" ansible-config {{ action }} {{ argv }}

# Dump ansible facts for a target
facts target:
  ANSIBLE_CONFIG="$(pwd)/ansible.cfg" ansible -i inventory '{{ target }}' -m ansible.builtin.setup

# Create a new role
new name:
  export ANSIBLE_CONFIG="$(pwd)/ansible.cfg"; cd roles && ansible-galaxy init '{{ name }}'
