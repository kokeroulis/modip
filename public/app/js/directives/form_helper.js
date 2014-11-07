var modipDirectives = angular.module('modipDirectives', []);

modipDirectives.directive('formHelper', function() {
  return {

    restrict: 'E',
    scope: {
      categoryId: '=',
      subCategoryId: '='
    },

    controller: function($scope, $attrs, $http, TeacherService) {
      $scope.templateData;
      $scope.resultJson ;
      $http.get('app/data_json/' + $attrs.subcategoryid + '.json').
        success(function(data) {
          $scope.templateData = data;
        })

      $scope.removeData = function(index) {
        if ($scope.templateData.tableModels.data.length === 1) {
          //we cannot remove the entire table
          return;
        }
        $scope.templateData.tableModels.data.splice(index,1);
      }

      $scope.saveData = function() {
        TeacherService.categorySave($attrs.subcategoryid, $scope.templateData).then(function(data) {
          console.log(data)
        }, function(error, status) {
          console.log(status)
          console.log(error)
        });
      }

      $scope.createNewData = function() {
        var copyOriginal = [];
        angular.copy($scope.templateData.tableModels.data[0], copyOriginal)

        angular.forEach(copyOriginal, function(value) {
          //clear the data that already exists
          value.data = '';
        });

        $scope.templateData.tableModels.data.push(copyOriginal);
      }
    },

    templateUrl: 'app/templates/directives/form_helper.html'
  }

});
