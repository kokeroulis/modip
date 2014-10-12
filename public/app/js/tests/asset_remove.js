describe('Teacher', function () {

  describe('Asset Delete', function() {

    var tester, service;

    beforeEach(function() {
      tester = ngMidwayTester('modipApp');
      service = tester.inject('TeacherService');
    });

    afterEach(function() {
      tester.destroy();
      tester = null;
    });

    it('should succeed', function(done) {
      var expected = {
        id: 1,
        content: 'Super Book',
        teacher: {
          id: 1,
          name: 'superteacher1',
          email: 'superteacher1@teilar.gr',
          department: {
            id: 1,
            name: 'T.P.T.'
          }
        },
        assetType: 1
      };

      service.deleteAsset(1).then(function(data) {
        expect(data).toEqual(expected)
        done();
      });
    });

    it('should fail', function(done) {
      var expected = {
        body: {
          Code: 200,
          Name: 'InvalidAsset'
        }
      };

      service.deleteAsset(1000).then(function(data) {},
      function(error) {
        expect(error).toEqual(expected)
        done();
      });
    });
  });

});

