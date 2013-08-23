angular.module(APP_NAME).controller 'HomeCtrl', ($scope, $rootScope) ->

    $scope.markerStyle =
        top: '20px'
        left: '20px'

    $scope.locationStyle =
        top: '50px'
        left: '50px'

    $scope.mapClick = ($event) ->
        console.log('here', $event);
        $scope.markerStyle.top = $event.y + 'px'
        $scope.markerStyle.left = $event.x + 'px'


    svg = d3.select("svg")
        .attr("height", 500)
        .style("background-color", "#E7E7E7")

    myLoc = svg.append("circle")
        .attr("id", "my-location")
        .attr("cx", 25)
        .attr("cy", 25)
        .attr("r", 5)
        .style("fill", "#E7E7E7")

    myLoc.transition()
        .delay(500)
        .style("fill", "#777")
