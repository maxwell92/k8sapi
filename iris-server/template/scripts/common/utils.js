define([], function(){
	var yce = {};
	yce.preUrl = '';

	yce.http = function($http, method, url, param, success, error){
		$http[method](yce.preUrl + url, param, {headers: {'Content-Type': 'application/x-www-form-urlencoded'}})
		.success(function(data){
			success(data);
		})
		.error(function(){
			if(error && typeof error == 'function') return error();
			console.log('error');
		});
	};
    console.log('this is util.js')
	return yce;
});
