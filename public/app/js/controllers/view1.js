'use strict';

var modipControllers = angular.module('modipControllers', []);

modipControllers.controller('View1Ctrl', ['$scope',
  function($scope) {
    $scope.hello = 'hello';
  }
]);
