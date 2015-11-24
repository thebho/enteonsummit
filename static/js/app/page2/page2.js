
myApp = angular.module('myApp', []);

myApp.controller("IndexController", function($scope, EchoService) {
  
  $scope.getVal = "";
  $scope.getRet = "";
  $scope.postVal = "";
  $scope.postRet = "";
  
  $scope.getEcho = function() {
    var promise = EchoService.get($scope.getVal);
    promise.then(function(data){
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