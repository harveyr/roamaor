angular.module(APP_NAME).controller 'HomeCtrl', ($scope, $rootScope, $http, $timeout) ->

    mapColor = "#E7E7E7"
    toonRadius = 5
    scaleX = null
    scaleY = null

    svg = d3.select("svg")
        .attr("height", 500)
        .style("background-color", mapColor)
        .style("box-shadow", "1px 1px 1px #999")

    lineFunc = d3.svg.line()
        .x((d) -> d.x)
        .y((d) -> d.y)
        .interpolate('linear')

    triFunc = d3.svg.symbol()
        .type("triangle-up")

    myDest = svg.append("circle")
        .attr("id", "my-dest")
        .attr("cx", 0)
        .attr("cy", 0)
        .attr("r", 5)
        .style("fill", mapColor)

    myDestPath = svg.append("path")
        .attr("id", "my-dest-path")

    toonSvgCoords = (toon) ->
        {
            x: toon.locx + toonRadius
            y: svg.attr("height") - toon.locy - toonRadius
        }

    svg.append("image")
        .attr("xlink:href", "/static/img/pin.png")
        .attr("width", 20)
        .attr("height", 20)

    $scope.mapClick = ($event) ->
        toon = $rootScope.allToons[0]
        toonCoords = toonSvgCoords(toon)
        destX = $event.offsetX
        destY = $event.offsetY

        lineData = [
            {x: toonCoords.x, y: toonCoords.y},
            {x: destX, y: destY},
        ]
        # myDestPath.attr("d", lineFunc(lineData))
        #     .attr("stroke", "#cfcfcf")
        #     .attr("stroke-width", 1)
        #     .attr("fill", "none")

        yOffset = 10
        width = 10
        height = 10
        destPointData = [
            {x: destX, y: destY - yOffset},
            {x: destX - width / 2, y: destY - yOffset - height},
            {x: destX + width / 2, y: destY - yOffset - height},
            {x: destX, y: destY - yOffset},
        ]
        console.log 'destPointData:', destPointData

        d3.select("#my-dest-point").remove()
        svg.append("path")
            .attr("id", "my-dest-point")
            .attr("d", lineFunc(destPointData))
            .attr("stroke", "#555")
            .attr("stroke-width", 1)
            .attr("fill", "none")
            .attr("opacity", 0)
            .transition()
            .duration(600)
            .attr("opacity", 1)
            .attr("transform", "translate(0, #{yOffset})")


        svgHeightScale = parseInt(svg.style("height")) / $rootScope.worldHeight
        svgWidthScale = parseInt(svg.style("width")) / $rootScope.worldWidth
        scaledDestX = destX / svgWidthScale
        scaledDestY = destY / svgHeightScale
        postData =
            x: scaledDestX
            y: scaledDestY

        $http.post("/api/destination", postData).then (response) ->
            console.log 'response:', response


    renderToons = ->
        console.log 'rendering:', $rootScope.allToons
        svgHeight = svg.attr("height")
        _.each $rootScope.allToons, (toon, idx) ->
            coords = toonSvgCoords(toon)
            toonLoc = svg.append("circle")
                .attr("id", "toon-#{toon._id}")
                .attr("cx", coords.x)
                .attr("cy", coords.y)
                .attr("r", toonRadius)
                .style("fill", "#E7E7E7")
            toonLoc.transition()
                .delay(500 + 20 * idx)
                .style("fill", "#777")

    fetchData = ->
        $http.get("/api/admin/alltoons").then (response) ->
            $rootScope.allToons = response.data
            renderToons()
        $http.get("/api/bootstrap").then (response) ->
            $rootScope.worldHeight = response.data.worldHeight
            $rootScope.worldWidth = response.data.worldWidth
    fetchData()
