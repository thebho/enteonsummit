
myApp = angular.module('myApp', []);

myApp.controller("IndexController", function($scope, EchoService) {

/*
type ImageData struct {
  Repo string
  Key string
  Created string
  Size string
}*/

  $scope.imageMap = {};

  $scope.getEcho = function() {
    var imagesPromise = EchoService.get($scope.organizationKey, $scope.tierKeys);
    imagesPromise.then(function (data) {
      var ii, key, newImageMap;
      newImageMap = {};
      for (ii = 0; ii < data.length; ii++) {
        newImageMap[data[ii].ID] = data[ii];
      }
      for (key in newImageMap) {
            $scope.imageMap[Key].Repo = newImageMap[Key].Repo;
            $scope.imageMap[Key].Key = newImageMap[Key].Key;
            $scope.imageMap[Key].Created = newImageMap[Key].Created;
            $scope.imageMap[Key].Size = newImageMap[Key].Size;


        }
      $scope.getRet = data;
    });
  };

  $scope.postEcho = function() {
    var promise = EchoService.post($scope.postVal);
    promise.then(function(data){
      $scope.postRet = data;
    });
  };
});

myApp.factory("EchoService", function($http, $q) {
  var echoService = {};
  echoService.get = function(val) {
    var getData = $q.defer();
    $http.get('/rest/echo/' + val)
      .success(function (data) {
        getData.resolve(data.Val);
      });
    return getData.promise;
  };

  echoService.post = function(val) {
    var postData = $q.defer(), valstruct = {};
    valstruct.Val = val;
    $http.post('/rest/echo/', valstruct)
      .success(function (data) {
        postData.resolve(data.Val);
      });
    return postData.promise;
  };

  return echoService;
});
