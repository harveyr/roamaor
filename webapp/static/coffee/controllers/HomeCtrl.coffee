angular.module(APP_NAME).controller 'HomeCtrl', ($scope, $rootScope, $http, $timeout) ->

    mapColor = "#E7E7E7"
    toonRadius = 5
    scaleX = null
    scaleY = null
    $scope.renderedLocationIds = []

    svgHeight = 500
    svg = d3.select("svg")
        .attr("height", svgHeight)
        # .style("background-color", mapColor)
        # .style("box-shadow", "1px 1px 1px #999")

    map = $(".game-map")
    mapWidth = map.width()
    mapHeight = svgHeight + 56

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

        myLoc = svg.selectAll("#my-toon")
            .data([coords])
        console.log 'toon coords:', coords
        
        translate = "translate(#{coords.x}, #{coords.y})"

        maxHealthBarHeight = 15
        hpPercent = (toon.Hp / toon.MaxHp)
        healthBarHeight = hpPercent * maxHealthBarHeight
        healthBarColor = "#15ff00"
        if hpPercent < .4
            healthBarColor = "red"
        else if hpPercent < 0.6
            healthBarColor = "#ffea00"

        if myLoc.attr("opacity") < 1
            myLoc.selectAll("#my-health-bar")
                .attr("height", healthBarHeight)
                .attr("y", maxHealthBarHeight - healthBarHeight + 1)
                .style("fill", healthBarColor)

            myLoc.attr("transform", translate)
                .transition()
                .delay(500)
                .duration(500)
                .attr("opacity", 1)
        else
            myLoc.transition()
                .delay(500)
                .duration(500)
                .attr("transform", translate)

            myLoc.selectAll("#my-health-bar")
                .transition()
                .attr("height", healthBarHeight)
                .attr("y", maxHealthBarHeight - healthBarHeight + 1)
                .style("fill", healthBarColor)


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
        console.log 'renderLocations'
        allCoords = []

        if !$rootScope.displayedLocations or $rootScope.displayedLocations.length == 0 
            return

        _.each $rootScope.displayedLocations, (loc, idx) ->
            coords = gameToMapCoords(loc.X1 + loc.X2 / 2, loc.Y1 + loc.Y2 / 2)
            coords.id = loc.Id
            allCoords.push coords

        transformFunc = (d) ->
            x = d.X1 + d.X2 / 2
            y = d.Y1 + d.Y2 / 2
            "translate (#{x}, #{y})"

        $timeout ->
            locs = svg.selectAll(".svg-town")
                .data($rootScope.displayedLocations)
                .attr("transform", transformFunc)

            if locs.attr("opacity") < 1
                locs.transition()
                    .delay((d, i) -> i * 200)
                    .duration(1000)
                    .attr("opacity", 1)

                locs.insert("rect", "rect")
                    # .attr("width", (d) -> 20)
                    # .attr("height", (d) -> 20)
                    .attr("width", (d) -> d.X2 - d.X1)
                    .attr("height", (d) -> d.Y2 - d.Y1)
                    .attr("fill", "rgba(6, 212, 0, 0.3)")

            # color = "#555"
            # locs.selectAll("polyline")
            #     .attr("stroke", color)
            # locs.selectAll("rect")
            #     .attr("stroke", color)
            # locs.selectAll("path")
            #     .attr("stroke", color)

        , 0


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


    zoom = d3.behavior.zoom()
        .on "zoom", ->
            svg.selectAll("g")
                .attr("transform")


    if $rootScope.myToon
        renderToon()

    $rootScope.$watch "displayedLocations", ->
        renderLocations()

    $rootScope.$watch "myToon", ->
        if $rootScope.myToon
            renderToon()
