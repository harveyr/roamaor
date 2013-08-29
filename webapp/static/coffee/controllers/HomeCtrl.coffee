angular.module(APP_NAME).controller 'HomeCtrl', ($scope, $rootScope, $http, $timeout, GameConstants) ->

    # See ColorBrewer! https://github.com/mbostock/d3/wiki/Ordinal-Scales

    lastZoom = new Date()

    redGreenGradient = ["#FF000E", "#FF2D0B", "#FF5B08", "#FF8905", "#FFB702", "#D0EA00", "#A2EF00", "#73F400", "#45F900", "#17FF00"]

    svgHeight = 500
    map = $(".game-map")
    svg = d3.select("svg")
        .attr("height", svgHeight)

    mapWidth = map.width()
    mapHeight = svgHeight + 60
    svgWidth = parseInt(svg.style("width"))
    svgHeight = parseInt(svg.style("height"))
    svgWidthScale = svgWidth / $rootScope.worldWidth
    svgHeightScale = svgHeight / $rootScope.worldHeight
    walkPromise = null

    gridDiameter = 200
    gridLinesX = d3.range(0, $rootScope.worldWidth, gridDiameter)
    gridLinesY = d3.range(0, $rootScope.worldHeight, gridDiameter)


    xScale = d3.scale.linear()
        # .domain([0, $rootScope.worldWidth])
        # .range([0, 500])
    yScale = d3.scale.linear()
        # .domain([0, $rootScope.worldHeight])
        # .range([0, 500])

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
            x: xScale(inputX)
            y: yScale(inputY)

    mapToGameCoords = (inputX, inputY) ->
        coords =
            x: xScale.invert(inputX)
            y: yScale.invert(inputY)

    selectToonSvg = ->
        svg.selectAll("#my-toon")

    walkToon = (step = 1) ->
        $timeout.cancel walkPromise
        toon = $rootScope.myToon

        if toon.Hp / toon.MaxHp < 0.15
            return

        legs = selectToonSvg().select(".toon-legs")

        if toon.LocX == toon.DestX and toon.LocY == toon.DestY
            legs.attr("transform", null)
            return

        delay = 300

        if step == 1
            data = [20, 0]
        else
            data = [0, -20]

        legLines = legs.selectAll("line")
            .data(data)
            .attr "transform", (d) ->
                if d == 0
                    return ""
                else
                    return "rotate(#{d}, 10, 20)"

            # console.log 'legLines:', legLines
            # legLines.attr "transform", (d) ->
            #     angle = d * 20
            #     "rotate(#{angle}, 10, 20)"

        walkPromise = $timeout ->
            walkToon(step * -1)
        , delay


    drawHealthBar = (anim = true) ->
        toon = $rootScope.myToon
        maxHealthBarHeight = 15
        hpPercent = (toon.Hp / toon.MaxHp)
        healthBarHeight = Math.max(1, hpPercent * maxHealthBarHeight)
        healthBarColor = "#15ff00"
        if hpPercent < .4
            healthBarColor = "red"
        else if hpPercent < 0.6
            healthBarColor = "#ffea00"

        healthBar = selectToonSvg().selectAll(".toon-health-bar")
        if anim
            healthBar.transition()
            .attr("height", healthBarHeight)
            .attr("y", maxHealthBarHeight - healthBarHeight + 1)
            .style("fill", healthBarColor)
        else
            healthBar.attr("height", healthBarHeight)
            .attr("y", maxHealthBarHeight - healthBarHeight + 1)
            .style("fill", healthBarColor)


    drawToon = ->
        if !$rootScope.myToon
            throw "myToon not set"
        toon = $rootScope.myToon
        coords = gameToMapCoords(toon.LocX, toon.LocY)

        toonSvg = svg.select("#my-toon")

        transform = "translate(#{coords.x}, #{coords.y}) scale(#{$scope.zoomScale})"

        maxHealthBarHeight = 15
        hpPercent = (toon.Hp / toon.MaxHp)
        healthBarHeight = hpPercent * maxHealthBarHeight
        healthBarColor = "#15ff00"
        if hpPercent < .4
            healthBarColor = "red"
        else if hpPercent < 0.6
            healthBarColor = "#ffea00"

        if hpPercent < 0.1
            toonSvg.selectAll(".healthy-toon")
                .attr("opacity", "0")
            deadToon = toonSvg.selectAll(".dead-toon")
                .attr("opacity", "1")
            deadToon.selectAll("polygon")
                .style("fill", "#ff7e86")
            deadToon.selectAll("polyline")
                .style("fill", "#ff7e86")
            toonSvg.attr("transform", transform)
            toonSvg.attr("opacity", 1)
            drawHealthBar(true)
            return
        else
            toonSvg.selectAll(".dead-toon")
                .attr("opacity", 0)
            toonSvg.selectAll(".healthy-toon")
                .attr("opacity", 1)

        if toonSvg.attr("opacity") < 1
            drawHealthBar(false)
            toonSvg.attr("transform", transform)
                .transition()
                .delay(500)
                .duration(500)
                .attr("opacity", 1)
        else
            toonSvg.transition()
                .duration(100)
                .attr("transform", transform)
            drawHealthBar(true)
        walkToon()

    drawDestination = (animate = false) ->
        coords = gameToMapCoords($rootScope.myToon.DestX, $rootScope.myToon.DestY)
        width = 18
        height = 21
        yAnimOffset = 10
        targetX = coords.x - width / 2 * $scope.zoomScale
        targetY = coords.y - height * $scope.zoomScale
        startY = targetY - yAnimOffset

        myDest = svg.selectAll("#my-destination")
        startTrans = "translate(#{targetX}, #{startY}) scale(#{$scope.zoomScale})"
        endTrans = "translate(#{targetX}, #{targetY}) scale(#{$scope.zoomScale})"

        if !animate
            myDest.attr("transform", endTrans)
        else
            myDest.attr("transform", startTrans)
            .transition()
            .duration(500)
            .attr("transform", endTrans)


    selectLocations = ->
        d3.selectAll(".svg-location")

    drawLocations = ->
        if !$rootScope.displayedLocations or $rootScope.displayedLocations.length == 0 
            return

        types = GameConstants.get("locationTypes")

        sortedLocations = _.sortBy $rootScope.displayedLocations, (loc) ->
            -1 * ((loc.X2 - loc.X1) + (loc.Y2 - loc.Y1))

        $timeout ->
            locs = selectLocations()
                .data(sortedLocations)
                .attr("transform", (d) ->
                    coords = gameToMapCoords(d.X1, d.Y1)
                    "translate (#{coords.x}, #{coords.y}) scale(#{$scope.zoomScale})"
                )

            if locs.attr("opacity") < 1
                locs.transition()
                    .delay((d, i) -> i * 100)
                    .duration(300)
                    .attr("opacity", 1)

                locs.select(".location-bounds")
                    .attr("width", (d) -> Math.max(0, (d.X2 - d.X1)))
                    .attr("height", (d) -> Math.max(0, (d.Y2 - d.Y1)))
                    .attr("opacity", 0.3)
                    .attr("stroke", "#000")
                    .style("fill", (d) ->
                        return redGreenGradient[10 - Math.floor(d.Danger * 10)])
            locs.select(".location-town")
                .attr("opacity", (d) ->
                    if d.LocationType == types.town
                        return 1
                    0
                )
        , 0

    setDestination = (offsetX, offsetY) ->
        gameCoords = mapToGameCoords(offsetX, offsetY)
        if gameCoords.x < 0 || gameCoords.y < 0
            return
        # drawDestination(scaledMapCoords.x, scaledMapCoords.y, true)
        $rootScope.myToon.DestX = gameCoords.x
        $rootScope.myToon.DestY = gameCoords.y
        drawDestination(true)
        $http.post("/api/destination", gameCoords).then (response) ->
            if response.data.success
                $rootScope.setMyToon response.data.toon
            else 
                $rootScope.alertUser "Failed to set destination: #{response.data.reason}" 
        walkToon()

    drawGrid = ->
        grid = svg.select("#gridlines")
        gridX = grid.selectAll(".grid-line-x")
            .data(gridLinesX)

        gridX.enter()
            .append("svg:line")
            .attr("class", "grid-line-x")

        strokeFunc = (d) ->
            if d == 0
                return "#333"
            else
                return "#999"

        strokeWidthFunc = (d) ->
            1

        gridX.attr("x1", (d) -> xScale(d))
            .attr("y1", yScale(0))
            .attr("x2", (d) -> xScale(d))
            .attr("y2", yScale($rootScope.worldHeight))
            .style("stroke", strokeFunc)
            .style("stroke-width", strokeWidthFunc)

        gridX.exit().remove()

        gridY = grid.selectAll(".grid-line-y")
            .data(gridLinesY)

        gridY.enter()
            .append("svg:line")
            .attr("class", "grid-line-y")

        gridY.attr("x1", xScale(0))
            .attr("y1", (d) -> yScale(d))
            .attr("x2", xScale($rootScope.worldWidth))
            .attr("y2", (d) -> yScale(d))
            .style("stroke", strokeFunc)
            .style("stroke-width", strokeWidthFunc)

        gridY.exit().remove()

    updateView = _.throttle ->
        drawGrid()
        drawToon()
        drawLocations()
        drawDestination(false)
    , 100

    zoom = d3.behavior.zoom()
        .on "zoom", ->
            $scope.translate = d3.event.translate
            $scope.zoomScale = d3.event.scale
            lastZoom = new Date()
            updateView()
    
    zoom.scaleExtent([0.4, 3.0])
    zoom.x(xScale)
    zoom.y(yScale)
    zoom.size([svgWidth, svgHeight])
    zoom(svg)

    svg.on "click", ->
        now = new Date()
        if now.getTime() - lastZoom.getTime() > 1000
            setDestination(d3.event.offsetX, d3.event.offsetY)

    $scope.toonZoom = ->
        console.log 'here'
        toon = $rootScope.myToon
        coords = gameToMapCoords(toon.LocX, toon.LocY)
        zoom.center([coords.x, coords.y])
        zoom.scale(2)
        zoom.event(svg)

    if $rootScope.myToon
        drawToon()
        drawDestination()

    $rootScope.$watch "myToon", ->
        if $rootScope.myToon
            updateView()
