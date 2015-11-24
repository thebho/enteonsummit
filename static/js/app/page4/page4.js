
myApp = angular.module('myApp', []);

myApp.controller("IndexController", function($scope, EchoService) {

  $scope.newItem = {
    Name: "",
    Cost: 0
  };

  $scope.success = 0;
  $scope.items = [];

  $scope.addToDB = function() {
    var promise = EchoService.addToDB($scope.newItem);
    promise.then(function(data){
      $scope.success++;
    });
  };

  $scope.deleteAll = function() {
    var promise = EchoService.deleteAll();
    promise.then(function(data){
      if (data === "") {
        $scope.getAll();
      }
    });
  };

  $scope.deleteRecord = function(val) {
    var promise = EchoService.deleteRecord(val);
    promise.then(function(data){
      if (data === "") {
        $scope.getAll();
      }
    });
  };

  $scope.getAll = function() {
    var promise = EchoService.getAll();
    promise.then(function(data){
      $scope.items = data;
    });
  };
});

myApp.factory("EchoService", function($http, $q) {
  var echoService = {};

  echoService.get = function(val) {
    var getData = $q.defer();
    $http.get('/rest/echo4/' + val)
      .success(function (data) {
        getData.resolve(data.Val);
      });
    return getData.promise;
  };

  echoService.getAll = function() {
    var getData = $q.defer();
    $http.get('/rest/echo4/')
      .success(function (data) {
        getData.resolve(data.Val);
      });
    return getData.promise;
  };


  echoService.addToDB = function(val) {
    var postData = $q.defer();
    $http.post('/rest/echo4/', val)
    return postData.promise;
  };

  echoService.deleteAll = function() {
    var delData = $q.defer();
    $http.delete('/rest/echo4/')
    .success(function (data) {
      delData.resolve(data.Err);
    });
    return delData.promise;
  };

  echoService.deleteRecord = function(val) {
    var delData = $q.defer();
    $http.delete('/rest/echo4/' + val)
    .success(function (data) {
      delData.resolve(data.Err);
    });
    return delData.promise;
  };


  return echoService;
});
