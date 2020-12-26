# -*- coding: utf-8 -*-

import os
import pymysql
pymysql.install_as_MySQLdb()

SERVICE_NAME = 'teamdo'
SERVICE_PORT = 7001

#运行BDD时是否输出response
ENABLE_BDD_DUMP_REQUEST = True
ENABLE_BDD_DUMP_RESPONSE = True

# 是否处于BDD执行中，一般在features/environmens.py中设置
IS_UNDER_BDD = True

BDD_USER_PWD = 'test'


#
# mysql相关配置
#
DB_HOST = os.environ.get('_DB_HOST', 'db.dev.com')
DB_PORT = os.environ.get('_DB_PORT', '3306')
DB_NAME = os.environ.get('_DB_NAME', 'teamdo')
DB_USER = os.environ.get('_DB_USER', 'root')
DB_PASSWORD = os.environ.get('_DB_PASSWORD', 'root')

DATABASES = {
	'default': {
		'ENGINE': 'mysql+retry',
		'NAME': DB_NAME,
		'USER': DB_USER,
		'PASSWORD': DB_PASSWORD,
		'HOST': DB_HOST,
		'PORT': DB_PORT,
		'CONN_MAX_AGE': 100
	}
}

API_GATEWAY = 'devapi.vxiaocheng.com'
MODE = 'develop'
LOCK_ENGINE = 'dummy'
ENABLE_SQL_LOG = False