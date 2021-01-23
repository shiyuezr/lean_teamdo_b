# -*- coding: utf-8 -*-

import threading

import requests
import time


DEST_USER_ID = {
	0: 826,
	1: 14,
	2: 11,
}

class CurlThread(threading.Thread):
	"""
	transfer并发测试
	"""
	def __init__(self, index):
		super(CurlThread, self).__init__()
		self.index = index

	def run(self):
		print 'start thread {}...'.format(self.index)
		for i in range(50):
			requests.get("http://127.0.0.1:6001/dev/all_corps")
		print 'finish thread {}...'.format(self.index)


def run():
	threads = []

	for i in xrange(100):
		thread = CurlThread(i)
		thread.start()
		time.sleep(0.4)
		threads.append(thread)

	for t in threads:
		t.join()

	print 'done...'

if __name__ == '__main__':
	run()