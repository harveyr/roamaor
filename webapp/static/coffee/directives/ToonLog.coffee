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

                <div ng-show="item.Data.weaponWonName">
                    <span class="label">Weapon Acquired</span>
                    <span class="dim">
                        Level {{item.Data.weaponWonLevel}} {{item.Data.weaponWonName}}
                    </span>
                </div>
                <div>
                    <small>
                        <ng-pluralize count="age"
                            when="{
                                0: 'Moments ago',
                                1: 'One minute ago',
                                'other': '{} minutes ago'
                            }">
                        </ng-pluralize>
                    </small>
                </div>
            </div>
        </div>
        """
        link: (scope) ->
            scope.name = $rootScope.myToon.Name

            parseDate = (goDate) ->
                rex = /(\d{4})-(\d{2})-(\d{2})T(\d{2}):(\d{2}):(\d{2}).+?-(\d{2}:\d{2})/g
                matches = rex.exec(goDate)
                new Date(matches[1], parseInt(matches[2]) - 1, matches[3], matches[4], matches[5], matches[6])

            createdDate = parseDate(scope.item.Created)
            now = new Date()
            scope.age = now.getUTCMinutes() - createdDate.getUTCMinutes()
            console.log 'scope.item.Data:', scope.item.Data

            switch scope.item.LogType

                when $rootScope.logTypes.fight
                    scope.action = """man-danced with a <span class="dim">Level #{scope.item.Data.opponentLevel} #{scope.item.Data.opponentName}</span>."""
                    if scope.item.Data.victor
                        scope.labelText = "victory"
                        scope.labelClass = "victory"
                    else
                        scope.labelText = "defeat"
                        scope.labelClass = "defeat"

                when $rootScope.logTypes.locationDiscovery
                    scope.action = "discovered the location of #{scope.item.Data.locationName}!"

