angular.module(DIRECTIVE_MODULE).directive "toonSummary", ($rootScope) ->
    directive =
        replace: true
        scope: true
        template: """
        <div class="row">
            <div class="small-12">
                <p>
                    <strong>{{name}}</strong>
                </p>
                <p>
                    Hp: {{locX}}, {{locY}}
                </p>
                <p>
                    Location: {{locX}}, {{locY}}
                </p>
                <p>
                    Destination: {{destX}}, {{destY}}
                </p>
                <p>
                    Fights Won: {{fightsWon}} / {{fights}}
                </p>
            </div>
        </div>
        """
        link: (scope) ->
            applyToon = (toon) ->
                scope.name = toon.Name
                scope.locX = toon.LocX.toFixed(2)
                scope.locY = toon.LocY.toFixed(2)
                scope.fights = toon.Fights
                scope.fightsWon = toon.FightsWon
                scope.destX = toon.DestX.toFixed(2)
                scope.destY = toon.DestY.toFixed(2)

            if $rootScope.myToon
                applyToon $rootScope.myToon

            console.log '[toonSummary] $rootScope.myToon:', $rootScope.myToon

            $rootScope.$watch 'myToon', ->
                applyToon $rootScope.myToon
