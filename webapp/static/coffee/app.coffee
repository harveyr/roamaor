app = angular.module(APP_NAME, [
    "#{APP_NAME}.directives"
    "#{APP_NAME}.services"
]).run ($rootScope, $http, GameConstants) ->
    $rootScope.appName = "Roamoar"
    $rootScope.myToon = null
    $rootScope.displayedLocations = []

    $rootScope.alertUser = (string) ->
        $rootScope.userAlert = string

    $rootScope.setMyToon = (toon) ->
        if !toon
            throw "setMyToon received null toon"
        $rootScope.myToon = toon

    updateDisplayedLocations = (locations) ->
        notDisplayed = []
        displayedIds = _.pluck $rootScope.displayedLocations, 'Id'
        _.each locations, (loc) ->
            if displayedIds.indexOf(loc.Id) == -1 
                notDisplayed.push loc

        if notDisplayed.length > 0
            $rootScope.displayedLocations = $rootScope.displayedLocations.concat(notDisplayed)

    $rootScope.fetchBundle = ->
        $http.get("/api/bootstrap").then (response) ->
            if !response.data.success
                $rootScope.alertUser "Failed to fetch bootstrap bundle. (#{response.data.reason})"
                return

            console.log 'bundle:', response.data
            $rootScope.worldHeight = response.data.worldHeight
            $rootScope.worldWidth = response.data.worldWidth
            $rootScope.myUser = response.data.user
            $rootScope.toonLogs = response.data.toonLogs

            GameConstants.set "logTypes", response.data.logTypes
            GameConstants.set "locationTypes", response.data.locationTypes

            if response.data.toon 
                $rootScope.setMyToon(response.data.toon)
                updateDisplayedLocations(response.data.visited)

    $rootScope.fetchBundle()
