
myApp = angular.module('myApp', []);

myApp.controller("IndexController", function($scope, EchoService) {

  $scope.address =  {
    Add1: "",
    Add2: "",
    City: "",
    State: "",
    Zip: ""
  };

  $scope.postConcat = "";

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
