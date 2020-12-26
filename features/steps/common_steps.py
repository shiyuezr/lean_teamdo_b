# -*- coding: utf-8 -*-
import json
from datetime import datetime
from behave import *
from bdd import client as bdd_client
from bdd import util as bdd_util
from util import user_util, db_util
import settings

from util.db_util import SQLService

@given(u"重置服务")
def step_impl(context):
	from features.bdd.client import RestClient
	rest_client = RestClient()
	response = rest_client.put('dev.bdd_reset')
	bdd_util.assert_api_call_success(response)
