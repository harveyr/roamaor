angular.module(APP_NAME).controller 'HomeCtrl', ($scope, $rootScope, $http, $timeout) ->

    # See ColorBrewer! https://github.com/mbostock/d3/wiki/Ordinal-Scales

    svgHeight = 500
    map = $(".game-map")
    svg = d3.select("svg")
        .attr("height", svgHeight)

    mapWidth = map.width()
    mapHeight = svgHeight + 56
    svgWidth = parseInt(svg.style("width"))
    svgHeight = parseInt(svg.style("height"))
    svgWidthScale = svgWidth / $rootScope.worldWidth
    svgHeightScale = svgHeight / $rootScope.worldHeight

    xScale = d3.scale.linear()
        .domain([0, $rootScope.worldWidth])
        .range([0, 500])
    yScale = d3.scale.linear()
        .domain([0, $rootScope.worldHeight])
        .range([0, 500])

    $scope.zoomScale = 1
    $scope.translate = [0, 0]

    $scope.mapStyle =
        "background-size": "#{mapWidth}px #{mapHeight}px"

    lineFunc = d3.svg.line()
        .x((d) -> d.x)
        .y((d) -> d.y)
        .interpolate('linear')

    myDestPath = svg.append("path")
        .attr("id", "my-dest-path")

    gameToMapCoords = (inputX, inputY) ->

        scaled =
            x: xScale(inputX) * $scope.zoomScale + $scope.translate[0]
            y: yScale(inputY) * $scope.zoomScale + $scope.translate[1]

    mapToGameCoords = (inputX, inputY) ->
        svgWidthScale = parseInt(svg.style("width")) / $rootScope.worldWidth
        svgHeightScale = parseInt(svg.style("height")) / $rootScope.worldHeight
        scaled =
            x: inputX / svgWidthScale
            y: inputY / svgHeightScale

    $scope.applyZoom = _.throttle ->
        svg.selectAll("#my-toon")

    , 300

    renderToon = ->
        if !$rootScope.myToon
            throw "myToon not set"
        toon = $rootScope.myToon
        coords = gameToMapCoords(toon.LocX, toon.LocY)

        myLoc = svg.selectAll("#my-toon")
            .data([coords])

        translate = "translate(#{coords.x}, #{coords.y}) scale(#{$scope.zoomScale})"

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
                .duration(100)
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

    selectLocations = ->
        d3.selectAll(".svg-town")

    locationTransform = (d) ->
        coords = gameToMapCoords(d.X1 + d.X2 / 2, d.Y1 + d.Y2 / 2)
        "translate (#{coords.x}, #{coords.y}) scale(#{$scope.zoomScale})"

    renderLocations = ->

        if !$rootScope.displayedLocations or $rootScope.displayedLocations.length == 0 
            return

        $timeout ->
            locs = selectLocations()
                .data($rootScope.displayedLocations)
                .attr("transform", locationTransform)

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

    applyZoom = _.throttle ->
        renderToon()
        renderLocations()
    , 100

    zoom = d3.behavior.zoom()
        .on "zoom", ->
            $scope.translate = d3.event.translate
            $scope.zoomScale = d3.event.scale
            applyZoom()
    
    zoom.scaleExtent([0.4, 3.0])
    zoom.x(xScale)
    zoom.y(xScale)

    svg.call(zoom)

    if $rootScope.myToon
        renderToon()

    $rootScope.$watch "displayedLocations", ->
        renderLocations()

    $rootScope.$watch "myToon", ->
        if $rootScope.myToon
            renderToon()
