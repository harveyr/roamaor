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
    var updateDisplayedLocations;
    $rootScope.appName = "Roamoar";
    $rootScope.myToon = null;
    $rootScope.displayedLocations = [];
    $rootScope.alertUser = function(string) {
      return $rootScope.userAlert = string;
    };
    $rootScope.setMyToon = function(toon) {
      if (!toon) {
        throw "setMyToon received null toon";
      }
      return $rootScope.myToon = toon;
    };
    updateDisplayedLocations = function(locations) {
      var notDisplayed;
      notDisplayed = _.difference(locations, $rootScope.displayedLocations);
      return $rootScope.displayedLocations = $rootScope.displayedLocations.concat(notDisplayed);
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
          return updateDisplayedLocations(response.data.visited);
        }
      });
    };
    return $rootScope.fetchBundle();
  });

  angular.module(APP_NAME).controller('HomeCtrl', function($scope, $rootScope, $http, $timeout) {
    var applyZoom, gameToMapCoords, initialXRange, initialYRange, lineFunc, locationTransform, map, mapHeight, mapToGameCoords, mapWidth, myDestPath, renderDestination, renderLocations, renderToon, selectLocations, svg, svgHeight, svgHeightScale, svgWidth, svgWidthScale, toonRadius, updateZoomScale, xScale, yScale, zoom, zoomCenterX, zoomCenterY;
    svgHeight = 500;
    map = $(".game-map");
    svg = d3.select("svg").attr("height", svgHeight);
    mapWidth = map.width();
    mapHeight = svgHeight + 56;
    svgWidth = parseInt(svg.style("width"));
    svgHeight = parseInt(svg.style("height"));
    svgWidthScale = svgWidth / $rootScope.worldWidth;
    svgHeightScale = svgHeight / $rootScope.worldHeight;
    zoomCenterX = svgWidth / 2;
    zoomCenterY = svgHeight / 2;
    initialXRange = svgWidth;
    initialYRange = svgHeight;
    $scope.zoomScale = 1;
    $scope.translate = [0, 0];
    toonRadius = 5;
    xScale = d3.scale.linear().domain([0, $rootScope.worldWidth]).range([0, initialXRange]);
    yScale = d3.scale.linear().domain([0, $rootScope.worldHeight]).range([0, initialYRange]);
    $scope.mapStyle = {
      "background-size": "" + mapWidth + "px " + mapHeight + "px"
    };
    lineFunc = d3.svg.line().x(function(d) {
      return d.x;
    }).y(function(d) {
      return d.y;
    }).interpolate('linear');
    myDestPath = svg.append("path").attr("id", "my-dest-path");
    gameToMapCoords = function(inputX, inputY) {
      var scaled;
      return scaled = {
        x: inputX * svgWidthScale + $scope.translate[0],
        y: inputY * svgHeightScale + $scope.translate[1]
      };
    };
    mapToGameCoords = function(inputX, inputY) {
      var scaled;
      svgWidthScale = parseInt(svg.style("width")) / $rootScope.worldWidth;
      svgHeightScale = parseInt(svg.style("height")) / $rootScope.worldHeight;
      return scaled = {
        x: inputX / svgWidthScale,
        y: inputY / svgHeightScale
      };
    };
    $scope.applyZoom = _.throttle(function() {
      return svg.selectAll("#my-toon");
    }, 300);
    renderToon = function() {
      var coords, healthBarColor, healthBarHeight, hpPercent, maxHealthBarHeight, myLoc, toon, translate;
      if (!$rootScope.myToon) {
        throw "myToon not set";
      }
      toon = $rootScope.myToon;
      coords = gameToMapCoords(toon.LocX, toon.LocY);
      myLoc = svg.selectAll("#my-toon").data([coords]);
      translate = "translate(" + coords.x + ", " + coords.y + ") scale(" + $scope.zoomScale + ")";
      maxHealthBarHeight = 15;
      hpPercent = toon.Hp / toon.MaxHp;
      healthBarHeight = hpPercent * maxHealthBarHeight;
      healthBarColor = "#15ff00";
      if (hpPercent < .4) {
        healthBarColor = "red";
      } else if (hpPercent < 0.6) {
        healthBarColor = "#ffea00";
      }
      if (myLoc.attr("opacity") < 1) {
        myLoc.selectAll("#my-health-bar").attr("height", healthBarHeight).attr("y", maxHealthBarHeight - healthBarHeight + 1).style("fill", healthBarColor);
        return myLoc.attr("transform", translate).transition().delay(500).duration(500).attr("opacity", 1);
      } else {
        myLoc.transition().duration(100).attr("transform", translate);
        return myLoc.selectAll("#my-health-bar").transition().attr("height", healthBarHeight).attr("y", maxHealthBarHeight - healthBarHeight + 1).style("fill", healthBarColor);
      }
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
    selectLocations = function() {
      return d3.selectAll(".svg-town");
    };
    locationTransform = function(d) {
      var coords;
      coords = gameToMapCoords(d.X1 + d.X2 / 2, d.Y1 + d.Y2 / 2);
      return "translate (" + coords.x + ", " + coords.y + ") scale(" + $scope.zoomScale + ")";
    };
    renderLocations = function() {
      if (!$rootScope.displayedLocations || $rootScope.displayedLocations.length === 0) {
        return;
      }
      return $timeout(function() {
        var locs;
        locs = selectLocations().data($rootScope.displayedLocations).attr("transform", locationTransform);
        if (locs.attr("opacity") < 1) {
          locs.transition().delay(function(d, i) {
            return i * 200;
          }).duration(1000).attr("opacity", 1);
          return locs.insert("rect", "rect").attr("width", function(d) {
            return d.X2 - d.X1;
          }).attr("height", function(d) {
            return d.Y2 - d.Y1;
          }).attr("fill", "rgba(6, 212, 0, 0.3)");
        }
      }, 0);
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
    applyZoom = _.throttle(function() {
      renderToon();
      return renderLocations();
    }, 300);
    updateZoomScale = function() {
      var wheelDelta, zoomMod;
      console.log('d3.event:', d3.event);
      console.log('d3.event.translate:', d3.event.translate);
      if (d3.event.sourceEvent.type === "mousemove") {
        wheelDelta = d3.event.sourceEvent.wheelDelta;
        if (wheelDelta > 0) {
          if ($scope.zoomScale > 3) {
            return;
          }
          zoomMod = 0.1;
        }
        if (wheelDelta < 0) {
          if ($scope.zoomScale < 0.4) {
            return;
          }
          zoomMod = -0.1;
        }
        $scope.zoomScale = $scope.zoomScale + zoomMod;
        $scope.translate = d3.event.translate;
      }
      return applyZoom();
    };
    zoom = d3.behavior.zoom().on("zoom", function() {
      return updateZoomScale();
    });
    zoom.scaleExtent([0.4, 3.0]);
    svg.call(zoom);
    if ($rootScope.myToon) {
      renderToon();
    }
    $rootScope.$watch("displayedLocations", function() {
      return renderLocations();
    });
    return $rootScope.$watch("myToon", function() {
      if ($rootScope.myToon) {
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
      template: "<div class=\"row\">\n    <div class=\"small-12\">\n        <p>\n            <strong>{{name}}</strong>\n        </p>\n        <p>\n            Level {{level}}\n        </p>\n        <p>\n            Hp: {{hp}} / {{maxHp}}\n            <div class=\"progress\"><span class=\"meter\" style=\"width: {{hpPercentage}}%\"></span></div>\n        </p>\n        <p>\n            Location: {{locX}}, {{locY}}\n        </p>\n        <p>\n            Destination: {{destX}}, {{destY}}\n        </p>\n        <p>\n            Fights Won: {{fightsWon}} / {{fights}}\n        </p>\n        <p>\n            Locations Visited: {{myToon.LocationsVisited}}\n        </p>\n        <p>\n            <div class=\"label\">Weapon</div>\n            <div ng-show=\"myToon.Weapon.Level\">\n                Level {{myToon.Weapon.Level}}\n                {{myToon.Weapon.Name}}\n            </div>\n        </p>\n    </div>\n</div>",
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
      template: "<div class=\"row log-item\">\n    <div class=\"small-3 large-2 columns\">\n        <div class=\"label log-label {{labelClass}}\">\n            <small>\n                {{labelText | uppercase}}\n            </small>\n        </div>\n    </div>\n    <div class=\"small-9 large-10 columns\">\n        {{name}}\n        <span ng-bind-html-unsafe=\"action\"></span>\n\n        <div ng-show=\"item.Data.weaponWonName\">\n            <span class=\"label\">Weapon Acquired</span>\n            <span class=\"dim\">\n                Level {{item.Data.weaponWonLevel}} {{item.Data.weaponWonName}}\n            </span>\n        </div>\n        <div>\n            <small>\n                <ng-pluralize count=\"age\"\n                    when=\"{\n                        0: 'Moments ago',\n                        1: 'One minute ago',\n                        'other': '{} minutes ago'\n                    }\">\n                </ng-pluralize>\n            </small>\n        </div>\n    </div>\n</div>",
      link: function(scope) {
        var createdDate, now, parseDate;
        scope.name = $rootScope.myToon.Name;
        parseDate = function(goDate) {
          var matches, rex;
          rex = /(\d{4})-(\d{2})-(\d{2})T(\d{2}):(\d{2}):(\d{2}).+?-(\d{2}:\d{2})/g;
          matches = rex.exec(goDate);
          return new Date(matches[1], parseInt(matches[2]) - 1, matches[3], matches[4], matches[5], matches[6]);
        };
        createdDate = parseDate(scope.item.Created);
        now = new Date();
        scope.age = now.getUTCMinutes() - createdDate.getUTCMinutes();
        switch (scope.item.LogType) {
          case $rootScope.logTypes.fight:
            scope.action = "man-danced with a <span class=\"dim\">Level " + scope.item.Data.opponentLevel + " " + scope.item.Data.opponentName + "</span>.";
            if (scope.item.Data.victor) {
              scope.labelText = "victory";
              return scope.labelClass = "victory";
            } else {
              scope.labelText = "defeat";
              return scope.labelClass = "defeat";
            }
            break;
          case $rootScope.logTypes.locationDiscovery:
            return scope.action = "discovered the location of " + scope.item.Data.locationName + "!";
        }
      }
    };
  });

  angular.module(APP_NAME).controller('ToonLogCtrl', function($scope, $rootScope, $http, $timeout) {});

}).call(this);
