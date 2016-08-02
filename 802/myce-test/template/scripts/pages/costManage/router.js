define([], function(){
	'use strict';

	var router = {
		costManage: {
			url: '/costManage',
			templateUrl: 'views/costManage/costManage.html',
			controller: 'costManageController'
		}
	};

	return router;
});