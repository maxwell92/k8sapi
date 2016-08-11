define([], function(){
	'use strict';

	var router = {
		appManage: {
			url: '/appManage',
			templateUrl: 'views/appManage/appManage.html',
			controller: 'appManageController'
		},
		deployment: {
			url: '/deployment',
			templateUrl: 'views/appManage/deployment.html',
			controller: 'appManageController'
		},
		rollback: {
			url: '/rollback',
			templateUrl: 'views/appManage/rollback.html',
			controller: 'appManageController'
		},
		rollup: {
			url: '/rollup',
			templateUrl: 'views/appManage/rollup.html',
			controller: 'appManageController'
		},
		cancel: {
			url: '/cancel',
			templateUrl: 'views/appManage/cancel.html',
			controller: 'appManageController'
		},
		history: {
			url: '/history',
			templateUrl: 'views/appManage/history.html',
			controller: 'appManageController'
		}
	};

	return router;
});