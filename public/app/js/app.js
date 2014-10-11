'use strict';

// Declare app level module which depends on views, and components
var modipApp = angular.module('modipApp', [
  'ngRoute',
  'modipControllers',
  'modipServices'
]);

modipApp.config(['$routeProvider',
  function($routeProvider) {
    $routeProvider.
      when('/', {
        templateUrl: 'app/templates/home/home.html',
        controller: 'HomeCtrl'
      }).
      otherwise({redirectTo: '/'})
  }
]);
