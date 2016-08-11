define([
	'utils'
	], function(utils){
	'use strict';

	var getApis = function($http){
		var apis = {};

		apis.getAppList = function(param, success, error){
			/*return utils.http($http, 'get', './getapplist', param, success, error);*/
			return utils.http($http, 'get', '/getapplist', param, success, error);
		};

		return apis;
	};	

	var services = {
		module: 'appManage',
		name: 'appManageService',
		getApis: getApis
	};

    console.log('In service.js: appManage service.js')
    console.log('In service.js: ' + getApis)
	return services;
});
