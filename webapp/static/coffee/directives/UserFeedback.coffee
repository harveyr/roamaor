angular.module(DIRECTIVE_MODULE).directive 'userFeedback', () ->
    directive =
        replace: true
        template: """
            <div class="row" ng-show="userAlert">
                <div class="small-12">
                    <div data-alert class="alert-box">
                        {{userAlert}}
                        <a href="#" class="close">&times;</a>
                    </div>
                </div>
            </div>
        """
        link: (scope) ->
            # pass
