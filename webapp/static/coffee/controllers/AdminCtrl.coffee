angular.module(APP_NAME).controller 'AdminCtrl', ($scope, $rootScope, $http, $timeout) ->
    $scope.admin = {}
    $scope.autoUpdate = false

    $scope.submitNewPlayer = (name) ->
        console.log 'name:', name
        data = 
            name: name
        $http.post("/api/admin/newtoon", data).then (response) ->
            console.log 'response:', response


    $scope.selectedToonChange = (toon) ->
        console.log 'selectedToonChange:', toon
        data =
            toonId: toon._id

        $rootScope.setMyToon toon
        $http.post("/api/activetoon", data).then (response) ->
            console.log 'response:', response.status, response.data

    $scope.showAllLocs = ->
        $http.get("/api/admin/alllocations").then (response) ->
            $rootScope.displayedLocations = response.data
    
    updateData = ->
        if !$scope.autoUpdate
            return
        $rootScope.fetchBundle()
        $timeout ->
            updateData()
        , 2000

    $scope.toggleAutoUpdate = (auto) ->
        $scope.autoUpdate = !$scope.autoUpdate
        updateData()

    $http.get("/api/admin/alltoons").then (response) ->
        $rootScope.allToons = response.data
        if $rootScope.myToon
            $scope.admin.selectedToon = $rootScope.myToon.Id

