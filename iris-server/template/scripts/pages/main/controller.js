/**
 * Created by Jora on 2016/7/29.
 */
define([
        'base64'
    ], function(Base64){
        'use strict';

        var ctrl = ['$scope','$state','mainService', function($scope, $state,mainService){
            $scope.$state = $state;
            $state.go('main.dashboard');
            $scope.data = {
                showSubnav : []
            };

            mainService.getNavlist(null,function(data){
                $scope.navList = data.list;
            });

            $scope.showSubnav = function(index){
                $scope.data.showSubnav[index] = !$scope.data.showSubnav[index];
            };

        }];



        var controllers = [
            {module: 'yce-manage', name: 'mainController', ctrl: ctrl}
        ];

        return controllers;
    }
);