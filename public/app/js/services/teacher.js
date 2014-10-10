'use strict';

var modipServices = angular.module('modipServices', []);

modipServices.factory('TeacherService', ['$http', '$q', function($http, $q) {
  function login(username, password) {
    var data = {
      username: username,
      password: password
    };

    var deffered = $q.defer();
    $http.post('teacher/login', data).success(function (result) {
      deffered.resolve(result.data);
    }).error(function(data, status) {
      deffered.reject({
        body: data.Common.error,
        status: status
      });
    });
    return deffered.promise;
  }

  var service = {
    login: login
  };

  return service;
}]);
