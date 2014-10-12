'use strict';

modipControllers.controller('TeacherCtrl', ['$scope', 'TeacherService',
  function($scope, TeacherService) {
    $scope.alerts = [];
    $scope.teacherInfo = TeacherService.info();
  }
]);
