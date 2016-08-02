define([
	'utils'
	], function(utils){
	'use strict';

	var getApis = function($http){
		var apis = {};

		return apis;
	};	

	var services = {
		module: 'extensionsManage',
		name: 'extensionsService',
		getApis: getApis
	};

	return services;
});