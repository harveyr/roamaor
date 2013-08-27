angular.module(APP_NAME).controller 'HomeCtrl', ($scope, $rootScope, $http, $timeout) ->

    mapColor = "#E7E7E7"
    toonRadius = 5
    scaleX = null
    scaleY = null

    svg = d3.select("svg")
        .attr("height", 500)
        # .style("background-color", mapColor)
        # .style("box-shadow", "1px 1px 1px #999")

    map = $(".game-map")
    mapWidth = map.width()
    mapHeight = map.height() - 9

    $scope.mapStyle =
        "background-size": "#{mapWidth}px #{mapHeight}px"

    lineFunc = d3.svg.line()
        .x((d) -> d.x)
        .y((d) -> d.y)
        .interpolate('linear')

    myDestPath = svg.append("path")
        .attr("id", "my-dest-path")

    toonSvgCoords = (toon) ->
        {
            x: toon.LocX + toonRadius
            y: svg.attr("height") - toon.LocY - toonRadius
        }


    mapToGameCoords = (inputX, inputY) ->
        svgWidthScale = parseInt(svg.style("width")) / $rootScope.worldWidth
        svgHeightScale = parseInt(svg.style("height")) / $rootScope.worldHeight
        scaled =
            x: inputX / svgWidthScale
            y: inputY / svgHeightScale

    gameToMapCoords = (inputX, inputY) ->
        svgWidthScale = parseInt(svg.style("width")) / $rootScope.worldWidth
        svgHeightScale = parseInt(svg.style("height")) / $rootScope.worldHeight
        scaled =
            x: inputX * svgWidthScale
            y: inputY * svgHeightScale
        # console.log 'gameToMapCoords scaled:', inputX, inputY, scaled
        scaled
        

    renderToon = ->
        if !$rootScope.myToon
            throw "myToon not set"
        toon = $rootScope.myToon
        coords = gameToMapCoords(toon.LocX, toon.LocY)

        myLoc = svg.selectAll("#my-location")
            .data([coords])
        console.log 'toon coords:', coords
        
        toonWidth = 15
        myLoc.enter()
            .append("image")
        
        myLoc.attr("id", "my-location")
            .attr("xlink:href", "/static/img/guy.png")
            .attr("width", toonWidth)
            .attr("height", toonWidth)
            .attr("x", (d) -> d.x - toonWidth / 2)
            .attr("y", (d) -> d.y)
        
        myLoc.exit().remove()

    renderDestination = (destX, destY) ->
        yOffset = 10
        width = 10
        height = 10
        destPointData = [
            {x: destX, y: destY - yOffset},
            {x: destX - width / 2, y: destY - yOffset - height},
            {x: destX + width / 2, y: destY - yOffset - height},
            {x: destX, y: destY - yOffset},
        ]

        myDest = svg.selectAll("#my-destination")
            .data(destPointData)

        d3.select("#my-dest-point").remove()
        svg.append("path")
            .attr("id", "my-dest-point")
            .attr("d", lineFunc(destPointData))
            .attr("stroke", "white")
            .attr("stroke-width", 1)
            .attr("fill", "none")
            .attr("opacity", 0)
            .transition()
            .duration(600)
            .attr("opacity", 1)
            .attr("transform", "translate(0, #{yOffset})")

    renderLocations = ->
        allCoords = []

        if !$rootScope.displayedLocations or $rootScope.displayedLocations.length == 0 
            return

        _.each $rootScope.displayedLocations, (loc, idx) ->
            coords = gameToMapCoords(loc.X1 + loc.X2 / 2, loc.Y1 + loc.Y2 / 2)
            coords.id = loc.Id
            allCoords.push coords

        locations = svg.selectAll(".world-location")
            .data($rootScope.displayedLocations)
        
        locations.enter()
            .append("image")
        
        locations.attr("xlink:href", "/static/img/town.png")
            .attr("class", "world-location")
            .attr("width", 15)
            .attr("height", 15)
            .attr("x", (d) -> d.X1 + d.X2 / 2)
            .attr("y", (d) -> d.Y1 + d.Y2 / 2)
        
        locations.exit()
            .remove()

    $scope.mapClick = ($event) ->
        destX = $event.offsetX
        destY = $event.offsetY

        renderDestination(destX, destY)
        console.log 'destX:', destX
        console.log 'destY:', destY

        postData = mapToGameCoords(destX, destY)
        console.log 'postData:', postData
        $http.post("/api/destination", postData).then (response) ->
            if response.data.success
                $rootScope.setMyToon response.data.toon
            else 
                $rootScope.alertUser "Failed to set destination: #{response.data.reason}" 

    if $rootScope.myToon
        renderToon()

    $rootScope.$watch "displayedLocations", ->
        renderLocations()

    $rootScope.$watch "myToon", ->
        if $rootScope.myToon
            renderLocations()
            renderToon()
