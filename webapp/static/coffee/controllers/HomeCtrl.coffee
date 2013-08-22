angular.module(APP_NAME).controller 'HomeCtrl', ($scope, $rootScope) ->

    $scope.markerStyle =
        top: '20px'
        left: '20px'

    $scope.locationStyle =
        top: '50px'
        left: '50px'

    $scope.mapClick = ($event) ->
        console.log('here', $event);
        $scope.markerStyle.top = $event.y + 'px'
        $scope.markerStyle.left = $event.x + 'px'
