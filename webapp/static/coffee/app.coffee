app = angular.module(APP_NAME, [
    "#{APP_NAME}.directives"
]).run ($rootScope) ->
    $rootScope.appName = "Roamoar"
