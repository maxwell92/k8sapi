define([
	'jQuery',
	'Underscore',
	'Angular',
	'uiRouter',
	'mock',
	'mockAngular',
	'mockData',
	'./common/constant',
	'./common/factory',
	'./common/filter',
	'./common/directive',
	'./common/services',
	'./common/routers',
	'./common/controllers'
], function($, _, angular, uiRouter, Mock, mockAngular, mockData, constant, factory, filter, directive, services, routers, controllers){
	'use strict';

	//通用模块
	angular.module('common', []);
	angular.module('dashboardManage', ['common']);
	angular.module('appManage', ['common']);
	angular.module('rbdManage', ['common']);
	angular.module('costManage', ['common']);
	angular.module('extensionsManage', ['common']);
	angular.module('imageManage', ['common']);


	//主应用
	var app = angular.module('yce-manage', ['ui.router', 'common', 'dashboardManage', 'appManage', 'rbdManage', 'costManage', 'extensionsManage', 'imageManage']);
	//路由引导
	Mock.mockjax(app);

	//常量
	constant.init();

	//方法
	factory.init();

	//过滤
	filter.init();

	//指令
	directive.init();

	//服务
	services.init();

	//路由引导
	routers.init(app);

	//控制器引导
	controllers.init();

	angular.bootstrap(window.document, ['yce-manage']);
});