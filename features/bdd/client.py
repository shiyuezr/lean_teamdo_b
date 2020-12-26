# -*- coding: utf-8 -*-
import json
import requests

import settings

CUR_CONTEXT = None

class ApiResponse(object):
	"""
	api call的response
	"""
	def __init__(self, response):
		self.text = response.text.strip()
		self.raw_response = response

	@property
	def json_data(self):
		return json.loads(self.text)['data']

	@property
	def data(self):
		return json.loads(self.text)['data']

	@property
	def body(self):
		return json.loads(self.text)

	@property
	def json(self):
		return json.loads(self.text)

	@property
	def is_success(self):
		"""
		判断该次请求是否成功
		"""
		r = self.raw_response
		if r.status_code != 200:
			assert False, "http status code is %d, http call is FAILED!!!!" % r.status_code

		if 'html>' in self.text:
			assert False, "NOT a valid json string, call api FAILED!!!!"

		if self.json['code'] == 200:
			return True
		else:
			print '-*-' * 20
			print self.text
			print '-*-' * 20
			assert 200 == self.json['code'], "json[code] != 200, call api FAILED!!!!"
			return False

	@property
	def is_fail(self):
		if self.json["code"] != 200:
			return True
		else:
			return False

	def __repr__(self):
		return self.text

class RestClient(object):
	"""
	访问rest资源的client
	"""
	def __init__(self):
		self.jwt_token = None

	def __get_url(self, type, resource):
		service = None
		if ':' in resource:
			service, resource = resource.split(':')

		pos = resource.rfind('.')
		if pos == -1:
			raise RuntimeError('INVALID RESOURCE: %s' % resource)
		app = resource[:pos].replace('.', '/')
		app_resource = resource[pos+1:]

		if service:
			url = 'http://devapi.vxiaocheng.com/%s/%s/%s/' % (service, app, app_resource)
		else:
			url = 'http://127.0.0.1:%s/%s/%s/' % (settings.SERVICE_PORT, app, app_resource)
		if type == 'put':
			url = '%s?_method=put' % url
		elif type == 'delete':
			url = '%s?_method=delete' % url

		print "url: ", url

		return url

	def get(self, resource, data={}, context=None):
		url = self.__get_url('get', resource)

		headers = {}
		if self.jwt_token:
			headers = {
				'AUTHORIZATION': self.jwt_token
			}
		r = requests.get(url, data, headers=headers)
		return ApiResponse(r)

	def post(self, resource, data={}, context=None):
		url = self.__get_url('post', resource)

		headers = {}
		if self.jwt_token:
			headers = {
				'AUTHORIZATION': self.jwt_token
			}
		r = requests.post(url, data, headers=headers)
		return ApiResponse(r)

	def put(self, resource, data={}, context=None):
		url = self.__get_url('put', resource)

		headers = {}
		if self.jwt_token:
			headers = {
				'AUTHORIZATION': self.jwt_token
			}
		r = requests.post(url, data, headers=headers)
		return ApiResponse(r)

	def delete(self, resource, data={}, context=None):
		url = self.__get_url('delete', resource)

		headers = {}
		if self.jwt_token:
			headers = {
				'AUTHORIZATION': self.jwt_token
			}
		r = requests.post(url, data, headers=headers)
		return ApiResponse(r)

class Obj(object):
	def __init__(self):
		pass

def login(type='backend', user='manager', password=None, **kwargs):
	if not password:
		password = '55e421ee9bdc9d9f6b6c6518E590b0ee'

	client = RestClient()
	if type == 'app':
		resp = client.put("gskep:login.logined_bdd_user", {
			'name': user,
		})
		assert resp.is_success
	elif type == 'backend':
		data = {
			"username": user if user not in ('xiaocheng', 'manager') else 'manager_小诚科技',
			"password": password
		}
		resp = client.put('gskep:login.logined_corp_user', data)
		assert resp.is_success

	client.jwt_token = resp.data['sid']
	client.cur_platform_id = resp.data['pid']
	client.cur_user_id = resp.data['id']
	user = Obj()
	user.id = resp.data['id']

	if 'context' in kwargs:
		context = kwargs['context']
		if context:
			context.client = client
			context.user = user

	return client