define([], function(){
	'use strict';

	var router = {
		dashboard: {
			url: '/dashboard',
			templateUrl: 'views/dashboard/dashboard.html',
			controller: 'dashboardController'
		}
	};

	return router;
});