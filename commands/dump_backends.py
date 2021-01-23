# -*- coding: utf-8 -*-

import os
import json
import shutil
import subprocess
import copy

import servicecli

from command_base import BaseCommand

class Command(BaseCommand):
	help = "dump_backends [namespace] \n\tnamespace:\tk8s namespace"
	args = ''
		
	def handle(self, namespace, **options):
		servicecli.dump_backends(namespace)