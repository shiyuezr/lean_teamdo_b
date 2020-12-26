# -*- coding: utf-8 -*-

import os
import sys

from util.db_util import SQLService  # important !!

path = os.path.abspath(os.path.join('.', '..'))
sys.path.insert(0, path)
reload(sys)
sys.setdefaultencoding('utf8')

import unittest
from bdd import util as bdd_util
from util import user_util

clean_modules = []
def __clear_all_app_data():
	"""
	清空应用数据
	"""
	if len(clean_modules) == 0:
		for clean_file in os.listdir('./features/clean'):
			if clean_file.startswith('__'):
				continue

			if clean_file.startswith('.'):
				continue

			module_name = 'features.clean.%s' % clean_file[:-3]
			module = __import__(module_name, {}, {}, ['*',])	
			clean_modules.append(module)

	for clean_module in clean_modules:
		clean_module.clean()

def init_mysql_trigger():
	"""
	创建mysql函数和触发器
	"""



def before_all(context):
	#创建test case，使用assert
	context.tc = unittest.TestCase('__init__')
	bdd_util.tc = context.tc

	init_mysql_trigger()


def after_all(context):
	pass


def before_scenario(context, scenario):
	context.scenario = scenario
	context.execute_steps(u"Given 重置服务")


def after_scenario(context, scenario):
	pass

