# -*- coding: utf-8 -*-
import json
import time
import logging
import requests
from datetime import datetime, timedelta


import db_util
exec_sql = db_util.exec_sql

tc = None

def convert_to_same_type(a, b):
	def to_same_type(target, other):
		target_type = type(target)
		other_type = type(other)
		if other_type == target_type:
			return True, target, other

		if (target_type == int) or (target_type == float):
			try:
				other = target_type(other)
				return True, target, other
			except:
				return False, target, other

		return False, target, other

	is_success, new_a, new_b = to_same_type(a, b)
	if is_success:
		return new_a, new_b
	else:
		is_success, new_b, new_a = to_same_type(b, a)
		if is_success:
			return new_a, new_b

	return a, b


###########################################################################
# assert_dict: 验证expected中的数据都出现在了actual中
###########################################################################
def assert_dict(expected, actual):
	global tc
	is_dict_actual = isinstance(actual, dict)
	for key in expected:
		expected_value = expected[key]
		if is_dict_actual:
			actual_value = actual[key]
		else:
			actual_value = getattr(actual, key)

		if isinstance(expected_value, dict):
			assert_dict(expected_value, actual_value)
		elif isinstance(expected_value, list):
			assert_list(expected_value, actual_value)
		else:
			try:
				expected_value, actual_value = convert_to_same_type(expected_value, actual_value)
				tc.assertEquals(expected_value, actual_value)
			except Exception, e:
				items = ['\n<<<<<', 'e: %s' % str(expected), 'a: %s' % str(actual), 'key: %s' % key, e.args[0], '>>>>>\n']
				e.args = ('\n'.join(items),)
				raise e


def assert_list(expected, actual, options=None):
	"""
	验证expected中的数据都出现在了actual中
	"""
	global tc
	try:
		tc.assertEquals(len(expected), len(actual), 'list length DO NOT EQUAL: %d != %d' % (len(expected), len(actual)))
	except:
		if options and 'key' in options:
			print '      Outer Compare Dict Key: ', options['key']
		raise

	for i in range(len(expected)):
		expected_obj = expected[i]
		actual_obj = actual[i]
		if isinstance(expected_obj, dict):
			assert_dict(expected_obj, actual_obj)
		else:
			expected_obj, actual_obj = convert_to_same_type(expected_obj, actual_obj)
			tc.assertEquals(expected_obj, actual_obj)



def assert_api_call(response, context):
	if context.text:
		input_data = json.loads(context.text)
		if isinstance(input_data, dict) and 'error' in input_data:
			assert_api_call_failed(response, input_data['error'])
			return False
		elif isinstance(input_data, list) and 'error' in input_data[0]:
			assert_api_call_failed(response, input_data[0]['error'])
			return False
		else:
			assert_api_call_success(response)
			return True
	else:
		assert_api_call_success(response)
		return True


###########################################################################
# assert_api_call_success: 验证api调用成功
###########################################################################
def assert_api_call_success(response):
	if 200 != response.body['code']:
		buf = []
		buf.append('>>>>>>>>>>>>>>> response <<<<<<<<<<<<<<<')
		buf.append(str(response))
		logging.error("API calling failure: %s" % '\n'.join(buf))
	assert 200 == response.body['code'], "code != 200, call api FAILED234!!!!"


###########################################################################
# assert_api_call_failed: 验证api调用失败
###########################################################################
def assert_api_call_failed(response, expected_err_code=None):
	if 200 == response.body['code']:
		buf = []
		buf.append('>>>>>>>>>>>>>>> response <<<<<<<<<<<<<<<')
		buf.append(str(response))
		logging.error("API calling not expected: %s" % '\n'.join(buf))
	assert 200 != response.body['code'], "code == 200, call api NOT EXPECTED!!!!"
	if expected_err_code:
		actual_err_code = str(response.body['errCode'])
		assert expected_err_code in actual_err_code, "errCode != '%s', error code FAILED!!!" % expected_err_code



###########################################################################
# assert_expected_list_in_actual: 验证expected中的数据都出现在了actual中
###########################################################################
def assert_expected_list_in_actual(expected, actual):
	global tc

	for i in range(len(expected)):
		expected_obj = expected[i]
		actual_obj = actual[i]
		if isinstance(expected_obj, dict):
			assert_dict(expected_obj, actual_obj)
		else:
			try:
				tc.assertEquals(expected_obj, actual_obj)
			except Exception, e:
				items = ['\n<<<<<', 'e: %s' % str(expected), 'a: %s' % str(actual), 'key: %s' % key, e.args[0], '>>>>>\n']
				e.args = ('\n'.join(items),)
				raise e