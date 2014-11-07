var modipDirectives = angular.module('modipDirectives', []);

modipDirectives.directive('formHelper', function() {
  return {

    restrict: 'E',
    scope: {
      categoryId: '=',
      subCategoryId: '='
    },

    controller: function($scope, $attrs, $http) {
      $scope.templateData;
      $scope.resultJson ;
      $http.get('app/data_json/' + $attrs.subcategoryid + '.json').
        success(function(data) {
          $scope.templateData = data;
        })

      $scope.createNewData = function() {
        var copyOriginal = [];
        angular.copy($scope.templateData.tableModels.data[0], copyOriginal)

        angular.forEach(copyOriginal, function(value) {
          //clear the data that already exists
          value.data = '';
        });

        $scope.templateData.tableModels.data.push(copyOriginal)
      }
    },

    templateUrl: 'app/templates/directives/form_helper.html'
  }

});
