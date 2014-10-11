'use strict';

// Declare app level module which depends on views, and components
var modipApp = angular.module('modipApp', [
  'ui.router',
  'modipControllers',
  'modipServices'
]);

modipApp.config(['$stateProvider', '$urlRouterProvider',
  function($stateProvider, $urlRouterProvider) {
    // For any unmatched url, redirect to /
    $urlRouterProvider.otherwise("/");

    $stateProvider.
      state('home', {
        url: '/',
        templateUrl: 'app/templates/home/home.html',
        controller: 'HomeCtrl'
      })
  }
]);
