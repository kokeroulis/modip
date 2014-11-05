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

  function assetOperation(url, data) {
    var deffered = $q.defer();

    $http.post(url, data).success(function (result) {
      if (result.Common.error.Name == 'InvalidAsset') {
        deffered.reject(reportErrorObj('InvalidAsset'));
      } else {
        deffered.resolve(result.data);
      }
    }).error(function(data, status) {
      deffered.reject(reportErrorObj(data.Common.error.Name, status));
    });

    return deffered.promise;
  }

  function login(username, password) {
    var data = {
      username: username,
      password: password
    };

    return assetOperation('teacher/login', data);
  }

  function info() {
    var deffered = $q.defer();
    $http.get('teacher/info').success(function (result) {
      deffered.resolve(result.data);
    }).error(function(data, status) {
      deffered.reject(reportErrorObj(data.Common.error.Name, status));
    });
    return deffered.promise;
  }

  function deleteAsset(assetId) {
    var deffered = $q.defer();

    var data = {
      assetId: parseInt(assetId)
    };

    return assetOperation('teacher/asset/remove', data);
  }

  function moveAsset(assetId, newAssetTypeId) {
    var data = {
      assetId: parseInt(assetId),
      assetTypeId: parseInt(newAssetTypeId)
    };

    return assetOperation('teacher/asset/move', data);
  }

  function modifyAsset(assetId, newContent) {
    var data = {
      assetId: parseInt(assetId),
      content: newContent
    };

    return assetOperation('teacher/asset/modify', data);
  }

  function addAsset(content, assetTypeId) {
    var deffered = $q.defer();

    var data = {
      content: content,
      assettype: parseInt(assetTypeId)
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

  function categoryList() {
    var deffered = $q.defer();
    $http.get('category/list').success(function (result) {
      deffered.resolve(result.data);
    }).error(function(data, status) {
      deffered.reject(reportErrorObj(data.Common.error.Name, status));
    });
    return deffered.promise;
  }

  function categorySave(subCategoryId, data) {
    var data = {
      data: angular.toJson(data),
      id: parseInt(subCategoryId)
    }

    var deffered = $q.defer();
    $http.post('category/save', data).success(function (result) {
      deffered.resolve(result.data);
    }).error(function(data, status) {
      deffered.reject(reportErrorObj(data.Common.error.Name, status));
    });
    return deffered.promise;
  }

  var service = {
    login: login,
    info: info,
    deleteAsset: deleteAsset,
    moveAsset: moveAsset,
    modifyAsset: modifyAsset,
    addAsset: addAsset,
    categoryList: categoryList,
    categorySave: categorySave
  };

  return service;
}]);
