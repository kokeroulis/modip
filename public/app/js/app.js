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
      }).
      state('teacher', {
        url: '/teacher',
        templateUrl: 'app/templates/teacher/teacher.html',
        controller: 'TeacherCtrl'
      }).
      state('teacher.category', {
        url: '/category/{categoryId}',
        templateUrl: 'app/templates/teacher/teacher_category.html',
        controller: ['$scope', '$state', function($scope, $state) {
          $scope.currentCategoryId = $state.params.categoryId;
        }]
      }).
      state('teacher.category.subCategory', {
        url: '/subCategory/{subCategoryId}',
        templateUrl: function($state) {
          return 'app/templates/teacher/subCategories/' + $state.subCategoryId + '.html'
        }
      }).
      state('teacher.assetType', {
        url: '/assetType/{assetTypeId}',
        templateUrl: 'app/templates/teacher/teacher_assetType.html',
        controller: ['$scope', '$state', function($scope, $state) {
          $scope.currentAssetType = $state.params.assetTypeId;
        }]
      })
  }
]);
