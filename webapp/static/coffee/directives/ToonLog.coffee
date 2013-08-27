angular.module(DIRECTIVE_MODULE).directive "toonLog", ($rootScope) ->
    directive =
        replace: true
        scope:
            item: '='
        template: """
        <div class="row log-item">
            <div class="small-12 columns">
                {{name}}
                {{action}}
            </div>
        </div>
        """
        link: (scope) ->
            scope.name = $rootScope.myToon.Name

            switch scope.item.LogType

                when $rootScope.logTypes.fight
                    scope.action = "man-danced with a Level #{scope.item.Data.opponentLevel} #{scope.item.Data.opponentName}"

                when $rootScope.logTypes.locationDiscovery
                    scope.action = "discovered the location of #{scope.item.Data.locationName}!"

