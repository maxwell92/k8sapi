define([
	'utils'
	], function(utils){
	'use strict';

	var getApis = function($http){
		var apis = {};

		return apis;
	};	

	var services = {
		module: 'dashboardManage',
		name: 'dashboardService',
		getApis: getApis
	};

	return services;
});