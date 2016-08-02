/**
 * Created by Jora on 2016/7/29.
 */
requirejs.config({
    baseUrl: 'scripts',
    paths: {
        jQuery: 'lib/jquery/jquery',
        Underscore: 'lib/underscore/underscore',
        jQueryUI: 'lib/jquery-ui/jquery-ui.min',
        Angular: 'lib/angular/angular',
        uiRouter: 'lib/angular-ui-router/angular-ui-router.min',
        base64: 'lib/base64/base64',
        utils: 'common/utils',
        mockAngular: 'mock/mock.angular',
        mock: 'mock/mock',
        mockData: 'mock/mockData'
    },
    shim: {
        'Angular': {
            deps: ['jQuery']
        },
        'uiRouter': {
            deps: ['Angular']
        },
        'jQueryUI': {
            deps: ['jQuery']
        },
        'mockAngular':{
            deps: ['mock']
        },
        'mockData':{
            deps: ['mock']
        }
    }
});

require(['app'], function(app){

});

