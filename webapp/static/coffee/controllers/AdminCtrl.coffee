angular.module(APP_NAME).controller 'AdminCtrl', ($scope, $rootScope, $http) ->
    $scope.admin = {}

    $scope.submitNewPlayer = (name) ->
        console.log 'name:', name
        data = 
            name: name
        $http.post("/api/admin/newtoon", data).then (response) ->
            console.log 'response:', response


    $scope.selectedToonChange = (toon) ->
        console.log 'toon:', toon
        data =
            toonId: toon._id

        $rootScope.setMyToon toon
        $http.post("/api/activetoon", data).then (response) ->
            console.log 'response:', response.status, response.data

        
