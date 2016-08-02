define([
	'utils'
	], function(utils){
	'use strict';

	var getApis = function($http){
		var apis = {};

		return apis;
	};	

	var services = {
		module: 'rbdManage',
		name: 'rbdManageService',
		getApis: getApis
	};

	return services;
});