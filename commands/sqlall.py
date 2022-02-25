# -*- coding: utf-8 -*-

import os
import json
import shutil

from command_base import BaseCommand

class Command(BaseCommand):
	help = "sqlall"
	args = ''

	def handle(self, **options):
		cmd = "BEEGO_RUNMODE=dev go run commands/cmd.go orm sqlall"
		print 'run> ', cmd
		os.system(cmd)
