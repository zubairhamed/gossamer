angular.module('gossamerApp', []).controller('EntityController', ['$scope','$http',
    function EntityController($scope, $http) {
        $scope.data = {};
        $http.get('v1.0/Datastreams').success(function(data) {

            var list = [];

            for (var i=0; i < data.value.length; i++) {
                var d = data.value[i];
                var o = {
                    text: d.unitOfMeasurement.name,
                    id: d["@iot.id"]
                };
                list.push(o)
            }

            $scope.data.count = data.count;
            $scope.data.entities = list;
            $scope.data.iconColor = "entity-box-col-mediumrare";

            $scope.entityClick = function(id) {
                $http.get('v1.0/Datastreams(' + id + ')').success(function(data) {
                    $scope.data.entity = {
                        id: data["@iot.id"],
                        text: data.unitOfMeasurement.name,
                    }
                });
            }
        });
    }
]);
