define([], function(){
	'use strict';

	var router = {
		extensions: {
			url: '/extensions',
			templateUrl: 'views/extensions/extensions.html',
			controller: 'extensionsController'
		},
		service: {
			url: '/service',
			templateUrl: 'views/extensions/service.html',
			controller: 'extensionsController'
		},
		endpoint: {
			url: '/endpoint',
			templateUrl: 'views/extensions/endpoint.html',
			controller: 'extensionsController'
		}
	};

	return router;
});