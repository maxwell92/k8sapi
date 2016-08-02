/**
 * Created by Jora on 2016/7/29.
 */
define([
        'base64'
    ], function(Base64){
        'use strict';

        var ctrl = ['$scope', 'appManageService', function($scope,appManageService){

            appManageService.getAppList(null,function(data){
                $scope.appList = data;
                console.log("In controller.js/ctl: " + data)
            });
        }];



        var controllers = [
            {module: 'appManage', name: 'appManageController', ctrl: ctrl}
        ];

        return controllers;
    }
);
