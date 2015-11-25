
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
    var imagesPromise = EchoService.get();
    imagesPromise.then(function (data) {
      console.log("Promise Data: " + data)
      var ii, key, newImageMap;
      var newImageMap = {};
      for (ii = 0; ii < data.length; ii++) {
        newImageMap[data[ii].Key] = data[ii];
      }
      for (Key in $scope.newImageMap) {
        $scope.imageMap[Key] = {};
            $scope.imageMap[Key].Repo = newImageMap[Key].Repo;
            $scope.imageMap[Key].Key = newImageMap[Key].Key;
            $scope.imageMap[Key].Created = newImageMap[Key].Created;
            $scope.imageMap[Key].Size = newImageMap[Key].Size;
        }
        console.log($scope.imageMap)
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
  echoService.get = function() {
    var getData = $q.defer();
    $http.get('/rest/dockerimages/')
      .success(function (data) {
        console.log(data);

        getData.resolve(data);
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
