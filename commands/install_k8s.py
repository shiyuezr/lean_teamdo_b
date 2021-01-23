# -*- coding: utf-8 -*-

import os
import json
import shutil
import subprocess
import copy

import servicecli

from command_base import BaseCommand

class Command(BaseCommand):
	help = "install_k8s [namespace] [mode]\n\tnamespace:\tk8s namespace\n\t     mode:\trest, cron\n"
	args = ''

	def handle(self, namespace, mode, **options):
		servicecli.install_k8s(namespace, mode)

