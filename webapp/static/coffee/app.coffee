app = angular.module(APP_NAME, [
    "#{APP_NAME}.directives"
]).run ($rootScope, $http) ->
    $rootScope.appName = "Roamoar"
    $rootScope.myToon = null

    $rootScope.alertUser = (string) ->
        $rootScope.userAlert = string

    $rootScope.setMyToon = (toon) ->
        if !toon
            throw "setMyToon received null toon"
        $rootScope.myToon = toon

    $rootScope.fetchBundle = ->
        $http.get("/api/bootstrap").then (response) ->
            if !response.data.success
                $rootScope.alertUser "Failed to fetch bootstrap bundle. (#{response.data.reason})"
                return

            console.log 'bundle:', response.data
            $rootScope.worldHeight = response.data.worldHeight
            $rootScope.worldWidth = response.data.worldWidth
            $rootScope.myUser = response.data.user
            if response.data.toon 
                $rootScope.setMyToon(response.data.toon)
                $rootScope.locationsVisited = response.data.visited

    $rootScope.fetchBundle()
