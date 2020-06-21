#!/bin/bash
ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -i inventory setup.yml
