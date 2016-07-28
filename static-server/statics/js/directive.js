function pageTitle($rootScope){
	return {
		link:function(scope,element){
			var func=function(event,toState){
					//console.log(event)
					var title='1406E'
					if(toState.data && toState.data.title) title += toState.data.title;
				element.text(title)
				}
			$rootScope.$on('$stateChangeStart',func)
		}
	}
}

function createDom($rootScope){
	return {
		restrict:'A',
		link:function(scope,element){
			console.log(element);
			var str="";
			for(var i=0;i<scope.arr.length;i++){
				str += '<p>id:'+scope.arr[i]+'</p>'
			}
			element.append(str);
		/*		
			var func = function(event,toState){
				console.log(toState);
				var str='';
				for(var i=0;i<toState.data.length;i++){
					str += '<li>'+ toState.data[i] +'</li>';
				}
				element.append(str);
			}
			$rootScope.$on('$stateChangeStart',func)*/
		}
	}
}

angular.module('app')
	.directive('pageTitle',pageTitle)
	.directive('createDom',createDom)