angular.module(APP_NAME).controller 'HomeCtrl', ($scope, $rootScope, $http, $timeout) ->

    # See ColorBrewer! https://github.com/mbostock/d3/wiki/Ordinal-Scales

    lastZoom = new Date()

    healthColors = ["#FF000E", "#FF2D0B", "#FF5B08", "#FF8905", "#FFB702", "#D0EA00", "#A2EF00", "#73F400", "#45F900", "#17FF00"]

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

        toonSvg = svg.selectAll("#my-toon")

        transform = "translate(#{coords.x}, #{coords.y}) scale(#{$scope.zoomScale})"

        maxHealthBarHeight = 15
        hpPercent = (toon.Hp / toon.MaxHp)
        healthBarHeight = hpPercent * maxHealthBarHeight
        healthBarColor = "#15ff00"
        if hpPercent < .4
            healthBarColor = "red"
        else if hpPercent < 0.6
            healthBarColor = "#ffea00"

        if hpPercent < 0.15
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
            toonSvg.selectAll(".healthy-toon-toon")
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
        d3.selectAll(".svg-town")

    locationTransform = (d) ->
        coords = gameToMapCoords(d.X1 + d.X2 / 2, d.Y1 + d.Y2 / 2)
        "translate (#{coords.x}, #{coords.y}) scale(#{$scope.zoomScale})"

    drawLocations = ->
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
        console.log 'draw grid'
        gridX = svg.selectAll(".grid-line-x")
            .data(gridLinesX)

        gridX.enter()
            .insert("svg:line")
            .attr("class", "grid-line-x")

        gridX.attr("x1", (d) -> xScale(d))
            .attr("y1", yScale(0))
            .attr("x2", (d) -> xScale(d))
            .attr("y2", yScale($rootScope.worldHeight))
            .style("stroke", "#555")
            .style("stroke-width", 1)

        gridX.exit().remove()

        gridY = svg.selectAll(".grid-line-y")
            .data(gridLinesY)

        gridY.enter()
            .insert("svg:line")
            .attr("class", "grid-line-y")

        gridY.attr("x1", xScale(0))
            .attr("y1", (d) -> yScale(d))
            .attr("x2", xScale($rootScope.worldWidth))
            .attr("y2", (d) -> yScale(d))
            .style("stroke", "#555")
            .style("stroke-width", 1)

        gridY.exit().remove()

    updateView = _.throttle ->
        drawToon()
        drawLocations()
        drawDestination(false)
        drawGrid()
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
