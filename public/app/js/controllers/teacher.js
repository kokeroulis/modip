'use strict';

modipControllers.controller('TeacherCtrl', ['$scope', 'TeacherService', '$state',
  function($scope, TeacherService, $state) {
    $scope.alerts = [];
    $scope.teacherInfo = null;
    TeacherService.info().then(function(result) {
      $scope.teacherInfo = result;
    }, function(error, status) {
      $scope.alerts.push({msg: 'Σφάλμα συστήματος ' + error.body.Name, type: 'alert'})
    });

    $scope.createNewAsset = function(newContent, assetTypeId) {
      TeacherService.addAsset(newContent, $state.params.assetTypeId).then(function(result) {
        console.log("result")
        $scope.alerts = [];
        $scope.alerts.push({msg: "Η καινούργια εγγραφή προστέθηκε επιτυχώς", type: 'success'});
      }, function(error, status) {
        console.log(error)
        console.log(status)
      });
    }


  }
]);
