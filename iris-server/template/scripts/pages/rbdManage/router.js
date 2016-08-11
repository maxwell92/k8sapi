define([], function(){
	'use strict';

	var router = {
		rbdManage: {
			url: '/rbdManage',
			templateUrl: 'views/rbdManage/rbdManage.html',
			controller: 'rbdManageController'
		}
	};

	return router;
});