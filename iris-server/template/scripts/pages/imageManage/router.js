define([], function(){
	'use strict';

	var router = {
		imageManage: {
			url: '/imageManage',
			templateUrl: 'views/imageManage/imageManage.html',
			controller: 'imageManageController'
		},
		search: {
			url: '/search',
			templateUrl: 'views/imageManage/search.html',
			controller: 'imageManageController'
		},
		delete: {
			url: '/delete',
			templateUrl: 'views/imageManage/delete.html',
			controller: 'imageManageController'
		}
	};

	return router;
});