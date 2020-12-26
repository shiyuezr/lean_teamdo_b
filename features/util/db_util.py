# coding: utf8

from eaglet.core.hack_peewee import connect

DB_DICT = dict()

def get_db(name):
	global DB_DICT
	if DB_DICT.get(name):
		return DB_DICT[name]
	url = 'mysql+retry://{name}:xiaocheng@127.0.0.1:3306/{name}'.format(name=name)
	print 'bdd connect db: %s' % url
	db = connect(url)
	db.connect()
	DB_DICT[name] = db
	return db

class SQLService(object):
	"""
	直接执行sql的服务
	"""
	__slots__ = (
		'__db', # database实例
	)

	def __init__(self, db):
		self.__db = db

	@classmethod
	def use(cls, db_name):
		db = get_db(db_name)
		if not db:
			raise Exception
		return cls(db)

	def execute_sql(self, sql):
		print sql
		cursor = self.__db.execute_sql(sql)
		return cursor

	def set_foreign_key_checks(self, value=True):
		sql = """
			set foreign_key_checks={};
		""".format(int(value))
		self.execute_sql(sql)