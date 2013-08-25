angular.module(APP_NAME).controller 'HomeCtrl', ($scope, $rootScope, $http, $timeout) ->

    mapColor = "#E7E7E7"
    toonRadius = 5
    scaleX = null
    scaleY = null

    svg = d3.select("svg")
        .attr("height", 500)
        # .style("background-color", mapColor)
        # .style("box-shadow", "1px 1px 1px #999")

    console.log '$(.game-map).height():', 
    console.log '$(.game-map).width():', $(".game-map").width()

    map = $(".game-map")
    mapWidth = map.width()
    mapHeight = map.height()

    $scope.mapStyle =
        "background-size": "#{mapWidth}px #{mapHeight}px"
    console.log '$scope.mapStyle:', $scope.mapStyle

    # svg.append("image")
    #     .attr("xlink:href", "/static/img/mapbg.png")
    #     .attr("width", svgWidth)
    #     .attr("height", svgHeight)


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

    renderToon = ->
        if !$rootScope.myToon
            throw "myToon not set"
        svgHeight = svg.attr("height")
        toon = $rootScope.myToon
        coords = gameToMapCoords(toon.LocX, toon.LocY)
        coords.x = Math.max(coords.x, toon.LocX + toonRadius / 2 + 1)
        coords.y = Math.max(coords.y, toon.LocY + toonRadius / 2)

        elemId = "toon-#{toon.Id}"
        d3.select(elemId).remove()
        svg.append("circle")
            .attr("id", elemId)
            .attr("cx", coords.x)
            .attr("cy", coords.y)
            .attr("r", toonRadius)
            .style("fill", "#ccc")

    renderLocations = ->
        _.each $rootScope.displayedLocations, (loc, idx) ->
            console.log 'loc:', loc
            coords = gameToMapCoords(loc.X1 + loc.X2 / 2, loc.Y1 + loc.Y2 / 2)
            console.log 'coords:', coords
            svg.append("circle")
                .attr("class", "world-location")
                .attr("cx", coords.x)
                .attr("cy", coords.y)
                .attr("r", 5)
                .attr("stroke-width", 1)
                .attr("stroke", "blue")

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


    $scope.mapClick = ($event) ->
        destX = $event.offsetX
        destY = $event.offsetY

        renderDestination(destX, destY)

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
            renderToon()
