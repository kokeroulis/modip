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
        },
        controller: ['$scope', '$state', function($scope, $state) {
          $scope.currentSubCategoryId = $state.params.subCategoryId;
          $scope.dataJson = []
          $scope.dataTypes = [
            {id: 1, name: "Βιβλία-Mονογραφίες"},
            {id: 2, name: "Επιστημονικά Περιοδικά με Κριτές"},
            {id: 3, name: "Επιστημονικά Περιοδικά Χωρίς Κριτές"},
            {id: 4, name: "Πρακτικά Διεθνών Συνεδρίων με Kριτές"},
            {id: 5, name: "Πρακτικά Εθνικών Συνεδρίων με Kριτές"},
            {id: 6, name: "Ανακοινώσεις σε Διεθνή Επιστημονικά Συνέδρια"},
            {id: 7, name: "Ανακοινώσεις σε Εθνικά Επιστημονικά Συνέδρια"},
            {id: 8, name: "Άλλες Εργασίες"},
            {id: 9, name: "Κεφάλαια σε Συλλογικούς Τόμους"},
            {id: 10, name: "Άλλα"},
          ];
          $scope.createNewData = function() {
            $scope.dataJson.push(
            {
              author: '',
              title: '',
              magazine: false,
              publicator: '',
              publicationDate: null,
              type: 0
            })

          }
        }]
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
