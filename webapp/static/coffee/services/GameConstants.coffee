angular.module(SERVICE_MODULE).service 'GameConstants', ->
    @constants = {}

    @set = (key, value) ->
        @constants[key] = value

    @get = (key) ->
        @constants[key]
