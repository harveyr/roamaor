// Generated by CoffeeScript 1.6.3
(function() {
  var APP_NAME, DIRECTIVE_MODULE, app;

  APP_NAME = 'roamaor';

  DIRECTIVE_MODULE = "" + APP_NAME + ".directives";

  angular.module(DIRECTIVE_MODULE, []);

  angular.module(DIRECTIVE_MODULE).directive('topNavbar', function($location) {
    var directive;
    return directive = {
      replace: true,
      template: "<div class=\"navbar\">\n    <div class=\"navbar-inner\">\n        <a class=\"brand\" href=\"#\">{{appName}}</a>\n        <ul class=\"nav\">\n            <li ng-repeat=\"link in navLinks\"\n                ng-class=\"{active: currentPath == link.href}\">\n                    <a href=\"{{link.href}}\">{{link.title}}</a>\n            </li>\n        </ul>\n    </div>\n</div>",
      link: function(scope) {
        scope.navLinks = [
          {
            href: '/',
            title: 'Home'
          }
        ];
        return scope.$on("$routeChangeSuccess", function(e, current, previous) {
          return scope.currentPath = $location.path();
        });
      }
    };
  });

  angular.module(DIRECTIVE_MODULE).directive('userFeedback', function() {
    var directive;
    return directive = {
      replace: true,
      template: "<div class=\"row\" ng-show=\"userAlert\">\n    <div class=\"small-12\">\n        <div data-alert class=\"alert-box\">\n            {{userAlert}}\n            <a href=\"#\" class=\"close\">&times;</a>\n        </div>\n    </div>\n</div>",
      link: function(scope) {}
    };
  });

  app = angular.module(APP_NAME, ["" + APP_NAME + ".directives"]).run(function($rootScope, $http) {
    $rootScope.appName = "Roamoar";
    $rootScope.myToon = null;
    $rootScope.alertUser = function(string) {
      return $rootScope.userAlert = string;
    };
    $rootScope.setMyToon = function(toon) {
      if (!toon) {
        throw "setMyToon received null toon";
      }
      return $rootScope.myToon = toon;
    };
    $rootScope.fetchBundle = function() {
      return $http.get("/api/bootstrap").then(function(response) {
        if (!response.data.success) {
          $rootScope.alertUser("Failed to fetch bootstrap bundle. (" + response.data.reason + ")");
          return;
        }
        console.log('bundle:', response.data);
        $rootScope.worldHeight = response.data.worldHeight;
        $rootScope.worldWidth = response.data.worldWidth;
        $rootScope.myUser = response.data.user;
        $rootScope.toonLogs = response.data.toonLogs;
        $rootScope.logTypes = response.data.logTypes;
        if (response.data.toon) {
          $rootScope.setMyToon(response.data.toon);
          return $rootScope.displayedLocations = response.data.visited;
        }
      });
    };
    return $rootScope.fetchBundle();
  });

  angular.module(APP_NAME).controller('HomeCtrl', function($scope, $rootScope, $http, $timeout) {
    var gameToMapCoords, lineFunc, map, mapColor, mapHeight, mapToGameCoords, mapWidth, myDestPath, renderDestination, renderLocations, renderToon, scaleX, scaleY, svg, toonRadius, toonSvgCoords;
    mapColor = "#E7E7E7";
    toonRadius = 5;
    scaleX = null;
    scaleY = null;
    svg = d3.select("svg").attr("height", 500);
    map = $(".game-map");
    mapWidth = map.width();
    mapHeight = map.height() - 9;
    $scope.mapStyle = {
      "background-size": "" + mapWidth + "px " + mapHeight + "px"
    };
    lineFunc = d3.svg.line().x(function(d) {
      return d.x;
    }).y(function(d) {
      return d.y;
    }).interpolate('linear');
    myDestPath = svg.append("path").attr("id", "my-dest-path");
    toonSvgCoords = function(toon) {
      return {
        x: toon.LocX + toonRadius,
        y: svg.attr("height") - toon.LocY - toonRadius
      };
    };
    mapToGameCoords = function(inputX, inputY) {
      var scaled, svgHeightScale, svgWidthScale;
      svgWidthScale = parseInt(svg.style("width")) / $rootScope.worldWidth;
      svgHeightScale = parseInt(svg.style("height")) / $rootScope.worldHeight;
      return scaled = {
        x: inputX / svgWidthScale,
        y: inputY / svgHeightScale
      };
    };
    gameToMapCoords = function(inputX, inputY) {
      var scaled, svgHeightScale, svgWidthScale;
      svgWidthScale = parseInt(svg.style("width")) / $rootScope.worldWidth;
      svgHeightScale = parseInt(svg.style("height")) / $rootScope.worldHeight;
      scaled = {
        x: inputX * svgWidthScale,
        y: inputY * svgHeightScale
      };
      return scaled;
    };
    renderToon = function() {
      var coords, myLoc, toon, toonWidth;
      if (!$rootScope.myToon) {
        throw "myToon not set";
      }
      toon = $rootScope.myToon;
      coords = gameToMapCoords(toon.LocX, toon.LocY);
      myLoc = svg.selectAll("#my-location").data([coords]);
      console.log('toon coords:', coords);
      toonWidth = 15;
      myLoc.enter().append("image");
      myLoc.attr("id", "my-location").attr("xlink:href", "/static/img/guy.png").attr("width", toonWidth).attr("height", toonWidth).attr("x", function(d) {
        return d.x - toonWidth / 2;
      }).attr("y", function(d) {
        return d.y;
      });
      return myLoc.exit().remove();
    };
    renderDestination = function(destX, destY) {
      var destPointData, height, myDest, width, yOffset;
      yOffset = 10;
      width = 10;
      height = 10;
      destPointData = [
        {
          x: destX,
          y: destY - yOffset
        }, {
          x: destX - width / 2,
          y: destY - yOffset - height
        }, {
          x: destX + width / 2,
          y: destY - yOffset - height
        }, {
          x: destX,
          y: destY - yOffset
        }
      ];
      myDest = svg.selectAll("#my-destination").data(destPointData);
      d3.select("#my-dest-point").remove();
      return svg.append("path").attr("id", "my-dest-point").attr("d", lineFunc(destPointData)).attr("stroke", "white").attr("stroke-width", 1).attr("fill", "none").attr("opacity", 0).transition().duration(600).attr("opacity", 1).attr("transform", "translate(0, " + yOffset + ")");
    };
    renderLocations = function() {
      var allCoords, locations;
      allCoords = [];
      if (!$rootScope.displayedLocations || $rootScope.displayedLocations.length === 0) {
        return;
      }
      _.each($rootScope.displayedLocations, function(loc, idx) {
        var coords;
        coords = gameToMapCoords(loc.X1 + loc.X2 / 2, loc.Y1 + loc.Y2 / 2);
        coords.id = loc.Id;
        return allCoords.push(coords);
      });
      locations = svg.selectAll(".world-location").data($rootScope.displayedLocations);
      locations.enter().append("image");
      locations.attr("xlink:href", "/static/img/town.png").attr("class", "world-location").attr("width", 15).attr("height", 15).attr("x", function(d) {
        return d.X1 + d.X2 / 2;
      }).attr("y", function(d) {
        return d.Y1 + d.Y2 / 2;
      });
      return locations.exit().remove();
    };
    $scope.mapClick = function($event) {
      var destX, destY, postData;
      destX = $event.offsetX;
      destY = $event.offsetY;
      renderDestination(destX, destY);
      console.log('destX:', destX);
      console.log('destY:', destY);
      postData = mapToGameCoords(destX, destY);
      console.log('postData:', postData);
      return $http.post("/api/destination", postData).then(function(response) {
        if (response.data.success) {
          return $rootScope.setMyToon(response.data.toon);
        } else {
          return $rootScope.alertUser("Failed to set destination: " + response.data.reason);
        }
      });
    };
    if ($rootScope.myToon) {
      renderToon();
    }
    $rootScope.$watch("displayedLocations", function() {
      return renderLocations();
    });
    return $rootScope.$watch("myToon", function() {
      if ($rootScope.myToon) {
        renderLocations();
        return renderToon();
      }
    });
  });

  angular.module(APP_NAME).config([
    '$routeProvider', '$locationProvider', function($routeProvider, $locationProvider) {
      var url, urlBase;
      urlBase = '/app';
      url = function(url) {
        return "" + urlBase + "/url";
      };
      $routeProvider.when(urlBase, {
        controller: 'HomeCtrl',
        templateUrl: 'static/partials/home.html'
      });
      return $locationProvider.html5Mode(true).hashPrefix('!');
    }
  ]);

  angular.module(APP_NAME).controller('AdminCtrl', function($scope, $rootScope, $http, $timeout) {
    var updateData;
    $scope.admin = {};
    $scope.autoUpdate = false;
    $scope.submitNewPlayer = function(name) {
      var data;
      console.log('name:', name);
      data = {
        name: name
      };
      return $http.post("/api/admin/newtoon", data).then(function(response) {
        return console.log('response:', response);
      });
    };
    $scope.selectedToonChange = function(toonId) {
      var data;
      data = {
        toonId: toonId
      };
      $rootScope.myToon = _.findWhere($rootScope.allToons, {
        "_id": toonId
      });
      console.log('$rootScope.myToon:', $rootScope.myToon);
      return $http.post("/api/activetoon", data).then(function(response) {
        return console.log('response:', response.status, response.data);
      });
    };
    $scope.showAllLocs = function() {
      return $http.get("/api/admin/alllocations").then(function(response) {
        return $rootScope.displayedLocations = response.data;
      });
    };
    updateData = function() {
      if (!$scope.autoUpdate) {
        return;
      }
      $rootScope.fetchBundle();
      return $timeout(function() {
        return updateData();
      }, 2000);
    };
    $scope.toggleAutoUpdate = function(auto) {
      $scope.autoUpdate = !$scope.autoUpdate;
      return updateData();
    };
    return $http.get("/api/admin/alltoons").then(function(response) {
      $rootScope.allToons = response.data;
      console.log('$rootScope.allToons:', $rootScope.allToons);
      if ($rootScope.myToon) {
        return $scope.admin.selectedToon = $rootScope.myToon.Id;
      }
    });
  });

  angular.module(DIRECTIVE_MODULE).directive("toonSummary", function($rootScope) {
    var directive;
    return directive = {
      replace: true,
      scope: true,
      template: "<div class=\"row\">\n    <div class=\"small-12\">\n        <p>\n            <strong>{{name}}</strong>\n        </p>\n        <p>\n            Level {{level}}\n        </p>\n        <p>\n            Hp: {{hp}} / {{maxHp}}\n            <div class=\"progress\"><span class=\"meter\" style=\"width: {{hpPercentage}}%\"></span></div>\n        </p>\n        <p>\n            Location: {{locX}}, {{locY}}\n        </p>\n        <p>\n            Destination: {{destX}}, {{destY}}\n        </p>\n        <p>\n            Fights Won: {{fightsWon}} / {{fights}}\n        </p>\n        <p>\n            Locations Visited: {{myToon.LocationsVisited}}\n        </p>\n    </div>\n</div>",
      link: function(scope) {
        var applyToon;
        applyToon = function(toon) {
          scope.name = toon.Name;
          scope.level = toon.Level;
          scope.locX = toon.LocX.toFixed(2);
          scope.locY = toon.LocY.toFixed(2);
          scope.hp = toon.Hp;
          scope.maxHp = toon.MaxHp;
          scope.hpPercentage = toon.Hp / toon.MaxHp * 100;
          scope.fights = toon.Fights;
          scope.fightsWon = toon.FightsWon;
          scope.destX = toon.DestX.toFixed(2);
          return scope.destY = toon.DestY.toFixed(2);
        };
        if ($rootScope.myToon) {
          applyToon($rootScope.myToon);
        }
        console.log('[toonSummary] $rootScope.myToon:', $rootScope.myToon);
        return $rootScope.$watch('myToon', function() {
          if ($rootScope.myToon) {
            return applyToon($rootScope.myToon);
          }
        });
      }
    };
  });

  angular.module(DIRECTIVE_MODULE).directive("toonLog", function($rootScope) {
    var directive;
    return directive = {
      replace: true,
      scope: {
        item: '='
      },
      template: "<div class=\"row log-item\">\n    <div class=\"small-12 columns\">\n        {{name}}\n        {{action}}\n    </div>\n</div>",
      link: function(scope) {
        scope.name = $rootScope.myToon.Name;
        switch (scope.item.LogType) {
          case $rootScope.logTypes.fight:
            return scope.action = "man-danced with a Level " + scope.item.Data.opponentLevel + " " + scope.item.Data.opponentName;
          case $rootScope.logTypes.locationDiscovery:
            return scope.action = "discovered the location of " + scope.item.Data.locationName + "!";
        }
      }
    };
  });

  angular.module(APP_NAME).controller('ToonLogCtrl', function($scope, $rootScope, $http, $timeout) {});

}).call(this);
