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
                    Hp: {{hp}} / {{maxHp}}
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
                <p>
                    Locations Visited: {{myToon.LocationsVisited}}
                </p>
            </div>
        </div>
        """
        link: (scope) ->
            applyToon = (toon) ->
                scope.name = toon.Name
                scope.locX = toon.LocX.toFixed(2)
                scope.locY = toon.LocY.toFixed(2)
                scope.hp = toon.Hp
                scope.maxHp = toon.MaxHp
                scope.fights = toon.Fights
                scope.fightsWon = toon.FightsWon
                scope.destX = toon.DestX.toFixed(2)
                scope.destY = toon.DestY.toFixed(2)

            if $rootScope.myToon
                applyToon $rootScope.myToon

            console.log '[toonSummary] $rootScope.myToon:', $rootScope.myToon

            $rootScope.$watch 'myToon', ->
                if $rootScope.myToon
                    applyToon $rootScope.myToon
