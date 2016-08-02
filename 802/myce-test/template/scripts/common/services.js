define([		
	'../pages/main/service',
	'../pages/dashboard/service',
	'../pages/appManage/service',
	'../pages/rbdManage/service',
	'../pages/costManage/service',
	'../pages/extensions/service',
	'../pages/imageManage/service'
	], function(mainService, dashboardService, appManageService, rbdManageService, costManageService, extensionsService, imageManageService){

		'use strict';
		//获取全部service
		var args = Array.prototype.slice.call(arguments, 0);		

		//services[]
		var services = args;

		//创建service
		var init = function(){				
			_.each(services, function(service, index, services){
				angular.module(service.module).factory(service.name,['$http', function($http){
					return service.getApis($http);
				}]);					
			});						
		};

		return {
			init: init
		};
	}
);