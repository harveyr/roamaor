angular.module(DIRECTIVE_MODULE).directive "toonLog", ($rootScope) ->
    directive =
        replace: true
        scope:
            item: '='
        template: """
        <div class="row log-item">
            <div class="small-3 large-2 columns">
                <div class="label log-label {{labelClass}}">
                    <small>
                        {{labelText | uppercase}}
                    </small>
                </div>
            </div>
            <div class="small-9 large-10 columns">
                {{name}}
                <span ng-bind-html-unsafe="action"></span>
            </div>
        </div>
        """
        link: (scope) ->
            scope.name = $rootScope.myToon.Name

            # console.log 'scope.item.Data:', scope.item.Data
            switch scope.item.LogType

                when $rootScope.logTypes.fight
                    scope.action = "man-danced with a Level #{scope.item.Data.opponentLevel} #{scope.item.Data.opponentName}."
                    if scope.item.Data.victor
                        scope.labelText = "victory"
                        scope.labelClass = "victory"
                    else
                        scope.labelText = "defeat"
                        scope.labelClass = "defeat"

                when $rootScope.logTypes.locationDiscovery
                    scope.action = "discovered the location of #{scope.item.Data.locationName}!"

