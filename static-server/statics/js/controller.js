function index($scope,$rootScope){
    console.log('index');
    $rootScope.indexArr = ['fsaf','safa','kkk'];
}
function list($scope,$rootScope){
    $scope.arr=[1,2,3,4,5,6,7];
    console.log($rootScope);
}
function product($scope){
    console.log('product')
    $scope.getInfo=function(){
        window.location.href="http://baidu.com"
    }
}
function about($scope){
    
}
function dom($scope){
    
}
angular.module('app')
    .controller('index',index)
    .controller('list',list)
    .controller('product',product)
    .controller('about',about)
    .controller('dom',dom)