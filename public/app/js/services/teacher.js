'use strict';

var modipServices = angular.module('modipServices', []);

modipServices.factory('TeacherService', ['$http', '$q', function($http, $q) {
  function login(username, password) {
    var data = {
      username: username,
      password: password
    };

    var deffered = $q.defer();
    $http.post('http://localhost:3001/teacher/login', data).success(function (result) {
      return deffered.resolve(result.data);
    }).error(function(error, status) {
      return deffered.reject(error);
    });
  }

  return {
    login: login
  }
}]);
