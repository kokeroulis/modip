'use strict';

var modipServices = angular.module('modipServices', []);

modipServices.factory('TeacherService', ['$http', '$q', function($http, $q) {
  function reportErrorObj(errorName, errorStatus) {
    return {
      body: {
        Code: errorStatus ? errorStatus : 200,
        Name: errorName
      }
    };
  }

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

  function info() {
    /*var deffered = $q.defer();
    $http.get('teacher/info').success(function (result) {
      deffered.resolve(result.data);
    }).error(function(data, status) {
      deffered.reject({
        body: data.Common.error,
        status: status
      });
    });
    return deffered.promise;*/
    return {
      teacher: {
        id:1,
        name:'superteacher1',
        email:'superteacher1@teilar.gr',
        department: {
          id:1,
          name: 'T.P.T.'
        }
      },
      assetTypes: [
        {
          id :1,
          name: 'books',
          count: 10,
          assets: [
            {
              id: 1,
              assetType: 1,
              content: 'Super Book'
            }
          ]
        }
      ]
    }
  }

  function deleteAsset(assetId) {
    var deffered = $q.defer();

    var data = {
      assetId: assetId
    };

    $http.post('teacher/asset/remove', data).success(function (result) {
      if (result.Common.error.Name == 'InvalidAsset') {
        deffered.reject(reportErrorObj('InvalidAsset'));
      } else {
        deffered.resolve(result.data);
      }
    }).error(function(data, status) {
      deffered.reject(reportErrorObj(data.Common.error.Name, status));
    });

    return deffered.promise;


    return {
      id: 1,
      content: 'Super Book',
      assetType: 1
    }
  }

  function moveAsset(assetId, newAssetTypeId) {
    return {
      id: 1,
      content: 'Super Book',
      assetType: 2
    }
  }

  function modifyAsset(assetId, newContent) {
    return {
      id: 1,
      content: 'Modified Awesome Book',
      assetType: 1
    }
  }

  function addAsset(content, assetTypeId) {
    var deffered = $q.defer();

    var data = {
      content: content,
      assettype: assetTypeId
    };

    $http.post('teacher/asset/add', data).success(function (result) {
      if (result.Common.error.Name == 'AlreadyExists') {
        deffered.reject(reportErrorObj('AlreadyExists'));
      } else {
        deffered.resolve(result.data);
      }
    }).error(function(data, status) {
      deffered.reject(reportErrorObj(deffered, data.Common.error.Name));
    });

    return deffered.promise;
  }

  var service = {
    login: login,
    info: info,
    deleteAsset: deleteAsset,
    moveAsset: moveAsset,
    modifyAsset: modifyAsset,
    addAsset: addAsset
  };

  return service;
}]);
