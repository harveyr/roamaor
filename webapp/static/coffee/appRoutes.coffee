angular.module(APP_NAME).config ['$routeProvider', '$locationProvider', ($routeProvider, $locationProvider) ->

    urlBase = '/app'

    url = (url) ->
        "#{urlBase}/url"

    $routeProvider
        .when(urlBase, {
            controller: 'HomeCtrl'
            templateUrl: 'static/partials/home.html',
        })

    $locationProvider
        .html5Mode(true)
        .hashPrefix('!');
]
