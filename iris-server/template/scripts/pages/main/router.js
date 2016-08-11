define([], function(){
	'use strict';

	var router = {
		index: {
			url: '/index',
			templateUrl: 'views/main/login.html',
			controller: 'mainController'
		},
		main: {
			url: '/main',
			templateUrl: 'views/main/main.html',
			controller: 'mainController'
		}		
	};

	return router;
});