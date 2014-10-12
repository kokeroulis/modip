'use strict';

var modipControllers = angular.module('modipControllers', []);

modipControllers.controller('HomeCtrl', ['$scope', 'TeacherService', '$state',
  function($scope, TeacherService, $state) {
    $scope.alerts = [];
    $scope.loginTeacher = function(username, password) {
      TeacherService.login(username, password).then(function(result) {
        $state.go('teacher');
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
