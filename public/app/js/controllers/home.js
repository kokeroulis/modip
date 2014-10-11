'use strict';

var modipControllers = angular.module('modipControllers', []);

modipControllers.controller('HomeCtrl', ['$scope', 'TeacherService',
  function($scope, TeacherService) {
    $scope.alerts = [];
    $scope.loginTeacher = function(username, password) {
      TeacherService.login(username, password).then(function(result) {
        //it should redirect
      }, function(error) {
        //clear our alerts
        $scope.alerts = [];
        if (error.body.Code === 401 && error.body.Name === "Authorization Failed") {
          $scope.alerts.push({msg:'Λάθος όνομα χρήστη ή συνθηματικό!', type: 'alert'});
        } else {
          $scope.alerts.push({msg:'Σφάλμα συστήματος ' + error.body.Name , type: 'alert'});
        }
      });
    }
  }
]);
