# -*- coding: utf-8 -*-
import MySQLdb as mysql

def get_db_name():
	with open('conf/dev.app.conf', 'rb') as f:
		for line in f:
			if "${_DB_NAME||" in line:
				beg = line.find('||')+2
				end = line.find('}', beg)

				return line[beg:end]

def exec_sql(sql, params=None):
	params = params or {}
	db_name = get_db_name()
	DB = mysql.connect(host='localhost',user='root',passwd='root',db=db_name,charset='utf8')
	cursor = DB.cursor()
	try:
		print ('exec sql======>', sql, params)
		cursor.execute(sql, params)
		DB.commit()
		cursor.description = cursor.description or []
		columns = [col[0] for col in cursor.description]
		results = cursor.fetchall()
		return [dict(zip(columns, row)) for row in results]
	except Exception as e:
		print ('[db_util] exception!!!!!!!!!', e)
	finally:
		cursor.close()
		DB.close()