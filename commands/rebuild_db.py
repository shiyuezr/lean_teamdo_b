# -*- coding: utf-8 -*-

import os
import json
import shutil
import platform

from command_base import BaseCommand

SQL_TMPL = r"""
DROP DATABASE IF EXISTS %(database)s;
CREATE DATABASE %(database)s DEFAULT CHAR SET 'utf8';
GRANT ALL ON %(database)s.* to '%(user)s'@'%%' IDENTIFIED BY '%(password)s';
GRANT ALL ON %(database)s.* to '%(user)s'@'localhost' IDENTIFIED BY '%(password)s';
"""

def parse_db_conf():
	lines = []

	def extract_value(key):
		for line in lines:
			if key in line:
				beg = line.find('||')+2
				end = line.find('}', beg)

				return line[beg:end]

	with open('conf/dev.app.conf', 'rb') as f:
		for line in f:
			lines.append(line.strip())

	database = extract_value("${_DB_NAME||")
	user = extract_value("${_DB_USER||")
	password = extract_value("${_DB_PASSWORD||")

	return {
		"database": database,
		"user": user,
		"password": password
	}

class Command(BaseCommand):
	help = "rebuild_db"
	args = ''

	def handle(self, **options):
		can_run_command = False
		if os.name == 'nt':
			can_run_command = True
		if os.name == 'posix' and platform.system().lower() == 'darwin':
			can_run_command = True

		# run mysql cmd
		db_conf = parse_db_conf()
		sql = SQL_TMPL % db_conf
		cmd = 'echo "%s" | mysql -u root --password=root -h db.dev.com' % sql
		os.system(cmd)

		# run syncdb cmd
		os.system("python manage.py syncdb")
