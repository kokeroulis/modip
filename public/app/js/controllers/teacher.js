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

    $scope.createNewAsset = function(newContent) {
      TeacherService.addAsset(newContent, $state.params.assetTypeId).then(function(result) {
        $scope.alerts = [];
        $scope.alerts.push({msg: "Η καινούργια εγγραφή προστέθηκε επιτυχώς", type: 'success'});
        //update our table
        TeacherService.info().then(function(data) {
          $scope.teacherInfo = data;
        });
      }, function(error, status) {
        $scope.alerts.push({msg: 'Σφάλμα συστήματος ' + error.body.Name, type:'danger'});
      });
    }

    $scope.modifyAsset = function(newContent, assetTypeId) {
      TeacherService.modifyAsset(assetTypeId, newContent).then(function(result) {
        $scope.alerts = [];
        $scope.alerts.push({msg: "Η εγγραφή επεξεργάστηκε επιτυχώς", type: 'success'});
        //update our table
        TeacherService.info().then(function(data) {
          $scope.teacherInfo = data;
        });
      }, function(error, status) {
        $scope.alerts.push({msg: 'Σφάλμα συστήματος ' + error.body.Name, type:'danger'});
      });
    }

    $scope.deleteAsset = function(assetTypeId) {
      TeacherService.deleteAsset(assetTypeId).then(function(result) {
        $scope.alerts = [];
        $scope.alerts.push({msg: "Η εγγραφή διαγράφτηκε επιτυχώς", type: 'success'});
        //update our table
        TeacherService.info().then(function(data) {
          $scope.teacherInfo = data;
        });
      }, function(error, status) {
        $scope.alerts.push({msg: 'Σφάλμα συστήματος ' + error.body.Name, type:'danger'});
      });
    }
  }
]);
