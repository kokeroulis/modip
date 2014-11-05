'use strict';

modipControllers.controller('TeacherCtrl', ['$scope', 'TeacherService', '$state',
  function($scope, TeacherService, $state) {
    $scope.alerts = [];
    $scope.teacherCategoryList = null;

    TeacherService.categoryList().then(function(result) {
      $scope.teacherCategoryList = result;
    }, function(error, status) {
      $scope.alerts.push({msg: 'Σφάλμα συστήματος ' + error.body.Name, type: 'alert'})
    });

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

    $scope.saveData = function(data) {
      TeacherService.categorySave($state.params.subCategoryId, data)
    }
  }
]);
