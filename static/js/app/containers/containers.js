
myApp = angular.module('myApp', []);

myApp.controller("IndexController", function($scope, EchoService) {

  $scope.containerMap = {};

  $scope.getEcho = function() {
    var containerPromise = echoService.get($scope.organizationKey, $scope.tierKeys);
    containerPromise.then(function (data) {
      var ii, key, newContainerMap;
      newImageMap = {};
      for (ii = 0; ii < data.length; ii++) {
        newImageMap[data[ii].ID] = data[ii];
      }
      for (key in newImageMap) {
            $scope.containerMap[Key].Image = newContainerMap[Key].Image;
            $scope.containerMap[Key].Key = newContainerMap[Key].Key;
            $scope.containerMap[Key].Created = newContainerMap[Key].Created;
            $scope.containerMap[Key].Size = newContainerMap[Key].Size;


        }
      $scope.getRet = data;
    });
  };

  $scope.postEcho = function() {
    var promise = EchoService.concat($scope.address);
    promise.then(function(data){
      $scope.postConcat = data;
    });
  };
});

myApp.factory("EchoService", function($http, $q) {
  var echoService = {};

  echoService.concat = function(val) {
    var postData = $q.defer();
    $http.post('/rest/echo3/', val)
      .success(function (data) {
        postData.resolve(data.Val);
      });
    return postData.promise;
  };

  return echoService;
});
